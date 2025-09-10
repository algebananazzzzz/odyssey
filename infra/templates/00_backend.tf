terraform {
  backend "s3" {
    bucket               = "{{.GlobalConfig.Bucket}}"
    key                  = "{{.ProjectConfig.Code}}"
    workspace_key_prefix = "{{.GlobalConfig.WorkspaceKeyPrefix}}"
    region               = "{{.GlobalConfig.Region}}"
  }
}
