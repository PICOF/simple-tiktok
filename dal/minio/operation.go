package minio

import (
	"bytes"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
	"net/url"
	"time"
)

func createBucket(bucketName string) {
	if len(bucketName) == 0 {
		klog.Fatal("invalid bucket name")
	}
	err := client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: "cn-south-1", ObjectLocking: false})
	if err != nil {
		if err.(minio.ErrorResponse).StatusCode == 409 {
			klog.Warnf("bucket: %s已经存在", bucketName)
			return
		}
		klog.Fatalf("Failed to create bucket: %v", err)
	} else {
		klog.Infof("Successfully created %s", bucketName)
	}
}

func UploadFile(ctx context.Context, bucketName string, objectName string, data []byte, objectsize int64, contentType string) error {
	n, err := client.PutObject(context.Background(), bucketName, objectName, bytes.NewReader(data), objectsize, minio.PutObjectOptions{
		ContentType: contentType,
	})
	println(n.Location)
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to upload file %s: %v", bucketName, err)
		return err
	}
	klog.CtxInfof(ctx, "Successfully uploaded file %s", objectName)
	return nil
}

// GetFileUrl : 获取到的连接有时限，应当即时获取，同时做缓存策略
func GetFileUrl(ctx context.Context, bucketName string, fileName string) (fileUrl string, err error) {
	expires := time.Hour * 24 * 7
	var resUrl *url.URL
	resUrl, err = client.PresignedGetObject(ctx, bucketName, fileName, expires, nil)
	if err != nil {
		klog.Errorf("Failed to get url of file %s in bucket '%s': %v", fileName, bucketName, err)
		return
	}
	fileUrl = resUrl.Scheme + "://" + resUrl.Host + resUrl.Path
	return
}
