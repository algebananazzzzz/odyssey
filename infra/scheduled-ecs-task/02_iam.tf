locals {
  eventbridge_role_name   = "${var.env}-mgmt-iamrole-${var.project_code}-eventbridge-role"
  eventbridge_policy_name = "${var.env}-mgmt-iampolicy-${var.project_code}-eventbridge-policy"
  execution_role_name     = "${var.env}-mgmt-iamrole-${var.project_code}-execution-role"
  execution_policy_name   = "${var.env}-mgmt-iampolicy-${var.project_code}-execution-policy"
  task_role_name          = "${var.env}-mgmt-iamrole-${var.project_code}-task-role"
  task_policy_name        = "${var.env}-mgmt-iampolicy${var.project_code}-task-policy"
}

module "eventbridge_role" {
  source = "./modules/iam_role"
  name   = local.eventbridge_role_name

  assume_role_allowed_principals = [{
    identifiers = ["scheduler.amazonaws.com"]
    type        = "Service"
  }]

  custom_policy = {
    name = local.eventbridge_policy_name
    statements = {
      ecsRunTask = {
        effect = "Allow"
        actions = [
          "ecs:RunTask",
        ]
        resources = [
          aws_ecs_cluster.this.arn,
          "${aws_ecs_task_definition.this.arn_without_revision}:*",
        ]
      }
      iamPermissions = {
        effect = "Allow"
        actions = [
          "iam:PassRole"
        ]
        resources = [
          module.ecs_execution_role.role.arn,
          module.ecs_task_role.role.arn,
        ]
      }
    }
  }
}

module "ecs_execution_role" {
  source = "./modules/iam_role"
  name   = local.execution_role_name

  assume_role_allowed_principals = [{
    identifiers = ["ecs-tasks.amazonaws.com"]
    type        = "Service"
  }]

  policy_attachments = [
    "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
  ]

  custom_policy = {
    name = local.execution_policy_name
    statements = {
      # ssmPermissions = {
      #   effect = "Allow"
      #   actions = [
      #     "ssm:GetParameters",
      #     "ssm:GetParameter",
      #   ]
      #   resources = [
      #     "arn:aws:ssm:${var.region}:${data.aws_caller_identity.current.account_id}:parameter${local.ssm_parameter_prefix}/*"
      #   ]
      # }
    }
  }
}

module "ecs_task_role" {
  source = "./modules/iam_role"
  name   = local.task_role_name

  assume_role_allowed_principals = [{
    identifiers = ["ecs-tasks.amazonaws.com"]
    type        = "Service"
  }]
}


