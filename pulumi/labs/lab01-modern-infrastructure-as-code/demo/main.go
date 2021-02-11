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

		site := getEnv(ctx, "s3:siteDir", "content")
		index := path.Join(site, "index.html")
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

func getEnv(ctx *pulumi.Context, key string, fallback ...string) string {
	if value, ok := ctx.GetConfig(key); ok {
		return value
	}

	if len(fallback) > 0 {
		return fallback[0]
	}

	return ""
}
