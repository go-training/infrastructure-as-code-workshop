# 設定 AWS 環境

在使用 Terraform 之前要先把 AWS 環境建立好

## 步驟一: 設定 AWS Region

請先將 AWS 環境設定完畢，請用 `aws configure` 完成 profile 設定

```sh
aws configure --profile demo
```

## 步驟二: 在 Terraform 設定 AWS Profile

打開 demo 目錄內的 `terraform.tf` 檔案，在內容後面補上

```tf
provider "aws" {
  profile = "demo"
  region  = "ap-northeast-1"
}
```
