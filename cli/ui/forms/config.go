package forms

import (
	"fmt"

	"github.com/algebananazzzzz/odyssey/cli/constants"
	"github.com/algebananazzzzz/odyssey/cli/types"
	"github.com/algebananazzzzz/odyssey/cli/ui/styles"
	"github.com/algebananazzzzz/odyssey/cli/validators"
	"github.com/charmbracelet/huh"
)

func NewConfigForm(globalConfig *types.GlobalConfig, projectConfig *types.ProjectConfig, confirm *bool) *huh.Form {

	projectOpts := make([]huh.Option[string], 0, len(constants.PROJECT_TEMPLATES))
	for value, tmpl := range constants.PROJECT_TEMPLATES {
		projectOpts = append(projectOpts, huh.NewOption(tmpl.Name, value))
	}

	envOpts := make([]huh.Option[int], 0, len(constants.ENVIRONMENTS))
	for value, label := range constants.ENVIRONMENTS {
		envOpts = append(envOpts, huh.NewOption(label, value))
	}

	return huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title(fmt.Sprintf("Welcome to Odyssey CLI - %s", constants.Version)).
				Description("Let's start creating a production-ready AWS project! I'll generate Terraform infrastructure, set up CI/CD pipelines, and create the project files so you can start developing quickly and safely."),

			huh.NewConfirm().
				Title("Sounds cool?").
				Affirmative("Spin up the cloud! ☁️").
				Negative("Let's start terraforming 🌱").
				Value(confirm),
		),
		huh.NewGroup(
			huh.NewNote().
				Title("Configure Global Settings").
				Description("These settings apply across all Odyssey projects you manage. They define where your Terraform state will be stored."),

			huh.NewInput().
				Title("Bucket").
				Description("The AWS S3 bucket where Odyssey will store Terraform state files. Ensure the bucket already exists in your AWS account.").
				Placeholder("my-odyssey-bucket").
				Key("bucket").
				Value(&globalConfig.Bucket).
				Validate(validators.NotEmpty),

			huh.NewInput().
				Title("Workspace Key Prefix").
				Description("A folder prefix inside your bucket to separate state files for different purposes.").
				Placeholder("tfstate").
				Key("workspace_key_prefix").
				Value(&globalConfig.WorkspaceKeyPrefix).
				Validate(validators.Alphanumeric),

			huh.NewInput().
				Title("Region").
				Description("The AWS region to deploy your AWS resources to (e.g., ap-southeast-1).").
				Placeholder("ap-southeast-1").
				Key("region").
				Value(&globalConfig.Region).
				Validate(validators.AWSRegion),
		),
		huh.NewGroup(
			huh.NewNote().
				Title("Configure Project Settings").
				Description("Define the specifics of the project you are setting up. Odyssey will generate infrastructure, CI/CD, and project files based on these settings."),
			huh.NewInput().
				Title("Project Code").
				Description("A unique identifier for this project (used in resource naming and state management).").
				Placeholder("my-project").
				Value(&projectConfig.Code).
				Validate(validators.Alphanumeric),

			huh.NewSelect[string]().
				Title("Project Type").
				Description("Select the type of project you are creating. This determines what infrastructure Odyssey will generate.").
				Value(&projectConfig.Type).
				Options(projectOpts...),

			huh.NewSelect[int]().
				Title("Environments").
				Description("Choose how many environments Odyssey should set up. More environments allow for safer testing before production.").
				Value(&projectConfig.Environments).
				Options(envOpts...),
		),
		huh.NewGroup(
			huh.NewNote().
				Title("Final Confirmation").
				Description("Review your selections above. Once confirmed, Odyssey will generate all required project files, infrastructure templates, and CI/CD configuration."),
			huh.NewConfirm().
				Title("Confirm Setup").
				Affirmative("Yes, proceed").
				Negative("No, exit").
				Value(confirm),
		),
	).WithTheme(styles.NewTheme())
}
