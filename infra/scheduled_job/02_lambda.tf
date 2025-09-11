locals {
  function_name = "${var.env}-app-func-${var.project_code}"
}

module "lambda_function" {
  source             = "./modules/lambda_function"
  function_name      = local.function_name
  execution_role_arn = module.lambda_execution_role.role.arn
  deployment_package = {
    image_uri = "${aws_ecr_repository.this.repository_url}:placeholder"
  }
  ignore_deployment_package_changes = true

  depends_on = [null_resource.push_placeholder_image]
}
