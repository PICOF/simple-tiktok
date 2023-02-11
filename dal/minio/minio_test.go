package minio

import (
	"context"
	"fmt"
	"os"
	"testing"
)

const bucketName = "my-test"

func getFile() ([]byte, os.FileInfo) {
	filePath := "./big_buck_bunny.mp4"
	data, _ := os.ReadFile(filePath)
	f, _ := os.Stat(filePath)
	return data, f
}

// 第一次测试时先创建测试 bucket
func TestCreateBucket(t *testing.T) {
	createBucket(bucketName)
}

func TestUpload(t *testing.T) {
	file, f := getFile()
	err := UploadFile(context.Background(), bucketName, f.Name(), file, f.Size(), "video/mp4")
	fmt.Println(err)
}
func TestUrl(t *testing.T) {
	_, f := getFile()
	url, err := GetFileUrl(context.Background(), bucketName, f.Name())
	if err != nil {
		return
	}
	fmt.Println(url)
}
