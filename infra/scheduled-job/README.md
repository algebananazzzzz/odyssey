<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >=1.5.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~>5.31.0 |

## Resources

| Name | Type |
|------|------|
| [aws_ecr_repository.this](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecr_repository) | resource |
| [aws_scheduler_schedule.scheduler](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/scheduler_schedule) | resource |
| [aws_scheduler_schedule_group.scheduler](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/scheduler_schedule_group) | resource |
| [null_resource.push_placeholder_image](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_aws_region"></a> [aws\_region](#input\_aws\_region) | The AWS region where all resources will be created. | `any` | n/a | yes |
| <a name="input_env"></a> [env](#input\_env) | The deployment environment (e.g., dev, staging, prod). | `any` | n/a | yes |
| <a name="input_project_code"></a> [project\_code](#input\_project\_code) | A short project code or identifier used to tag and name resources. | `any` | n/a | yes |
<!-- END_TF_DOCS -->