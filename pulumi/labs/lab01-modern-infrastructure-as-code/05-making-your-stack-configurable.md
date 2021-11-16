# 設定 Pulumi Stack 環境變數

大家可以看到，現在所有 `main.go` 的程式碼，都是直接 hardcode 的，那怎麼透過一些環境變數來動態改變設定呢？這時候可以透過 pulumi config 指令來調整喔，底下來看看怎麼實作，假設我們要讀取的 index.html 放在其他目錄底下，該怎麼動態調整？

## 步驟一: 撰寫讀取 Config 函式

```go
func getEnv(ctx *pulumi.Context, key string, fallback ...string) string {
  if value, ok := ctx.GetConfig(key); ok {
    return value
  }

  if len(fallback) > 0 {
    return fallback[0]
  }

  return ""
}
```

pulumi 的 context 內有一個讀取環境變數函式叫 `GetConfig`，接著我們在設計一個 fallback 當作 default 回傳值。底下設定一個變數 `s3:siteDir`

```sh
pulumi config set s3:siteDir production
```

打開 `Pulumi.dev.yaml` 可以看到

```yaml
config:
  aws:profile: demo
  aws:region: ap-northeast-1
  s3:siteDir: production
```

接著將程式碼改成如下:

```go
    site := getEnv(ctx, "s3:siteDir", "content")
    index := path.Join(site, "index.html")
    _, err = s3.NewBucketObject(ctx, "index.html", &s3.BucketObjectArgs{
      Bucket:      bucket.Bucket,
      Source:      pulumi.NewFileAsset(index),
      Acl:         pulumi.String("public-read"),
      ContentType: pulumi.String(mime.TypeByExtension(path.Ext(index))),
    })
```

## 步驟二: 進行部署

```sh
$ pulumi up
Previewing update (dev)

View Live: https://app.pulumi.com/appleboy/demo/dev/previews/d76d2f9b-16c8-4bfd-820d-d5368d29f592

     Type                    Name        Plan       Info
     pulumi:pulumi:Stack     demo-dev
 ~   └─ aws:s3:BucketObject  index.html  update     [diff: ~source]

Resources:
    ~ 1 to update
    2 unchanged

Do you want to perform this update? details
  pulumi:pulumi:Stack: (same)
    [urn=urn:pulumi:dev::demo::pulumi:pulumi:Stack::demo-dev]
    ~ aws:s3/bucketObject:BucketObject: (update)
        [id=index.html]
        [urn=urn:pulumi:dev::demo::aws:s3/bucketObject:BucketObject::index.html]
      - source: asset(file:77aab46) { content/index.html }
      + source: asset(file:01c09f4) { production/index.html }
```

可以看到 source 會被換成 `production/index.html`

## 步驟三: 讀取更多檔案

整個 Web 專案肯定不止一個檔案，所以再來改一下原本的讀取檔案列表流程

```go
    site := getEnv(ctx, "s3:siteDir", "content")
    files, err := ioutil.ReadDir(site)
    if err != nil {
      return err
    }

    for _, item := range files {
      name := item.Name()
      if _, err = s3.NewBucketObject(ctx, name, &s3.BucketObjectArgs{
        Bucket:      bucket.Bucket,
        Source:      pulumi.NewFileAsset(filepath.Join(site, name)),
        Acl:         pulumi.String("public-read"),
        ContentType: pulumi.String(mime.TypeByExtension(path.Ext(filepath.Join(site, name)))),
      }); err != nil {
        return err
      }
    }
```

執行部署

```sh
     Type                    Name        Status      Info
     pulumi:pulumi:Stack     demo-dev
 +   ├─ aws:s3:BucketObject  about.html  created
 ~   └─ aws:s3:BucketObject  index.html  updated     [diff: ~source]

Outputs:
    bucketEndpoint: "foobar-1234.s3-website-ap-northeast-1.amazonaws.com"
    bucketID      : "foobar-1234"
    bucketName    : "foobar-1234"

Resources:
    + 1 created
    ~ 1 updated
    2 changes. 2 unchanged

Duration: 9s
```

完整程式碼如下:

```go
package main

import (
  "io/ioutil"
  "mime"
  "path"
  "path/filepath"

  "github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
  "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
  pulumi.Run(func(ctx *pulumi.Context) error {
    // Create an AWS resource (S3 Bucket)
    bucket, err := s3.NewBucket(ctx, "my-bucket", &s3.BucketArgs{
      Bucket: pulumi.String("foobar-1234"),
      Website: s3.BucketWebsiteArgs{
        IndexDocument: pulumi.String("index.html"),
      },
    })
    if err != nil {
      return err
    }

    site := getEnv(ctx, "s3:siteDir", "content")
    files, err := ioutil.ReadDir(site)
    if err != nil {
      return err
    }

    for _, item := range files {
      name := item.Name()
      if _, err = s3.NewBucketObject(ctx, name, &s3.BucketObjectArgs{
        Bucket:      bucket.Bucket,
        Source:      pulumi.NewFileAsset(filepath.Join(site, name)),
        Acl:         pulumi.String("public-read"),
        ContentType: pulumi.String(mime.TypeByExtension(path.Ext(filepath.Join(site, name)))),
      }); err != nil {
        return err
      }
    }

    // Export the name of the bucket
    ctx.Export("bucketID", bucket.ID())
    ctx.Export("bucketName", bucket.Bucket)
    ctx.Export("bucketEndpoint", bucket.WebsiteEndpoint)

    return nil
  })
}

func getEnv(ctx *pulumi.Context, key string, fallback ...string) string {
  if value, ok := ctx.GetConfig(key); ok {
    return value
  }

  if len(fallback) > 0 {
    return fallback[0]
  }

  return ""
}
```

## 下一個章節

=> [建立第二個 Pulumi Stack 環境](./06-create-new-stack.md)
