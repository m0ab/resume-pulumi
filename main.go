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
			Name:     pulumi.String("www.markbennett.info"),
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

		// Read and Upload index.html
		indexContentBytes, err := ioutil.ReadFile("index.html")
		if err != nil {
			return err
		}
		_, err = storage.NewBucketObject(ctx, "indexHtml", &storage.BucketObjectArgs{
			Bucket:      bucket.Name,
			Content:     pulumi.String(string(indexContentBytes)),
			ContentType: pulumi.String("text/html"),
			Name:        pulumi.String("index.html"),
		})
		if err != nil {
			return err
		}

		// Read and Upload 404.html
		notFoundContentBytes, err := ioutil.ReadFile("404.html")
		if err != nil {
			return err
		}
		_, err = storage.NewBucketObject(ctx, "error404Html", &storage.BucketObjectArgs{
			Bucket:      bucket.Name,
			Content:     pulumi.String(string(notFoundContentBytes)),
			ContentType: pulumi.String("text/html"),
			Name:        pulumi.String("404.html"),
		})
		if err != nil {
			return err
		}

        resumeContentBytes, err := ioutil.ReadFile("resume.html")
		_, err = storage.NewBucketObject(ctx, "resumeHtml", &storage.BucketObjectArgs{
			Bucket:      bucket.Name,
			Content:     pulumi.String(string(resumeContentBytes)),
			ContentType: pulumi.String("text/html"),
			Name:        pulumi.String("resume.html"),
		})
		if err != nil {
			return err
		}

		ctx.Export("bucketURL", pulumi.Sprintf("https://storage.googleapis.com/%s", bucket.Name))



		return nil
	})
}
