# 更新 Infra 架構

上個步驟教大家如何建立 Infra 架構，那這單元教大家如何將使用 S3 當一個簡單的 Web Hosting。

1. 將 index.html 放入 S3 內
2. 設定 S3 當作 Web Hosting
3. 測試 S3 Hosting

## 步驟一: 建立 index.html 放入 S3 內

建立 `content/index.html` 檔案，內容如下

```html
<html>
  <body>
    <h1>Hello Pulumi S3 Bucket</h1>
  </body>
</html>
```

修改 `main.go`，將 `index.html` 加入到 S3 bucket 內

```go
		index := path.Join("content", "index.html")
		_, err = s3.NewBucketObject(ctx, "index.html", &s3.BucketObjectArgs{
			Bucket: bucket.Bucket,
			Source: pulumi.NewFileAsset(index),
		})

		if err != nil {
			return err
		}
```

其中目錄結構如下

```sh
├── demo
│   ├── Pulumi.dev.yaml
│   ├── Pulumi.yaml
│   ├── content
│   │   └── index.html
│   ├── go.mod
│   ├── go.sum
│   └── main.go
```

部署到 S3 Bucket 內

```sh
$ pulumi up
Previewing update (dev)

View Live: https://app.pulumi.com/appleboy/demo/dev/previews/a0ac1b69-06b9-4109-800d-20618b36e5c8

     Type                    Name        Plan
     pulumi:pulumi:Stack     demo-dev
 +   └─ aws:s3:BucketObject  index.html  create

Resources:
    + 1 to create
    2 unchanged

Do you want to perform this update? details
  pulumi:pulumi:Stack: (same)
    [urn=urn:pulumi:dev::demo::pulumi:pulumi:Stack::demo-dev]
    + aws:s3/bucketObject:BucketObject: (create)
        [urn=urn:pulumi:dev::demo::aws:s3/bucketObject:BucketObject::index.html]
        acl         : "private"
        bucket      : "foobar-1234"
        forceDestroy: false
        key         : "index.html"
        source      : asset(file:77aab46) { content/index.html }
```

## 步驟二: 設定 S3 為 Web Hosting

修改 main.go

```go
		bucket, err := s3.NewBucket(ctx, "my-bucket", &s3.BucketArgs{
			Bucket: pulumi.String("foobar-1234"),
			Website: s3.BucketWebsiteArgs{
				IndexDocument: pulumi.String("index.html"),
			},
		})

		index := path.Join("content", "index.html")
		_, err = s3.NewBucketObject(ctx, "index.html", &s3.BucketObjectArgs{
			Bucket:      bucket.Bucket,
			Source:      pulumi.NewFileAsset(index),
			Acl:         pulumi.String("public-read"),
			ContentType: pulumi.String(mime.TypeByExtension(path.Ext(index))),
		})
```

最後設定輸出 URL:

```go
ctx.Export("bucketEndpoint", bucket.WebsiteEndpoint)
```

最後完整程式碼如下:

```go
package main

import (
	"mime"
	"path"

	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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

		index := path.Join("content", "index.html")
		_, err = s3.NewBucketObject(ctx, "index.html", &s3.BucketObjectArgs{
			Bucket:      bucket.Bucket,
			Source:      pulumi.NewFileAsset(index),
			Acl:         pulumi.String("public-read"),
			ContentType: pulumi.String(mime.TypeByExtension(path.Ext(index))),
		})

		if err != nil {
			return err
		}

		// Export the name of the bucket
		ctx.Export("bucketID", bucket.ID())
		ctx.Export("bucketName", bucket.Bucket)
		ctx.Export("bucketEndpoint", bucket.WebsiteEndpoint)

		return nil
	})
}
```

執行 pulumi up

```sh
Previewing update (dev)

View Live: https://app.pulumi.com/appleboy/demo/dev/previews/edbaefca-f723-4ac5-aabd-7cb638636612

     Type                    Name        Plan       Info
     pulumi:pulumi:Stack     demo-dev
 ~   ├─ aws:s3:Bucket        my-bucket   update     [diff: +website]
 ~   └─ aws:s3:BucketObject  index.html  update     [diff: ~acl,contentType]

Outputs:
  + bucketEndpoint: output<string>

Resources:
    ~ 2 to update
    1 unchanged

Do you want to perform this update? details
  pulumi:pulumi:Stack: (same)
    [urn=urn:pulumi:dev::demo::pulumi:pulumi:Stack::demo-dev]
    ~ aws:s3/bucket:Bucket: (update)
        [id=foobar-1234]
        [urn=urn:pulumi:dev::demo::aws:s3/bucket:Bucket::my-bucket]
      + website: {
          + indexDocument: "index.html"
        }
    --outputs:--
  + bucketEndpoint: output<string>
    ~ aws:s3/bucketObject:BucketObject: (update)
        [id=index.html]
        [urn=urn:pulumi:dev::demo::aws:s3/bucketObject:BucketObject::index.html]
      ~ acl        : "private" => "public-read"
      ~ contentType: "binary/octet-stream" => "text/html; charset=utf-8"
Do you want to perform this update? yes
Updating (dev)

View Live: https://app.pulumi.com/appleboy/demo/dev/updates/7

     Type                    Name        Status      Info
     pulumi:pulumi:Stack     demo-dev
 ~   ├─ aws:s3:Bucket        my-bucket   updated     [diff: +website]
 ~   └─ aws:s3:BucketObject  index.html  updated     [diff: ~acl,contentType]

Outputs:
  + bucketEndpoint: "foobar-1234.s3-website-ap-northeast-1.amazonaws.com"
    bucketID      : "foobar-1234"
    bucketName    : "foobar-1234"

Resources:
    ~ 2 updated
    1 unchanged

Duration: 13s
```

## 步驟三: 測試 URL

透過底下指令可以拿到 S3 的 URL:

```sh
pulumi stack output bucketEndpoint
```

透過 CURL 指令測試看看

```sh
$ curl -v $(pulumi stack output bucketEndpoint)
*   Trying 52.219.16.96...
* TCP_NODELAY set
* Connected to foobar-1234.s3-website-ap-northeast-1.amazonaws.com (52.219.16.96) port 80 (#0)
> GET / HTTP/1.1
> Host: foobar-1234.s3-website-ap-northeast-1.amazonaws.com
> User-Agent: curl/7.64.1
> Accept: */*
>
< HTTP/1.1 200 OK
< x-amz-id-2: 0NrZfFxZNOs+toz0/86FiASG+MyQE6f+KbKNi4wzcDtmn5mTnQoxupVybR464X8Oi6HDMjSU+i8=
< x-amz-request-id: A55BD2534EDC94A9
< Date: Thu, 11 Feb 2021 03:30:42 GMT
< Last-Modified: Thu, 11 Feb 2021 03:29:14 GMT
< ETag: "46e94ba24774d0c4a768f9461e6b9806"
< Content-Type: text/html; charset=utf-8
< Content-Length: 70
< Server: AmazonS3
<
<html>
  <body>
    <h1>Hello Pulumi S3 Bucket</h1>
  </body>
</html>
* Connection #0 to host foobar-1234.s3-website-ap-northeast-1.amazonaws.com left intact
```
