# 刪除 Pulumi Stack 環境

最後步驟就是要學習怎麼一鍵刪除整個 Infrastructure 環境。現在我們已經建立兩個 Stack 環境，該怎麼移除？

## 步驟一: 刪除所有資源

用底下指令可以刪除全部資源

```sh
Previewing destroy (prod)

View Live: https://app.pulumi.com/appleboy/demo/prod/previews/92f9c4a4-f4a9-464d-be27-5040aff295ae

     Type                    Name        Plan
 -   pulumi:pulumi:Stack     demo-prod   delete
 -   ├─ aws:s3:BucketObject  about.html  delete
 -   ├─ aws:s3:BucketObject  index.html  delete
 -   └─ aws:s3:Bucket        my-bucket   delete

Outputs:
  - bucketEndpoint: "my-bucket-a7044ab.s3-website-ap-northeast-1.amazonaws.com"
  - bucketID      : "my-bucket-a7044ab"
  - bucketName    : "my-bucket-a7044ab"

Resources:
    - 4 to delete

Do you want to perform this destroy? details
- aws:s3/bucketObject:BucketObject: (delete)
    [id=about.html]
    [urn=urn:pulumi:prod::demo::aws:s3/bucketObject:BucketObject::about.html]
- aws:s3/bucketObject:BucketObject: (delete)
    [id=index.html]
    [urn=urn:pulumi:prod::demo::aws:s3/bucketObject:BucketObject::index.html]
- aws:s3/bucket:Bucket: (delete)
    [id=my-bucket-a7044ab]
    [urn=urn:pulumi:prod::demo::aws:s3/bucket:Bucket::my-bucket]
- pulumi:pulumi:Stack: (delete)
    [urn=urn:pulumi:prod::demo::pulumi:pulumi:Stack::demo-prod]
    --outputs:--
  - bucketEndpoint: "my-bucket-a7044ab.s3-website-ap-northeast-1.amazonaws.com"
  - bucketID      : "my-bucket-a7044ab"
  - bucketName    : "my-bucket-a7044ab"
```

選擇 `yse` 移除所以資源

```sh
Destroying (prod)

View Live: https://app.pulumi.com/appleboy/demo/prod/updates/2

     Type                    Name        Status
 -   pulumi:pulumi:Stack     demo-prod   deleted
 -   ├─ aws:s3:BucketObject  index.html  deleted
 -   ├─ aws:s3:BucketObject  about.html  deleted
 -   └─ aws:s3:Bucket        my-bucket   deleted

Outputs:
  - bucketEndpoint: "my-bucket-a7044ab.s3-website-ap-northeast-1.amazonaws.com"
  - bucketID      : "my-bucket-a7044ab"
  - bucketName    : "my-bucket-a7044ab"

Resources:
    - 4 deleted

Duration: 7s
```

## 步驟二: 移除 Stack 設定

上面步驟只是把所有資源移除，但是你還是保留了所以 stack history 操作，請看

```sh
$ pulumi stack history
Version: 2
UpdateKind: destroy
Status: succeeded
Message: chore(pulumi): 設定 Pulumi Stack 環境變數
+0-4~0 0 Updated 1 minute ago took 8s
    exec.kind: cli
    git.author: Bo-Yi Wu
    git.author.email: xxxxxxxx@gmail.com
    git.committer: Bo-Yi Wu
    git.committer.email: xxxxxxxx@gmail.com
    git.dirty: true
    git.head: 9d9f8182abefb0e90656ca45065bc07a8a3431f4
    git.headName: refs/heads/main
    vcs.kind: github.com
    vcs.owner: go-training
    vcs.repo: infrastructure-as-code-workshop

Version: 1
UpdateKind: update
Status: succeeded
Message: chore(pulumi): 設定 Pulumi Stack 環境變數
+4-0~0 0 Updated 8 minutes ago took 18s
    exec.kind: cli
    git.author: Bo-Yi Wu
    git.author.email: xxxxxxxx@gmail.com
    git.committer: Bo-Yi Wu
    git.committer.email: xxxxxxxx@gmail.com
    git.dirty: true
    git.head: 437e94e130ee3d31eb80075dd237cc17d09255d1
    git.headName: refs/heads/main
    vcs.kind: github.com
    vcs.owner: go-training
    vcs.repo: infrastructure-as-code-workshop
```

要整個完整移除，請務必要執行底下指令

```sh
pulumi stack rm
```

最後的確認

```sh
$ pulumi stack rm
This will permanently remove the 'prod' stack!
Please confirm that this is what you'd like to do by typing ("prod"):
```

## 移除其他的 Stack

按照上面的步驟重新移除其他的 Stack，先使用底下指令列出還有哪些 Stack:

```sh
$ pulumi stack ls
NAME  LAST UPDATE     RESOURCE COUNT  URL
dev   24 minutes ago  5               https://app.pulumi.com/appleboy/demo/dev
```

選擇 Stack

```sh
pulumi stack select dev
```

接著重複上面一跟二步驟即可
