variable "region" {
  description = "This is the cloud hosting region where your webapp will be deployed."
  default = "ap-northeast-1"
}

variable "bucket" {
  description = "default s3 bucket name."
  default = "foobar"
}

variable "profile" {
  description = "A named profile is a collection of settings and credentials that you can apply to a AWS CLI command."
  default = "demo"
}
