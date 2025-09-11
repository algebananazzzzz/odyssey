locals {
  security_group_name = "${var.env}-app-sg-${var.project_code}"
}

# Get the default VPC
data "aws_vpc" "default" {
  filter {
    name   = "tag:Name"
    values = ["com-all-vpc-vpcsg"]
  }
}

# Get the public subnets in the default VPC
data "aws_subnets" "public" {
  filter {
    name   = "vpc-id"
    values = [data.aws_vpc.default.id]
  }

  filter {
    name   = "tag:Environment"
    values = [var.env]
  }

  filter {
    name   = "tag:Zone"
    values = ["app"]
  }
}

# Create a security group in the default VPC
resource "aws_security_group" "this" {
  name        = local.security_group_name
  description = "Security group for ECS task"
  vpc_id      = data.aws_vpc.default.id

  # Outbound: Allow all
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = local.security_group_name
  }
}
