output "origin_bucket_name" {
  value = module.cloudfront_s3_origin.new_bucket.bucket
}

output "cloudfront_distribution_id" {
  description = "The ID of the CloudFront distribution."
  value       = module.cloudfront_s3_origin.distribution.id
}
