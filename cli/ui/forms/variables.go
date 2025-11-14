package forms

import (
	"fmt"
	"log"

	"github.com/algebananazzzzz/odyssey/cli/constants"
	"github.com/algebananazzzzz/odyssey/cli/types"
	"github.com/algebananazzzzz/odyssey/cli/validators"
	"github.com/charmbracelet/huh"
)

// projectVariableGroup renders the variable inputs for the selected template/project pair.
func projectVariableGroup(cfg types.Config) *huh.Group {
	if cfg.ProjectConfig == nil {
		return nil
	}
	if cfg.ProjectConfig.Template == "" || cfg.ProjectConfig.Type == "" {
		return nil
	}

	tmpl, err := constants.LoadTemplate(cfg.ProjectConfig.Template)
	if err != nil {
		log.Printf("failed to load template %q: %v", cfg.ProjectConfig.Template, err)
		return nil
	}

	project, err := tmpl.Project(cfg.ProjectConfig.Type)
	if err != nil {
		log.Printf("failed to load project %q for template %q: %v", cfg.ProjectConfig.Type, cfg.ProjectConfig.Template, err)
		return nil
	}

	variables := project.Variables
	if len(variables) == 0 {
		return nil
	}

	fields := []huh.Field{
		huh.NewNote().
			Title(fmt.Sprintf("Configure %s-specific settings", projectDisplayName(project))).
			Description(variableGroupDescription(project.ShortDescription)),
	}

	for _, variable := range variables {
		varDef := variable
		accessor := newVariableAccessor(cfg.ProjectConfig.Variables, varDef.Key)

		field := huh.NewInput().
			Title(varDef.Title).
			Description(varDef.Description).
			Placeholder(varDef.Placeholder).
			Key(varDef.Key).
			Accessor(accessor)

		if varDef.Required {
			field = field.Validate(validators.NotEmpty)
		}

		fields = append(fields, field)
	}

	return huh.NewGroup(fields...)
}

// ensureProjectVariables makes sure the config contains a map for variables.
func ensureProjectVariables(cfg *types.ProjectConfig) {
	if cfg == nil {
		return
	}
	if cfg.Variables == nil {
		cfg.Variables = make(map[string]string)
	}
}

// variableAccessor binds a template variable key to the shared map.
type variableAccessor struct {
	vars map[string]string
	key  string
}

// newVariableAccessor returns an Accessor backed by the shared variables map.
func newVariableAccessor(vars map[string]string, key string) huh.Accessor[string] {
	return &variableAccessor{
		vars: vars,
		key:  key,
	}
}

// Get returns the value for the accessor's key.
func (v *variableAccessor) Get() string {
	return v.vars[v.key]
}

// Set updates the value for the accessor's key.
func (v *variableAccessor) Set(value string) {
	v.vars[v.key] = value
}

func variableGroupDescription(shortDesc string) string {
	if shortDesc == "" {
		return "These values are specific to the project type you selected."
	}
	return shortDesc
}

func projectDisplayName(project *types.TemplateProject) string {
	if project == nil {
		return "project"
	}
	if project.Name != "" {
		return project.Name
	}
	return project.ID
}
