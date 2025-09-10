package types

type TemplateConfig struct {
	Name  string        `yaml:"name"`
	Files TemplateFiles `yaml:"files"`
}

type TemplateFiles struct {
	Infra        string `yaml:"infra"`
	CICD         string `yaml:"cicd"`
	ProjectFiles string `yaml:"project_files"`
}
