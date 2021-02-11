# 建立新專案

## 步驟一: 安裝 Pulumi CLI 工具

首先你要在自己電腦安裝上 Pulumi CLI 工具，請參考[官方網站][1]，根據您的作業環境有不同的安裝方式，底下以 Mac 環境為主

[1]:https://www.pulumi.com/docs/get-started/install/

```sh
brew install pulumi
```

透過 brew 即可安裝成功，那升級工具透過底下即可

```sh
brew upgrade pulumi
```

或者您沒有使用 brew，也可以透過 curl 安裝

```sh
curl -fsSL https://get.pulumi.com | sh
```

測試 CLI 指令

```sh
$ pulumi version
v2.20.0
```

有看到版本資訊就是安裝成功了

## 步驟二: 初始化專案

透過 `pulumi new -h` 可以看到說明

```sh
Usage:
  pulumi new [template|url] [flags]

Flags:
  -c, --config stringArray        Config to save
      --config-path               Config keys contain a path to a property in a map or list to set
  -d, --description string        The project description; if not specified, a prompt will request it
      --dir string                The location to place the generated project; if not specified, the current directory is used
  -f, --force                     Forces content to be generated even if it would change existing files
  -g, --generate-only             Generate the project only; do not create a stack, save config, or install dependencies
  -h, --help                      help for new
  -n, --name string               The project name; if not specified, a prompt will request it
  -o, --offline                   Use locally cached templates without making any network requests
      --secrets-provider string   The type of the provider that should be used to encrypt and decrypt secrets (possible choices: default, passphrase, awskms, azurekeyvault, gcpkms, hashivault) (default "default")
  -s, --stack string              The stack name; either an existing stack or stack to create; if not specified, a prompt will request it
  -y, --yes                       Skip prompts and proceed with default values
```

可以選擇的 Template 超多，那我們這次用 AWS 搭配 Go 語言的 Temaplate 當作範例

```sh
$ pulumi new aws-go --dir demo
This command will walk you through creating a new Pulumi project.

Enter a value or leave blank to accept the (default), and press <ENTER>.
Press ^C at any time to quit.

project name: (demo)
project description: (A minimal AWS Go Pulumi program)
Created project 'demo'

Please enter your desired stack name.
To create a stack in an organization, use the format <org-name>/<stack-name> (e.g. `acmecorp/dev`).
stack name: (dev)
Created stack 'dev'

aws:region: The AWS region to deploy into: (us-east-1) ap-northeast-1
Saved config

Installing dependencies...

Finished installing dependencies

Your new project is ready to go! ✨

To perform an initial deployment, run 'cd demo', then, run 'pulumi up'
```

## 步驟三: 檢查專案目錄結構

```sh
└── demo
    ├── Pulumi.dev.yaml
    ├── Pulumi.yaml
    ├── go.mod
    ├── go.sum
    └── main.go
```

其中 `main.go` 就是主程式檔案

```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an AWS resource (S3 Bucket)
		bucket, err := s3.NewBucket(ctx, "my-bucket", nil)
		if err != nil {
			return err
		}

		// Export the name of the bucket
		ctx.Export("bucketName", bucket.ID())
		return nil
	})
}
```

## 下一個章節

=> [設定 AWS 環境](./02-configuring-aws.md)
