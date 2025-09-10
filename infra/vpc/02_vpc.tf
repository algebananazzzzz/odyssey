module "vpc" {
  source       = "./vpc"
  aws_region   = var.aws_region
  project_code = var.project_code
}
