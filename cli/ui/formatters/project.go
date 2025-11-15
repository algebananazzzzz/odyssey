package formatters

import (
	"fmt"
	"log"
	"strings"

	"github.com/algebananazzzzz/odyssey/cli/constants"
	"github.com/algebananazzzzz/odyssey/cli/types"
	"github.com/algebananazzzzz/odyssey/cli/ui/styles"
	"github.com/charmbracelet/lipgloss"
)

// PrintProjectSummary returns a formatted summary of the current project settings.
func PrintProjectSummary(cfg *types.ProjectConfig) string {
	labelStyle := styles.LabelStyle.
		Bold(true)

	var projectDef *types.TemplateProject
	if cfg.Template != "" {
		if tmpl, err := constants.LoadTemplate(cfg.Template); err != nil {
			log.Printf("failed to load template %q: %v", cfg.Template, err)
		} else {
			if cfg.Type != "" {
				if project, err := tmpl.Project(cfg.Type); err == nil {
					projectDef = project
				} else {
					log.Printf("failed to load project %q: %v", cfg.Type, err)
				}
			}
		}
	}

	templateLabel := "Unknown"
	if tmpl, ok := constants.PROJECT_TEMPLATES[cfg.Template]; ok && tmpl.Name != "" {
		templateLabel = tmpl.Name
	} else if cfg.Template != "" {
		templateLabel = fmt.Sprintf("custom (%s)", cfg.Template)
	}

	projectLabel := "Unknown"
	if projectDef != nil {
		projectLabel = projectDef.DisplayName()
	} else if cfg.Type != "" {
		projectLabel = fmt.Sprintf("custom (%s)", cfg.Type)
	}

	workflowLabel := "Unknown"
	if projectDef != nil && cfg.Workflow != "" {
		if wf, err := projectDef.Workflow(cfg.Workflow); err == nil {
			if wf.Description != "" {
				workflowLabel = wf.Description
			} else {
				workflowLabel = wf.ID
			}
		} else {
			workflowLabel = fmt.Sprintf("custom (%s)", cfg.Workflow)
		}
	}

	envLabel := "Unknown"
	if projectDef != nil {
		envs := projectDef.WorkflowEnvironments(cfg.Workflow)
		if len(envs) > 0 {
			envLabel = strings.Join(envs, ", ")
		}
	}

	lines := []string{
		fmt.Sprintf("%s %s", labelStyle.Render("Project Code:"), styles.ValueStyle.Render(cfg.Code)),
		fmt.Sprintf("%s %s", labelStyle.Render("Directory:"), styles.ValueStyle.Render(cfg.Directory)),
		fmt.Sprintf("%s %s", labelStyle.Render("Template:"), styles.ValueStyle.Render(templateLabel)),
		fmt.Sprintf("%s %s", labelStyle.Render("Project Type:"), styles.ValueStyle.Render(projectLabel)),
		fmt.Sprintf("%s %s", labelStyle.Render("Workflow:"), styles.ValueStyle.Render(workflowLabel)),
		fmt.Sprintf("%s %s", labelStyle.Render("Environments:"), styles.ValueStyle.Render(envLabel)),
	}

	varVariableLines := variableSummaryLines(cfg, projectDef, labelStyle)
	if len(varVariableLines) > 0 {
		lines = append(lines, "")
		lines = append(lines, varVariableLines...)
	}

	content := strings.Join(lines, "\n")

	// Wrap with header + border
	return styles.BorderStyle.Render(
		styles.HeaderStyle.Render("Odyssey Project Configuration") + "\n\n" + content,
	)
}

// variableSummaryLines renders the configured template variables for display.
func variableSummaryLines(cfg *types.ProjectConfig, project *types.TemplateProject, labelStyle lipgloss.Style) []string {
	var lines []string

	if project == nil || len(project.Variables) == 0 {
		return nil
	}

	for _, def := range project.Variables {
		value := cfg.Variables[def.Key]
		if value == "" {
			value = "<not set>"
		}
		lines = append(lines, fmt.Sprintf("%s %s", labelStyle.Render(def.Title+":"), styles.ValueStyle.Render(value)))
	}

	return lines
}
