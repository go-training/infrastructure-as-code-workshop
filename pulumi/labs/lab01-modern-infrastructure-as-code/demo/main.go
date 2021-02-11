package main

import (
	"io/ioutil"
	"mime"
	"path"
	"path/filepath"

	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an AWS resource (S3 Bucket)
		bucket, err := s3.NewBucket(ctx, "my-bucket", &s3.BucketArgs{
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
