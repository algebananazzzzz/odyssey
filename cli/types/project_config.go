package types

type ProjectConfig struct {
	Code         string
	Type         string
	Environments int
}

func (pc *ProjectConfig) IsEmpty() bool {
	if pc == nil {
		return true
	}
	return pc.Code == "" && pc.Type == "" && pc.Environments == 0
}
