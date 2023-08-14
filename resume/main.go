package main

import (
    "github.com/pulumi/pulumi-google-native/sdk/go/google/storage/v1"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        // GCP Storage Bucket to serve the static website
        bucket, err := v1.NewBucket(ctx, "myBucket", &v1.BucketArgs{
            Name:     pulumi.String("mark-bennett-resume"),
            Location: pulumi.String("US"),
            Website: &v1.BucketWebsiteArgs{
                MainPageSuffix: pulumi.String("index.html"),
                NotFoundPage:  pulumi.String("404.html"),
            },
        })
        if err != nil {
            return err
        }

        // Export the website URL of the storage bucket
        ctx.Export("websiteUrl",.SelfLink.ApplyT(func(link string) (string, error) {
            return "http://storage.googleapis.com/" + link, nil
        }).(pulumi.StringOutput))

        return nil
    })
}
