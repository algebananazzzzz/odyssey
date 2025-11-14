package forms

import (
	"fmt"

	"github.com/algebananazzzzz/odyssey/cli/config"
	"github.com/algebananazzzzz/odyssey/cli/ui/styles"
	"github.com/charmbracelet/huh"
)

// NewMissingVariablePrompt returns a confirmation form prompting users to reconfigure missing variables.
func NewMissingVariablePrompt(missing config.ErrMissingProjectVariable, decision *bool) *huh.Form {
	description := fmt.Sprintf("The %q project variable is required for this template (%s). Reconfigure now to supply a value?", missing.Title, missing.Key)
	return newReconfigureForm("Project configuration needs attention", description, decision)
}

// NewCorruptionPrompt returns a confirmation form for handling corrupted project configs.
func NewCorruptionPrompt(decision *bool) *huh.Form {
	description := "Manual edits were detected in odyssey/project.yaml. Reconfigure now to replace infrastructure and workflow files?"
	return newReconfigureForm("Project configuration corrupted", description, decision)
}

func newReconfigureForm(title, description string, decision *bool) *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title(title).
				Description(description),
			huh.NewConfirm().
				Title("Reconfigure now?").
				Affirmative("Yes, reconfigure").
				Negative("No, exit").
				Value(decision),
		),
	).WithTheme(styles.NewTheme())
}
