package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/storage"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"io/ioutil"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a GCP Storage Bucket
		bucket, err := storage.NewBucket(ctx, "myBucket", &storage.BucketArgs{
			Name:     pulumi.String("markbennettresume"),
			Location: pulumi.String("US"),
			Website: &storage.BucketWebsiteArgs{
				MainPageSuffix: pulumi.String("index.html"),
				NotFoundPage:   pulumi.String("404.html"),
			},
		})
		if err != nil {
			return err
		}

		// Set IAM policy for the bucket to be publicly readable
		_, err = storage.NewBucketIAMMember(ctx, "publicReadPolicy", &storage.BucketIAMMemberArgs{
			Bucket: bucket.Name,
			Role:   pulumi.String("roles/storage.objectViewer"),
			Member: pulumi.String("allUsers"),
		})
		if err != nil {
			return err
		}

		contentBytes, err := ioutil.ReadFile("index.html")
		if err != nil {
			return err
		}
		content := string(contentBytes)

		_, err = storage.NewBucketObject(ctx, "indexHtml", &storage.BucketObjectArgs{
			Bucket:      bucket.Name,
			Content:     pulumi.String(content),
			ContentType: pulumi.String("text/html"),
			Name:        pulumi.String("index.html"),
		})
		if err != nil {
			return err
		}

		ctx.Export("bucketURL", pulumi.Sprintf("https://storage.googleapis.com/%s", bucket.Name))

		return nil
	})
}
