package operation

import (
	"context"
	"github.com/PICOF/simple-tiktok/dal"
	"github.com/PICOF/simple-tiktok/dal/minio"
	"github.com/PICOF/simple-tiktok/pkg/cover"
	"github.com/cloudwego/kitex/pkg/klog"
)

func GetVideoById(ctx context.Context, id int64) (list []TVideoInfo, err error) {
	err = dal.DB.WithContext(ctx).Order("publish_time").Where("author_id = ?", id).Find(&list).Error
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to get video list by id: %v", err)
		return
	}
	return
}

func PublishVideo(ctx context.Context, info *TVideoInfo, data []byte, objName string) (err error) {
	filename := objName + ".mp4"
	coverName := objName + ".jpg"
	err = minio.UploadFile(ctx, minio.BucketName, filename, data, int64(len(data)), "video/mp4")
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to upload video file: %v", err)
		return err
	}
	fileUrl, err := minio.GetFileUrl(ctx, minio.BucketName, filename)
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to get file url: %v", err)
		return err
	}
	info.PlayUrl = fileUrl
	videoCover, err := cover.GetCover(ctx, fileUrl)
	if err != nil {
		return err
	}
	err = minio.UploadFile(ctx, minio.BucketName, coverName, videoCover, int64(len(videoCover)), "image/jpeg")
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to upload cover file: %v", err)
		return err
	}
	coverUrl, err := minio.GetFileUrl(ctx, minio.BucketName, coverName)
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to get cover url: %v", err)
		return err
	}
	info.CoverUrl = coverUrl
	if err := dal.DB.Create(info).Error; err != nil {
		klog.CtxErrorf(ctx, "Failed to create video info: %v", err)
		return err
	}
	return
}
