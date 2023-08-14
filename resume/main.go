package main

import (
    "github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/storage"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        // Create a GCP Storage Bucket to serve your static website
        bucket, err := storage.NewBucket(ctx, "my &storage.BucketArgs{
            Name     pulumi.String("myBucket"),
            Location: pulumi.String("US"),
            Website: &storage.BucketWebsiteArgs{
                MainPageSuffix: pulumi.String("index.html"),
                NotFoundPage:  pulumi.String("404.html"),
            },
        })
        if err != nil {
            return err
        }

        // Upload index.html to the bucket
        _, err = storage.NewBucketObject "indexHtml", &storage.BucketObject			B:      bucket.Name			Name:        pulumi.String("index.html"),
            Content:     pulumi.String("1"), // Replace with actual HTML content
            ContentType: pulumi.String("text/html"),
        })
        if err != nil {
            return err
        }

        // Upload 404.html to the bucket
        _, err = storage.NewBucketObject(ctx, "not", &storage.BucketObjectArgs{
            Bucket:      bucket.Name			Name        pulumi.String("404.html"),
            Content:     pulumi.String("2"), // Replace with actual HTML content
            ContentType: pulumi.String("text/html"),
        })
        if err != nil {
            return err
        }

        // Export the DNS name of the bucket
        ctx.Export("bucketUrl", pulumi.Sprintf("http://storage.googleapis.com/%s", bucket.Name))

        return nil
    })
}
