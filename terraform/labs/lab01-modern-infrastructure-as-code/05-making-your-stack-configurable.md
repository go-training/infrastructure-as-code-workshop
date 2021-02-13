# 設定 Terrafrom 環境變數

大家可以看到，現在所有 `s3.tf` 的程式碼，都是直接 hardcode 的，那怎麼透過一些環境變數來動態改變設定呢？

## 步驟一: 設定變數內容

建立 `variables.tf`

```tf
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
```

接著打開 `s3.tf`:

```tf
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
```

可以看到現在可以透過 `${var.bucket}` 進行環境設定。

## 步驟二: 進行部署

```sh
$ terraform apply

An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # aws_s3_bucket.b will be created
  + resource "aws_s3_bucket" "b" {
      + acceleration_status         = (known after apply)
      + acl                         = "private"
      + arn                         = (known after apply)
      + bucket                      = (known after apply)
      + bucket_domain_name          = (known after apply)
      + bucket_regional_domain_name = (known after apply)
      + force_destroy               = false
      + hosted_zone_id              = (known after apply)
      + id                          = (known after apply)
      + region                      = (known after apply)
      + request_payer               = (known after apply)
      + website_domain              = (known after apply)
      + website_endpoint            = (known after apply)

      + versioning {
          + enabled    = (known after apply)
          + mfa_delete = (known after apply)
        }

      + website {
          + index_document = "index.html"
        }
    }

  # aws_s3_bucket_object.index will be created
  + resource "aws_s3_bucket_object" "index" {
      + acl                    = "public-read"
      + bucket                 = (known after apply)
      + content_type           = "text/html; charset=utf-8"
      + etag                   = (known after apply)
      + force_destroy          = false
      + id                     = (known after apply)
      + key                    = "index.html"
      + server_side_encryption = (known after apply)
      + source                 = "content/index.html"
      + storage_class          = (known after apply)
      + version_id             = (known after apply)
    }

  # random_pet.petname will be created
  + resource "random_pet" "petname" {
      + id        = (known after apply)
      + length    = 3
      + separator = "-"
    }

Plan: 3 to add, 0 to change, 0 to destroy.
```

我們可以自訂不同的變數內容，如果有 `dev` 及 `prod` 環境，可以建立兩個檔案 `dev.tfvars` 及 `prod.tfvars`，內容格式如下:

```tf
bucket = "example"
```

接著使用 `apply` 指令時需要加上額外的參數:

```sh
terraform apply -var-file=dev.tfvars
```

可以看到底下結果:

```sh
Terraform will perform the following actions:

  # aws_s3_bucket.b must be replaced
-/+ resource "aws_s3_bucket" "b" {
      + acceleration_status         = (known after apply)
      ~ arn                         = "arn:aws:s3:::foobar-rightly-healthy-griffon" -> (known after apply)
      ~ bucket                      = "foobar-rightly-healthy-griffon" -> "example-rightly-healthy-griffon" # forces replacement
      ~ bucket_domain_name          = "foobar-rightly-healthy-griffon.s3.amazonaws.com" -> (known after apply)
      ~ bucket_regional_domain_name = "foobar-rightly-healthy-griffon.s3.ap-northeast-1.amazonaws.com" -> (known after apply)
      ~ hosted_zone_id              = "Z2M4EHUR26P7ZW" -> (known after apply)
      ~ id                          = "foobar-rightly-healthy-griffon" -> (known after apply)
      ~ region                      = "ap-northeast-1" -> (known after apply)
      ~ request_payer               = "BucketOwner" -> (known after apply)
      - tags                        = {} -> null
      ~ website_domain              = "s3-website-ap-northeast-1.amazonaws.com" -> (known after apply)
      ~ website_endpoint            = "foobar-rightly-healthy-griffon.s3-website-ap-northeast-1.amazonaws.com" -> (known after apply)
        # (2 unchanged attributes hidden)

      ~ versioning {
          ~ enabled    = false -> (known after apply)
          ~ mfa_delete = false -> (known after apply)
        }

      ~ website {
            # (1 unchanged attribute hidden)
        }
    }

  # aws_s3_bucket_object.index must be replaced
-/+ resource "aws_s3_bucket_object" "index" {
      ~ bucket                 = "foobar-rightly-healthy-griffon" -> (known after apply) # forces replacement
      ~ etag                   = "f3a63be8f363c2478ffc79f169610d36" -> (known after apply)
      ~ id                     = "index.html" -> (known after apply)
      - metadata               = {} -> null
      + server_side_encryption = (known after apply)
      ~ storage_class          = "STANDARD" -> (known after apply)
      - tags                   = {} -> null
      + version_id             = (known after apply)
        # (5 unchanged attributes hidden)
    }

Plan: 2 to add, 0 to change, 2 to destroy.

Changes to Outputs:
  ~ s3_webhost_url = "foobar-rightly-healthy-griffon.s3-website-ap-northeast-1.amazonaws.com" -> (known after apply)
```

## 步驟三: 讀取更多檔案

如果每個 html 檔案都要寫一個 config，這樣相當麻煩，透過底下方式可以讀取單一目錄直接設定

```tf
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
```

透過 terraform 的 fileset 讀取全部資料。最後執行部署

```sh
$ terraform apply -var-file=dev.tfvars
random_pet.petname: Refreshing state... [id=rightly-healthy-griffon]
aws_s3_bucket_object.index: Refreshing state... [id=index.html]
aws_s3_bucket.b: Refreshing state... [id=example-rightly-healthy-griffon]

An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  + create
  - destroy

Terraform will perform the following actions:

  # aws_s3_bucket_object.content["about.html"] will be created
  + resource "aws_s3_bucket_object" "content" {
      + acl                    = "public-read"
      + bucket                 = "example-rightly-healthy-griffon"
      + content_type           = "text/html; charset=utf-8"
      + etag                   = "c00772d098b074161c79200aee996d0f"
      + force_destroy          = false
      + id                     = (known after apply)
      + key                    = "about.html"
      + server_side_encryption = (known after apply)
      + source                 = "content/about.html"
      + storage_class          = (known after apply)
      + version_id             = (known after apply)
    }

  # aws_s3_bucket_object.content["index.html"] will be created
  + resource "aws_s3_bucket_object" "content" {
      + acl                    = "public-read"
      + bucket                 = "example-rightly-healthy-griffon"
      + content_type           = "text/html; charset=utf-8"
      + etag                   = "f3a63be8f363c2478ffc79f169610d36"
      + force_destroy          = false
      + id                     = (known after apply)
      + key                    = "index.html"
      + server_side_encryption = (known after apply)
      + source                 = "content/index.html"
      + storage_class          = (known after apply)
      + version_id             = (known after apply)
    }

  # aws_s3_bucket_object.index will be destroyed
  - resource "aws_s3_bucket_object" "index" {
      - acl           = "public-read" -> null
      - bucket        = "example-rightly-healthy-griffon" -> null
      - content_type  = "text/html; charset=utf-8" -> null
      - etag          = "f3a63be8f363c2478ffc79f169610d36" -> null
      - force_destroy = false -> null
      - id            = "index.html" -> null
      - key           = "index.html" -> null
      - metadata      = {} -> null
      - source        = "content/index.html" -> null
      - storage_class = "STANDARD" -> null
      - tags          = {} -> null
    }

Plan: 2 to add, 0 to change, 1 to destroy.
```

這邊要注意 `content_type` 的設定，如果目錄底下有很多不同格式的檔案，這樣需要分別設定 content type，依照其他方式分類，像是 `images/png` 或 `text/html` 區分開來設定。這邊如果是換作 `Pulumi` 就可以用程式方式來透過副檔名來自動讀取 type，底下是 Go 語言範例:

```go
mime.TypeByExtension(path.Ext(filepath.Join(site, name)))
```
