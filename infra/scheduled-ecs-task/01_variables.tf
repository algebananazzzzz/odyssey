variable "env" {
  description = "The deployment environment (e.g., dev, staging, prod)."
}

variable "aws_region" {
  description = "The AWS region where all resources will be created."
}

variable "project_code" {
  description = "A short project code or identifier used to tag and name resources."
}
