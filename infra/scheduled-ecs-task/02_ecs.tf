locals {
  ecs_cluster_name          = "${var.env}-app-ecscluster-${var.project_code}"
  ecs_task_defn_name        = "${var.env}-app-ecstaskdefn-${var.project_code}"
  cloudwatch_log_group_name = "${var.env}/${var.project_code}/app/ecs-logs"
}

resource "aws_ecs_cluster" "this" {
  name = local.ecs_cluster_name

  setting {
    name  = "containerInsights"
    value = "enabled"
  }
}

resource "aws_ecs_cluster_capacity_providers" "fargate" {
  cluster_name = aws_ecs_cluster.this.name

  capacity_providers = ["FARGATE"]

  default_capacity_provider_strategy {
    base              = 1
    weight            = 100
    capacity_provider = "FARGATE"
  }
}


resource "aws_cloudwatch_log_group" "this" {
  name              = local.cloudwatch_log_group_name
  retention_in_days = 7
}

resource "aws_ecs_task_definition" "this" {
  family                   = local.ecs_task_defn_name
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  cpu                      = "256"
  memory                   = "512"
  execution_role_arn       = module.ecs_execution_role.role.arn
  task_role_arn            = module.ecs_task_role.role.arn

  container_definitions = jsonencode([
    {
      name      = "app"
      image     = "${aws_ecr_repository.this.repository_url}:placeholder"
      cpu       = 256
      memory    = 512
      essential = true
      command   = ["node", "dist/index.js"]

      logConfiguration = {
        logDriver = "awslogs"
        options = {
          awslogs-group         = aws_cloudwatch_log_group.this.name
          awslogs-region        = var.aws_region
          awslogs-stream-prefix = "ecs"
        }
      }
    }
  ])

  lifecycle {
    ignore_changes = [container_definitions]
  }
}
