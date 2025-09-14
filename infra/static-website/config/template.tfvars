# Required Variables
# -------------------------
# The target region to which the resources will be deployed.
aws_region = "{{.GlobalConfig.Region}}"
# The code name of the project used for naming convention.
project_code = "{{.ProjectConfig.Code}}"

# Optional Variables
# -------------------------
# If you want to use a custom domain, set this variable to your desired domain name.
# custom_aliases = ["site.domain.com"]

# The domain name for the ACM certificate and Route 53 hosted zone.
# acm_certificate_domain = "domain.com"

# The name of the Route 53 hosted zone.
# route53_zone_name = "domain.com"
