package types

import "fmt"

type TemplateProject struct {
	ID               string       `yaml:"id"`
	Name             string       `yaml:"name"`
	ShortDescription string       `yaml:"short_description"`
	Description      string       `yaml:"description"`
	Workflows        []Workflow   `yaml:"workflows"`
	Infra            Infra        `yaml:"infra"`
	Files            ProjectFiles `yaml:"project"`
	Submodules       []Submodule  `yaml:"submodules"`
	Variables        []Variable   `yaml:"variables"`
}

type Workflow struct {
	ID                   string        `yaml:"id"`
	Description          string        `yaml:"description"`
	RequiredEnvironments []string      `yaml:"required_environments"`
	Paths                []PathMapping `yaml:"paths"`
}

type Infra struct {
	Environments []string      `yaml:"environments"`
	Templates    []PathMapping `yaml:"templates"`
	Paths        []PathMapping `yaml:"paths"`
}

type ProjectFiles struct {
	Paths []PathMapping `yaml:"paths"`
}

type PathMapping struct {
	Src  string `yaml:"src"`
	Dest string `yaml:"dest"`
}

type Submodule struct {
	Repo string `yaml:"repo"`
	Ref  string `yaml:"ref"`
	Path string `yaml:"path"`
}

type Variable struct {
	Key         string `yaml:"key"`
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Placeholder string `yaml:"placeholder"`
	Required    bool   `yaml:"required"`
}

// DisplayName returns a human-friendly label for the project.
func (p *TemplateProject) DisplayName() string {
	if p == nil {
		return "project"
	}
	if p.Name != "" {
		return p.Name
	}
	return p.ID
}

// Workflow returns the workflow definition for the provided ID.
func (p *TemplateProject) Workflow(workflowID string) (*Workflow, error) {
	if p == nil {
		return nil, fmt.Errorf("project is nil")
	}
	for i := range p.Workflows {
		if p.Workflows[i].ID == workflowID {
			return &p.Workflows[i], nil
		}
	}
	return nil, fmt.Errorf("workflow %q not found in project %q", workflowID, p.ID)
}

// WorkflowEnvironments returns the required environments for a workflow.
func (p *TemplateProject) WorkflowEnvironments(workflowID string) []string {
	wf, err := p.Workflow(workflowID)
	if err != nil || len(wf.RequiredEnvironments) == 0 {
		return []string{"prd"}
	}
	envs := make([]string, len(wf.RequiredEnvironments))
	copy(envs, wf.RequiredEnvironments)
	return envs
}

// Summary returns the combined short + long description for the project.
func (p *TemplateProject) Summary() string {
	if p == nil {
		return "Select a template and project type to see its summary."
	}
	switch {
	case p.ShortDescription != "" && p.Description != "":
		return fmt.Sprintf("%s\n%s", p.ShortDescription, p.Description)
	case p.Description != "":
		return p.Description
	case p.ShortDescription != "":
		return p.ShortDescription
	default:
		return fmt.Sprintf("%s selected.", p.DisplayName())
	}
}

// Summary returns the workflow description.
func (w *Workflow) Summary() string {
	if w == nil {
		return "Select a workflow to continue."
	}
	if w.Description != "" {
		return w.Description
	}
	if w.ID != "" {
		return fmt.Sprintf("Workflow %s selected.", w.ID)
	}
	return "Workflow selected."
}
