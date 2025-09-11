locals {
  lambda_execution_role_name   = "${var.env}-mgmt-iamrole-${var.project_code}-function"
  lambda_execution_policy_name = "${var.env}-mgmt-iampolicy-${var.project_code}-function"

  schedule_execution_role_name   = "${var.env}-mgmt-iamrole-${var.project_code}-scheduler"
  schedule_execution_policy_name = "${var.env}-mgmt-iampolicy-${var.project_code}-scheduler"
}

module "lambda_execution_role" {
  source = "./modules/iam_role"
  name   = local.lambda_execution_role_name
  policy_attachments = [
    "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  ]
}

module "scheduler_execution_role" {
  source = "./modules/iam_role"
  name   = local.schedule_execution_role_name
  assume_role_allowed_principals = [{
    type        = "Service"
    identifiers = ["scheduler.amazonaws.com"]
  }]
  custom_policy = {
    name = local.schedule_execution_policy_name
    statements = {
      lambdaPermissions = {
        effect = "Allow"
        actions = [
          "lambda:InvokeFunction",
        ]
        resources = [module.lambda_function.function.arn]
      }
    }
  }
}

