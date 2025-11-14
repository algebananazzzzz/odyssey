package operations

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/algebananazzzzz/odyssey/cli/constants"
	"github.com/algebananazzzzz/odyssey/cli/fileops"
	"github.com/algebananazzzzz/odyssey/cli/types"
)

// CloneProjectFiles clones the template repository and copies the relevant files.
func CloneProjectFiles(currentDir string, cfg types.Config) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		if cfg.ProjectConfig == nil {
			return fmt.Errorf("project config is nil")
		}
		projectCfg := cfg.ProjectConfig

		source, ok := constants.PROJECT_TEMPLATES[projectCfg.Template]
		if !ok {
			return fmt.Errorf("unknown template %q", projectCfg.Template)
		}

		tmpl, err := constants.LoadTemplate(projectCfg.Template)
		if err != nil {
			return fmt.Errorf("failed to load template %s: %w", projectCfg.Template, err)
		}

		project, err := tmpl.Project(projectCfg.Type)
		if err != nil {
			return fmt.Errorf("failed to load project definition: %w", err)
		}

		// Create a temporary directory for cloning
		tempDir, err := os.MkdirTemp("", "git-clone-")
		if err != nil {
			return fmt.Errorf("failed to create temp directory: %v", err)
		}
		defer os.RemoveAll(tempDir)

		repo := strings.TrimSpace(source.Repo)
		if repo == "" {
			return fmt.Errorf("template %q is missing a repository reference", projectCfg.Template)
		}
		repoURL := fmt.Sprintf("https://github.com/%s.git", repo)

		// Clone the repository without checkout
		cmd := exec.CommandContext(ctx, "git", "clone", "--no-checkout", repoURL, tempDir)

		if _, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("failed to clone repository: %v", err)
		}

		ref := source.Ref
		if ref == "" {
			ref = "main"
		}
		// Checkout the specific ref
		cmd = exec.CommandContext(ctx, "git", "-C", tempDir, "checkout", ref)
		if output, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("failed to checkout ref %s: %v, %s", ref, err, output)
		}

		if err := syncMappings(tempDir, currentDir, project.Infra.Paths); err != nil {
			return fmt.Errorf("failed to copy infra files: %w", err)
		}

		if err := renderTemplateMappings(tempDir, currentDir, project.Infra.Templates, project.Infra.Environments, cfg); err != nil {
			return fmt.Errorf("failed to render infra templates: %w", err)
		}

		if workflowPaths, err := selectWorkflowPaths(project.Workflows, projectCfg.Workflow); err != nil {
			return err
		} else if err := syncMappings(tempDir, currentDir, workflowPaths); err != nil {
			return fmt.Errorf("failed to copy workflow files: %w", err)
		}

		if err := syncMappings(tempDir, currentDir, project.Files.Paths); err != nil {
			return fmt.Errorf("failed to copy project files: %w", err)
		}

		return nil
	}
}

// InitGit initializes a git repository in the given directory.
func InitGit(currentDir string) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		// Initialize a new repository
		cmd := exec.CommandContext(ctx, "git", "init")
		cmd.Dir = currentDir

		if _, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("failed to initialize repository: %v", err)
		}

		return nil
	}
}

func syncMappings(baseDir, currentDir string, mappings []types.PathMapping) error {
	for _, mapping := range mappings {
		if mapping.Src == "" || mapping.Dest == "" {
			continue
		}

		src := filepath.Join(baseDir, mapping.Src)
		dest := mapping.Dest
		if !filepath.IsAbs(dest) {
			dest = filepath.Join(currentDir, dest)
		}

		if err := fileops.Sync(src, dest); err != nil {
			return fmt.Errorf("failed to sync %s -> %s: %w", mapping.Src, dest, err)
		}
	}
	return nil
}

func renderTemplateMappings(baseDir, currentDir string, mappings []types.PathMapping, environments []string, cfg types.Config) error {
	if len(mappings) == 0 {
		return nil
	}

	envs := environments
	if len(envs) == 0 {
		envs = []string{""}
	}

	for _, env := range envs {
		envCfg := cfg
		envCfg.Env = env
		for _, mapping := range mappings {
			src := filepath.Join(baseDir, mapping.Src)
			dest := filepath.Join(currentDir, replaceEnvToken(mapping.Dest, env))
			if err := renderTemplatePath(src, dest, envCfg); err != nil {
				return err
			}
		}
	}

	return nil
}

func renderTemplatePath(src, dest string, cfg types.Config) error {
	info, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("failed to stat template source %s: %w", src, err)
	}

	if info.IsDir() {
		return filepath.WalkDir(src, func(path string, d fs.DirEntry, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}
			rel, err := filepath.Rel(src, path)
			if err != nil {
				return err
			}
			target := filepath.Join(dest, rel)
			if d.IsDir() {
				return os.MkdirAll(target, 0o755)
			}
			return renderTemplateFile(path, target, cfg)
		})
	}

	return renderTemplateFile(src, dest, cfg)
}

func renderTemplateFile(src, dest string, cfg types.Config) error {
	content, err := os.ReadFile(src)
	if err != nil {
		return fmt.Errorf("failed to read template %s: %w", src, err)
	}

	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		return fmt.Errorf("failed to create template destination %s: %w", dest, err)
	}

	tmpl, err := template.New(filepath.Base(src)).Parse(string(content))
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", src, err)
	}

	file, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("failed to create template output %s: %w", dest, err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, cfg); err != nil {
		return fmt.Errorf("failed to render template %s: %w", src, err)
	}

	return nil
}

func replaceEnvToken(path, env string) string {
	return strings.ReplaceAll(path, "${{environment}}", env)
}

func selectWorkflowPaths(workflows []types.Workflow, selectedID string) ([]types.PathMapping, error) {
	if len(workflows) == 0 {
		return nil, nil
	}
	if selectedID == "" {
		var paths []types.PathMapping
		for _, wf := range workflows {
			paths = append(paths, wf.Paths...)
		}
		return paths, nil
	}
	for _, wf := range workflows {
		if wf.ID == selectedID {
			return wf.Paths, nil
		}
	}
	return nil, fmt.Errorf("selected workflow %q not found", selectedID)
}
