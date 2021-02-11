# 建立第二個 Pulumi Stack 環境

在 Pulumi 可以很簡單的建立多種環境，像是 Testing 或 Production，只要將動態變數抽出來設定成 config 即可。底下來看看怎麼建立全先的環境，這步驟在 Pulumi 叫做 Stack。前面已經建立一個 dev 環境，現在我們要建立一個全新環境來部署 Testing 或 Production 該如何做呢？

## 步驟一: 建立全新 Stack 環境

透過 pulumi stack 可以建立全新環境

```sh
$ pulumi stack ls
NAME  LAST UPDATE   RESOURCE COUNT  URL
dev*  1 minute ago  5               https://app.pulumi.com/appleboy/demo/dev
```

建立 stack

```sh
$ pulumi stack init prod
Created stack 'prod'
$ pulumi stack ls
NAME   LAST UPDATE   RESOURCE COUNT  URL
dev    1 minute ago  5               https://app.pulumi.com/appleboy/demo/dev
prod*  n/a           n/a             https://app.pulumi.com/appleboy/demo/prod
```

設定參數

```sh
pulumi config set s3:siteDir www
pulumi config set aws:profile demo
pulumi config set aws:region ap-northeast-1
```

## 步驟二: 建立 www 內容

建立 `content/www` 目錄，一樣放上 index.htm + about.html

```html
<html>
  <body>
    <h1>Hello Pulumi S3 Bucket From New Stack</h1>
  </body>
</html>
```

about.html

```html
<html>
  <body>
    <h1>About us From New Stack</h1>
  </body>
</html>
```

## 步驟三: 部署 New Stack

先看看 Preview 結果

```sh
$ pulumi up
Previewing update (prod)

View Live: https://app.pulumi.com/appleboy/demo/prod/previews/3b85a340-0e71-455e-9b96-48dc38538d18

     Type                    Name        Plan
 +   pulumi:pulumi:Stack     demo-prod   create
 +   ├─ aws:s3:Bucket        my-bucket   create
 +   ├─ aws:s3:BucketObject  index.html  create
 +   └─ aws:s3:BucketObject  about.html  create

Resources:
    + 4 to create

Do you want to perform this update? details
+ pulumi:pulumi:Stack: (create)
    [urn=urn:pulumi:prod::demo::pulumi:pulumi:Stack::demo-prod]
    + aws:s3/bucket:Bucket: (create)
        [urn=urn:pulumi:prod::demo::aws:s3/bucket:Bucket::my-bucket]
        acl         : "private"
        bucket      : "my-bucket-ba8088c"
        forceDestroy: false
        website     : {
            indexDocument: "index.html"
        }
    + aws:s3/bucketObject:BucketObject: (create)
        [urn=urn:pulumi:prod::demo::aws:s3/bucketObject:BucketObject::index.html]
        acl         : "public-read"
        bucket      : "my-bucket-ba8088c"
        contentType : "text/html; charset=utf-8"
        forceDestroy: false
        key         : "index.html"
        source      : asset(file:460188b) { www/index.html }
    + aws:s3/bucketObject:BucketObject: (create)
        [urn=urn:pulumi:prod::demo::aws:s3/bucketObject:BucketObject::about.html]
        acl         : "public-read"
        bucket      : "my-bucket-ba8088c"
        contentType : "text/html; charset=utf-8"
        forceDestroy: false
        key         : "about.html"
        source      : asset(file:376c42a) { www/about.html }
```

如果看起來沒問題，就可以直接執行了

```sh
Updating (prod)

View Live: https://app.pulumi.com/appleboy/demo/prod/updates/1

     Type                    Name        Status
 +   pulumi:pulumi:Stack     demo-prod   created
 +   ├─ aws:s3:Bucket        my-bucket   created
 +   ├─ aws:s3:BucketObject  about.html  created
 +   └─ aws:s3:BucketObject  index.html  created

Outputs:
    bucketEndpoint: "my-bucket-a7044ab.s3-website-ap-northeast-1.amazonaws.com"
    bucketID      : "my-bucket-a7044ab"
    bucketName    : "my-bucket-a7044ab"

Resources:
    + 4 created

Duration: 18s
```

最後用 curl 執行看看

```sh
$ curl -v $(pulumi stack output bucketEndpoint)
*   Trying 52.219.8.20...
* TCP_NODELAY set
* Connected to my-bucket-a7044ab.s3-website-ap-northeast-1.amazonaws.com (52.219.8.20) port 80 (#0)
> GET / HTTP/1.1
> Host: my-bucket-a7044ab.s3-website-ap-northeast-1.amazonaws.com
> User-Agent: curl/7.64.1
> Accept: */*
>
< HTTP/1.1 200 OK
< x-amz-id-2: oGxc+rLPi3kLOZslMsOmJqPY/WGeMoxX9sXJDRj4wlJlGVq+7pMx3ers71jxnDiDkeM9JRrd+T8=
< x-amz-request-id: 528235DDFF40F365
< Date: Thu, 11 Feb 2021 04:49:21 GMT
< Last-Modified: Thu, 11 Feb 2021 04:48:41 GMT
< ETag: "ae41d1b3f0aeef6a490e1b2edc74d2b5"
< Content-Type: text/html; charset=utf-8
< Content-Length: 85
< Server: AmazonS3
<
<html>
  <body>
    <h1>Hello Pulumi S3 Bucket From New Stack</h1>
  </body>
</html>
* Connection #0 to host my-bucket-a7044ab.s3-website-ap-northeast-1.amazonaws.com left intact
* Closing connection 0
```
