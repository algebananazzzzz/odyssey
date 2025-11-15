package types

import "fmt"

type Template struct {
	Name        string            `yaml:"name"`
	Type        string            `yaml:"type"`
	Version     string            `yaml:"version"`
	Description string            `yaml:"description"`
	Projects    []TemplateProject `yaml:"options"`
}

// Project returns the project definition for the provided ID.
func (t *Template) Project(projectID string) (*TemplateProject, error) {
	if t == nil {
		return nil, fmt.Errorf("template is nil")
	}
	for i := range t.Projects {
		if t.Projects[i].ID == projectID {
			return &t.Projects[i], nil
		}
	}
	return nil, fmt.Errorf("project %q not found in template %q", projectID, t.Name)
}

// DisplayName returns a friendly label for the template.
func (t *Template) DisplayName() string {
	if t == nil {
		return "template"
	}
	if t.Name != "" {
		return t.Name
	}
	if t.Type != "" {
		return t.Type
	}
	return "template"
}

// Summary returns a description of the template.
func (t *Template) Summary() string {
	if t == nil {
		return "Select a template to see its summary."
	}
	if t.Description != "" {
		return t.Description
	}
	return fmt.Sprintf("%s template selected.", t.DisplayName())
}
