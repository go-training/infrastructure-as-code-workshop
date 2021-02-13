resource "random_pet" "petname" {
  length    = 3
  separator = "-"
}

resource "aws_s3_bucket" "b" {
  bucket = "${var.bucket}-${random_pet.petname.id}"
  acl    = "private"

  website {
    index_document = "index.html"
  }
}

resource "aws_s3_bucket_object" "content" {
  bucket       = aws_s3_bucket.b.id
  acl          = "public-read"
  content_type = "text/html; charset=utf-8"

  for_each = fileset("content/", "*")
  key    = each.value
  source = "content/${each.value}"
  # etag makes the file update when it changes; see https://stackoverflow.com/questions/56107258/terraform-upload-file-to-s3-on-every-apply
  etag   = filemd5("content/${each.value}")
}
