package minio

import (
	config2 "github.com/PICOF/simple-tiktok/pkg/config"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

var (
	config          *viper.Viper
	client          *minio.Client
	endpoint        string
	accessKeyID     string
	secretAccessKey string
	secure          bool
	BucketName      string
)

func init() {
	initConfig()
	var err error
	client, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: secure})
	if err != nil {
		klog.Fatalf("minio连接错误: %v", err)
	}
	klog.Infof("%#v\n", client)
	createBucket(BucketName)
}

func initConfig() {
	config = config2.GetConfig("minio")
	endpoint = config.GetString("endpoint")
	accessKeyID = config.GetString("accessKeyID")
	secretAccessKey = config.GetString("secretAccessKey")
	secure = config.GetBool("secure")
	BucketName = config.GetString("videoBucket")
}
