package operations

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/algebananazzzzz/odyssey/cli/constants"
	"github.com/algebananazzzzz/odyssey/cli/fileops"
	"github.com/algebananazzzzz/odyssey/cli/types"
)

// isEmptyDirectory returns true if the directory is empty or only contains hidden files (starting with .) or the "odyssey" folder.
func isEmptyDirectory(path string) (bool, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return false, err
	}

	for _, e := range entries {
		name := e.Name()
		// Ignore hidden files/folders and "odyssey"
		if name != "odyssey" && name[0] != '.' {
			return false, nil
		}
	}

	return true, nil
}

func CloneProjectFiles(currentDir string, config types.ProjectConfig) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		empty, err := isEmptyDirectory(currentDir)
		if err != nil {
			return fmt.Errorf("failed to check directory contents: %v", err)
		}

		// Create a temporary directory for cloning
		tempDir, err := os.MkdirTemp("", "git-clone-")
		if err != nil {
			return fmt.Errorf("failed to create temp directory: %v", err)
		}
		defer os.RemoveAll(tempDir)

		// Clone the repository without checkout
		cmd := exec.CommandContext(ctx, "git", "clone", "--no-checkout", constants.ODYSSEY_GIT_URL, tempDir)

		if _, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("failed to clone repository: %v", err)
		}

		// Checkout the specific commit
		cmd = exec.CommandContext(ctx, "git", "-C", tempDir, "checkout", constants.CommitSHA)
		if output, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("failed to checkout commit %s: %v, %s", constants.CommitSHA, err, output)
		}

		projectSrc, hasProjectFiles := constants.ProjectSrc(tempDir, config)

		infraSrc, infraDest := constants.InfraPaths(tempDir, currentDir, config)
		cicdSrc, cicdDest := constants.CICDPaths(tempDir, currentDir, config)

		if empty && hasProjectFiles {
			if err := fileops.CopyDir(projectSrc, currentDir); err != nil {
				return fmt.Errorf("failed to copy project files: %v", err)
			}
		}

		if err := fileops.CopyOrReplace(infraSrc, infraDest); err != nil {
			return fmt.Errorf("failed to replace infra files: %v", err)
		}
		if err := fileops.CopyOrReplace(cicdSrc, cicdDest); err != nil {
			return fmt.Errorf("failed to replace workflow files: %v", err)
		}

		return nil
	}
}

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

func AddSubmodule(repoPath, submodulePath, submoduleURL string) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		cmd := exec.CommandContext(ctx, "git", "submodule", "add", submoduleURL, submodulePath)
		cmd.Dir = repoPath

		// Run the command
		output, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("failed to add submodule: %v, output: %s", err, string(output))
		}

		return nil
	}
}
