# 更新 AWS 架構 (S3 Hosting)

上個步驟教大家如何建立 Infra 架構，那這單元教大家如何將使用 S3 當一個簡單的 Web Hosting。

1. 將 index.html 放入 S3 內
2. 設定 S3 當作 Web Hosting
3. 測試 S3 Hosting

## 步驟一: 建立 index.html 放入 S3 內

建立 `content/index.html` 檔案，內容如下

```html
<html>
  <body>
    <h1>Hello Terrafrom S3 Bucket From Dev</h1>
  </body>
</html>
```

修改 `s3.tf`，將 `index.html` 加入到 S3 bucket 內

```tf
resource "aws_s3_bucket_object" "index" {
  bucket       = aws_s3_bucket.b.id
  key          = "index.html"
  source       = "content/index.html"
  acl          = "public-read"
  content_type = "text/html; charset=utf-8"
}
```

其中目錄結構如下

```sh
└── demo
    ├── content
    │   └── index.html
    ├── s3.tf
    ├── terraform.tf
    ├── terraform.tfstate
    └── terraform.tfstate.backup
```

部署到 S3 Bucket 內

```sh
$ terraform apply
aws_s3_bucket.b: Refreshing state... [id=foobar4567]

An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # aws_s3_bucket_object.index will be created
  + resource "aws_s3_bucket_object" "index" {
      + acl                    = "public-read"
      + bucket                 = "foobar4567"
      + content_type           = "text/html; charset=utf-8"
      + etag                   = (known after apply)
      + force_destroy          = false
      + id                     = (known after apply)
      + key                    = "indx.html"
      + server_side_encryption = (known after apply)
      + source                 = "content/index.html"
      + storage_class          = (known after apply)
      + version_id             = (known after apply)
    }

Plan: 1 to add, 0 to change, 0 to destroy.
```

## 步驟二: 設定 S3 為 Web Hosting

修改 main.go

```tf
resource "aws_s3_bucket" "b" {
  bucket = "foobar4567"
  acl    = "private"

  website {
    index_document = "index.html"
  }
}
```

最後新增 `outputs.tf` 新增 `website_endpoint` 顯示結果

```tf
output "s3_webhost_url" {
  value = aws_s3_bucket.b.website_endpoint
}
```

執行 `terraform apply`

```sh
$ terraform apply
aws_s3_bucket.b: Refreshing state... [id=foobar4567]
aws_s3_bucket_object.index: Refreshing state... [id=indx.html]

An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:

Terraform will perform the following actions:

Plan: 0 to add, 0 to change, 0 to destroy.

Changes to Outputs:
  + s3_webhost_url = "foobar4567.s3-website-ap-northeast-1.amazonaws.com"

Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: yes


Apply complete! Resources: 0 added, 0 changed, 0 destroyed.

Outputs:

s3_webhost_url = "foobar4567.s3-website-ap-northeast-1.amazonaws.com"
```

## 步驟三: 測試 URL

透過 CURL 指令測試看看

```sh
$ curl -v foobar4567.s3-website-ap-northeast-1.amazonaws.com
*   Trying 52.219.16.80...
* TCP_NODELAY set
* Connected to foobar4567.s3-website-ap-northeast-1.amazonaws.com (52.219.16.80) port 80 (#0)
> GET / HTTP/1.1
> Host: foobar4567.s3-website-ap-northeast-1.amazonaws.com
> User-Agent: curl/7.64.1
> Accept: */*
>
< HTTP/1.1 200 OK
< x-amz-id-2: p4xTlZMjDafPDCB6JEHEglO9nWzkKSQWljZuFFaPhuqELKrAbJIbM3nkFFnBG5TM00M4iEha0FI=
< x-amz-request-id: 4F42D526789BDC9C
< Date: Fri, 12 Feb 2021 14:33:06 GMT
< Last-Modified: Fri, 12 Feb 2021 14:33:01 GMT
< ETag: "f3a63be8f363c2478ffc79f169610d36"
< Content-Type: text/html; charset=utf-8
< Content-Length: 82
< Server: AmazonS3
<
<html>
  <body>
    <h1>Hello Terrafrom S3 Bucket From Dev</h1>
  </body>
</html>
* Connection #0 to host foobar4567.s3-website-ap-northeast-1.amazonaws.com left intact
* Closing connection 0
```

## 下一個章節

=> [設定 Terrafrom 環境變數](./05-making-your-stack-configurable.md)
