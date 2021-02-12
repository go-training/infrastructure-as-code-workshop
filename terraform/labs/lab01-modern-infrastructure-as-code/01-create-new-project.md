# 建立新專案

## 步驟一: 安裝 Terraform 工具

首先你要在自己電腦安裝上 Pulumi CLI 工具，請參考[官方網站][1]，根據您的作業環境有不同的安裝方式，底下以 Mac 環境為主

[1]:https://learn.hashicorp.com/tutorials/terraform/install-cli

```sh
brew tap hashicorp/tap
brew install hashicorp/tap/terraform
```

透過 brew 即可安裝成功，那升級工具透過底下即可

```sh
brew upgrade hashicorp/tap/terraform
```

完成後，可以直接安裝自動完成 (tab completion)

```sh
terraform -install-autocomplete
```

測試 CLI 指令

```sh
$ terraform -version
Terraform v0.14.6
```

有看到版本資訊就是安裝成功了

## 步驟二: 初始化專案

建立 `demo` 專案目錄，並且新增 `terraform.tf` 檔案，內容如下:

```tf
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 2.70"
    }
  }
}
```

切換到 demo 目錄，並執行 terraform init 初始化即可

```sh
cd demo && terraform init
```

## 下一個章節

=> [設定 AWS 環境](./02-configuring-aws.md)
