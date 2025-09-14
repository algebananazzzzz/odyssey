locals {
  origin_bucket_name    = "${var.env}-web-s3-${var.project_code}"
  rewrite_function_name = "${var.env}-web-cloudfrontfn-${var.project_code}-rewrite-index-html"
  use_custom_aliases    = var.custom_aliases != null
}

resource "aws_cloudfront_function" "rewrite_index" {
  name    = local.rewrite_function_name
  runtime = "cloudfront-js-2.0"
  comment = "Rewrite paths without extension to /index.html"
  publish = true
  code    = file("${path.module}/rewrite.js")
}

module "cloudfront_s3_origin" {
  source             = "./modules/cloudfront_s3_origin"
  origin_bucket_name = local.origin_bucket_name

  cloudfront_aliases = var.custom_aliases
  cloudfront_viewer_certificate = local.use_custom_aliases ? {
    acm_certificate_arn = data.aws_acm_certificate.this[0].arn
    ssl_support_method  = "sni-only"
  } : null
  route53_create_records = local.use_custom_aliases ? true : false
  route53_zone_id        = local.use_custom_aliases ? data.aws_route53_zone.public[0].id : null

  cloudfront_default_cache_behaviour = {
    allowed_methods        = ["GET", "HEAD", "OPTIONS"]
    cached_methods         = ["GET", "HEAD"]
    compress               = true
    cache_policy_id        = "658327ea-f89d-4fab-a63d-7e88639e58f6"
    viewer_protocol_policy = "redirect-to-https"
    function_association = {
      event_type   = "viewer-request"
      function_arn = aws_cloudfront_function.rewrite_index.arn
    }
  }
}
