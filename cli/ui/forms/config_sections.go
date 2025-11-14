package forms

import (
	"errors"
	"fmt"
	"strings"

	"github.com/algebananazzzzz/odyssey/cli/constants"
	"github.com/algebananazzzzz/odyssey/cli/ui/formatters"
	"github.com/algebananazzzzz/odyssey/cli/validators"
	"github.com/charmbracelet/huh"
)

// welcomeGroup renders the opening splash + consent.
func (b *configFormBuilder) welcomeGroup() *huh.Group {
	return huh.NewGroup(
		huh.NewNote().
			Title(fmt.Sprintf("Welcome to Odyssey CLI - %s", constants.Version)).
			Description("Let's start creating a production-ready AWS project! I'll generate Terraform infrastructure, set up CI/CD pipelines, and create the project files so you can start developing quickly and safely."),

		huh.NewConfirm().
			Title("Sounds cool?").
			Affirmative("Yes!! ☁️").
			Negative("Lets Go! 🌱").
			Value(b.confirm),
	)
}

func (b *configFormBuilder) globalSettingsGroup() *huh.Group {
	return huh.NewGroup(
		huh.NewNote().
			Title("Configure Global Settings").
			Description("These settings apply across all Odyssey projects you manage. They define where Terraform state will live."),

		huh.NewInput().
			Title("Bucket").
			Description("S3 bucket where Odyssey will store Terraform state files.").
			Placeholder("my-odyssey-bucket").
			Value(&b.cfg.GlobalConfig.Bucket).
			Validate(validators.NotEmpty),

		huh.NewInput().
			Title("Workspace Key Prefix").
			Description("Folder/prefix inside the bucket for separating state files.").
			Placeholder("tfstate").
			Value(&b.cfg.GlobalConfig.WorkspaceKeyPrefix).
			Validate(validators.Alphanumeric),

		huh.NewInput().
			Title("Region").
			Description("AWS region hosting the shared infrastructure (e.g., ap-southeast-1).").
			Placeholder("ap-southeast-1").
			Value(&b.cfg.GlobalConfig.Region).
			Validate(validators.AWSRegion),
	)
}

// projectSettingsGroup handles project code + destination folder.
func (b *configFormBuilder) projectSettingsGroup() *huh.Group {
	dirField := huh.NewInput().
		Title("Project Directory").
		Description("Odyssey will create your project in this directory. Update this if you prefer a different path.").
		Value(&b.cfg.ProjectConfig.Directory).
		Validate(validators.FilePath)

	codeField := huh.NewInput().
		Title("Project Code").
		Placeholder("my-project").
		Value(&b.cfg.ProjectConfig.Code).
		Validate(func(val string) error {
			if err := validators.Alphanumeric(val); err != nil {
				return err
			}
			newDir := defaultProjectDirectory(b.cfg.ProjectConfig.Code)
			if newDir != b.cfg.ProjectConfig.Directory {
				b.cfg.ProjectConfig.Directory = newDir
				dirField.Value(&b.cfg.ProjectConfig.Directory)
			}
			return nil
		})

	return huh.NewGroup(
		huh.NewNote().
			Title("Project Settings").
			Description("Name your project and choose where Odyssey should generate the files."),
		codeField,
		dirField,
	)
}

// templateSelectionGroup renders template + project selects.
func (b *configFormBuilder) templateSelectionGroup() *huh.Group {
	projectSelect := huh.NewSelect[string]().
		Title("Project Type").
		Description("Select the specific project preset within the chosen template.").
		Value(&b.cfg.ProjectConfig.Type).
		OptionsFunc(func() []huh.Option[string] {
			return projectOptions(b.cfg.ProjectConfig)
		}, b.cfg.ProjectConfig).
		Validate(func(val string) error {
			if b.cfg.ProjectConfig.Template == "" {
				return errors.New("select a template first")
			}
			if len(projectOptions(b.cfg.ProjectConfig)) == 0 {
				return errors.New("template has no project types")
			}
			if strings.TrimSpace(val) == "" {
				return errors.New("project type is required")
			}
			return nil
		})
	b.projectSelect = projectSelect

	return huh.NewGroup(
		huh.NewNote().
			Title("Choose Template & Project Type").
			Description("Pick the Odyssey template repository and project preset you want to use."),

		huh.NewSelect[string]().
			Title("Template").
			Description("Select the template repository you want to use.").
			Value(&b.cfg.ProjectConfig.Template).
			Options(templateSelectOptions()...).
			Validate(func(val string) error {
				if strings.TrimSpace(val) == "" {
					return errors.New("template is required")
				}
				if val != b.lastTemplate {
					b.lastTemplate = val
					b.resetTemplateSelections()
				}
				return nil
			}),

		huh.NewNote().
			Title("Template description").
			DescriptionFunc(func() string {
				tmpl, err := loadTemplate(b.cfg.ProjectConfig)
				if err != nil {
					if b.cfg.ProjectConfig == nil || b.cfg.ProjectConfig.Template == "" {
						return "Select a template to see its summary."
					}
					return fmt.Sprintf("Failed to load template details: %v", err)
				}
				return tmpl.Summary()
			}, b.cfg.ProjectConfig),

		projectSelect,

		huh.NewNote().
			Title("Project description").
			DescriptionFunc(func() string {
				if b.cfg.ProjectConfig == nil || b.cfg.ProjectConfig.Type == "" {
					return "Select a template and project type to see its summary."
				}
				project, err := loadProject(b.cfg.ProjectConfig)
				if err != nil {
					return fmt.Sprintf("Failed to load project details: %v", err)
				}
				return project.Summary()
			}, b.cfg.ProjectConfig),
	)
}

// workflowSelectionGroup renders workflow/exposure selection.
func (b *configFormBuilder) workflowSelectionGroup() *huh.Group {
	workflowSelect := huh.NewSelect[string]().
		Title("Workflow").
		Description("Select the workflow variant you want to generate for this project.").
		Value(&b.cfg.ProjectConfig.Workflow).
		OptionsFunc(func() []huh.Option[string] {
			return workflowOptions(b.cfg.ProjectConfig)
		}, b.cfg.ProjectConfig).
		Validate(func(val string) error {
			if b.cfg.ProjectConfig.Template == "" || b.cfg.ProjectConfig.Type == "" {
				return nil
			}
			if len(workflowOptions(b.cfg.ProjectConfig)) == 0 {
				return nil
			}
			if strings.TrimSpace(val) == "" {
				return errors.New("workflow is required")
			}
			return nil
		})
	b.workflowSelect = workflowSelect

	return huh.NewGroup(
		workflowSelect,
		huh.NewNote().
			Title("Workflow summary").
			DescriptionFunc(func() string {
				if b.cfg.ProjectConfig == nil || b.cfg.ProjectConfig.Template == "" || b.cfg.ProjectConfig.Type == "" {
					return "Select a template and project to view available workflows."
				}
				project, err := loadProject(b.cfg.ProjectConfig)
				if err != nil {
					return fmt.Sprintf("Failed to load workflow details: %v", err)
				}
				if len(project.Workflows) == 0 {
					return "This project does not define workflow variants."
				}
				if b.cfg.ProjectConfig.Workflow == "" {
					return "Select a workflow to continue."
				}
				wf, err := project.Workflow(b.cfg.ProjectConfig.Workflow)
				if err != nil {
					return fmt.Sprintf("Failed to load workflow details: %v", err)
				}
				return wf.Summary()
			}, b.cfg.ProjectConfig),
	)
}

// finalConfirmationGroup renders the closing summary + confirm.
func (b *configFormBuilder) finalConfirmationGroup() *huh.Group {
	note := huh.NewNote().
		Title("Final Confirmation").
		DescriptionFunc(func() string {
			return fmt.Sprintf(
				"Review your selections:\n\n%s",
				formatters.PrintProjectSummary(b.cfg.ProjectConfig),
			)
		}, b.cfg.ProjectConfig)

	return huh.NewGroup(
		note,
		huh.NewConfirm().
			Title("Confirm Setup").
			Affirmative("Yes, proceed").
			Negative("No, exit").
			Value(b.confirm),
	)
}
