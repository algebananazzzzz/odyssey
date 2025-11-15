package types

// TemplateCatalog represents the embedded catalog of available templates.
type TemplateCatalog struct {
	Projects map[string]TemplateSource `yaml:"projects"`
}

// TemplateSource describes the Git reference for a remotely hosted template definition.
type TemplateSource struct {
	Name string `yaml:"name"`
	Repo string `yaml:"repo"`
	Ref  string `yaml:"ref"`
	File string `yaml:"file"`
}

// FilePath returns the explicit file path or the default template config name.
func (ts TemplateSource) FilePath() string {
	if ts.File == "" {
		return "odyssey.template-config.yaml"
	}
	return ts.File
}
