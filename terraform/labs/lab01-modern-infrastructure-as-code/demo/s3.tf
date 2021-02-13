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

resource "aws_s3_bucket_object" "index" {
  bucket       = aws_s3_bucket.b.id
  key          = "index.html"
  source       = "content/index.html"
  acl          = "public-read"
  content_type = "text/html; charset=utf-8"
}
