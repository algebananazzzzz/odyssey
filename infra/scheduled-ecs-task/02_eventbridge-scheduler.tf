locals {
  schedule_name       = "${var.env}-app-schedule-${var.project_code}"
  schedule_group_name = "${var.env}-app-schedulegrp-${var.project_code}"
}


resource "aws_scheduler_schedule_group" "scheduler" {
  name = local.schedule_group_name
}

resource "aws_scheduler_schedule" "scheduler" {
  name       = local.schedule_name
  group_name = aws_scheduler_schedule_group.scheduler.name

  flexible_time_window {
    mode = "OFF"
  }

  schedule_expression          = "cron(0/30 * * * ? *)"
  schedule_expression_timezone = "Asia/Singapore"

  target {
    arn      = aws_ecs_cluster.this.arn
    role_arn = module.eventbridge_role.role.arn

    ecs_parameters {
      task_definition_arn = aws_ecs_task_definition.this.arn_without_revision
      launch_type         = "FARGATE"

      network_configuration {
        assign_public_ip = true
        security_groups  = [aws_security_group.this.id]
        subnets          = data.aws_subnets.public.ids
      }
    }

    retry_policy {
      maximum_event_age_in_seconds = 300
      maximum_retry_attempts       = 3
    }
  }
}
