package types

type ProjectConfig struct {
	Template  string
	Code      string
	Type      string
	Directory string
	Workflow  string
	Variables map[string]string `json:"variables" yaml:"variables"`
	Corrupted bool              `json:"-" yaml:"-"`
}

// IsEmpty reports whether the project config is uninitialized.
func (pc *ProjectConfig) IsEmpty() bool {
	if pc == nil {
		return true
	}
	return pc.Template == "" && pc.Code == "" && pc.Type == "" && pc.Directory == "" && pc.Workflow == "" && len(pc.Variables) == 0
}
