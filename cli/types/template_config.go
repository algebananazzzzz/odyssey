package types

type TemplateConfig struct {
	Name         string `yaml:"name"`
	Infra        string `yaml:"infra"`
	CICD         string `yaml:"cicd"`
	ProjectFiles string `yaml:"project_files"`
}
