resource "aws_s3_bucket" "b" {
  bucket = "foobar4567"
  acl    = "private"

  website {
    index_document = "index.html"
  }
}

resource "aws_s3_bucket_object" "index" {
  bucket       = aws_s3_bucket.b.id
  key          = "index.html"
  source       = file("${path.module}/content/index.html")
  acl          = "public-read"
  content_type = "text/html; charset=utf-8"
}
