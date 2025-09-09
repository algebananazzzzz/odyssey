package types

type GlobalConfig struct {
	Bucket             string `json:"bucket"`
	WorkspaceKeyPrefix string `json:"workspace_key_prefix"`
	Region             string `json:"region"`
}

func (gc *GlobalConfig) IsEmpty() bool {
	if gc == nil {
		return true
	}
	return gc.Bucket == "" && gc.WorkspaceKeyPrefix == "" && gc.Region == ""
}
