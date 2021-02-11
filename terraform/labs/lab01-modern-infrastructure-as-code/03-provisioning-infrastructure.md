# 初始化 Pulumi 架構 (建立 S3 Bucket)

## 步驟一: 建立新的 S3 Bucket

建立 `s3.tf` 檔案，並且寫入底下資料

```tf
resource "aws_s3_bucket" "b" {
  bucket = "foobar4567"
  acl    = "private"
}
```

## 步驟二: 執行 Terraform 預覽

透過底下指令可以直接預覽每個操作步驟所做的改變:

```sh
terraform apply
```

可以看到底下預覽:

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
      + bucket                      = "foobar1234"
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
    }

Plan: 1 to add, 0 to change, 0 to destroy.
```

可以看到更詳細的建立步驟及權限，在此步驟可以詳細知道 Terraform 會怎麼設定 AWS 架構，透過此預覽方式避免人為操作失誤。

## 步驟三: 執行部署

看完上面的預覽，我們最後就直接鍵入 `yes` 執行:

```sh
Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: yes

aws_s3_bucket.b: Creating...
aws_s3_bucket.b: Creation complete after 5s [id=foobar4567]

Apply complete! Resources: 1 added, 0 changed, 0 destroyed.
```

接著你可以看到在目錄底下多了兩個檔案 `terraform.tfstate` 跟 `terraform.tfstate.backup`，請注意務必將 `.tfstate` 放入 `.gitignore` 列表內，不要將此檔案放到 git 版本控制內。
