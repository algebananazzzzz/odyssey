package formatters

import (
	"fmt"

	"github.com/algebananazzzzz/odyssey/cli/constants"
	"github.com/algebananazzzzz/odyssey/cli/types"
	"github.com/algebananazzzzz/odyssey/cli/ui/styles"
)

func PrintProjectSummary(cfg *types.ProjectConfig) string {
	labelStyle := styles.LabelStyle.
		Bold(true)

	// Build the summary text
	content := fmt.Sprintf(
		"%s %s\n%s %s\n%s %s\n",
		labelStyle.Render("Project Code:"), styles.ValueStyle.Render(cfg.Code),
		labelStyle.Render("Project Type:"), styles.ValueStyle.Render(constants.PROJECT_TEMPLATES[cfg.Type].Name),
		labelStyle.Render("Environments:"), styles.ValueStyle.Render(constants.ENVIRONMENTS[cfg.Environments]),
	)

	// Wrap with header + border
	return styles.BorderStyle.Render(
		styles.HeaderStyle.Render("Odyssey Project Configuration") + "\n\n" + content,
	)
}
