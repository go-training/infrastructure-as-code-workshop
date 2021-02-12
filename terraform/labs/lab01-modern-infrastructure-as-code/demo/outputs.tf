output "s3_webhost_url" {
  value = aws_s3_bucket.b.website_endpoint
}
