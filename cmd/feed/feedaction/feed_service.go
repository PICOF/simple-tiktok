package feedaction

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/constant"
	"github.com/PICOF/simple-tiktok/dal/operation"
	"github.com/PICOF/simple-tiktok/kitex_gen/feed"
	"github.com/cloudwego/kitex/pkg/klog"
	"strconv"
	"time"
)

func GetVideoListHandler(ctx context.Context, request *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	var latestTime int64
	if request.LatestTime == nil {
		latestTime = time.Now().UnixMilli()
	} else {
		latestTime, err = strconv.ParseInt(*request.LatestTime, 10, 64)
		if err != nil {
			klog.CtxErrorf(ctx, "Failed to format 'last_time': %v", err)
			return
		}
	}
	list, nextTime, err := GetVideoList(ctx, request.GetUserId(), latestTime)
	if err != nil {
		klog.CtxErrorf(ctx, "An error occurred while fetching the list of videos: %v", err)
		resp = PackResponse(nil, constant.Failed, nil)
		return
	}
	resp = PackResponse(list, constant.Success, nextTime)
	return
}
func GetVideoList(ctx context.Context, userId int64, latestTime int64) (list []*feed.VideoInfo, nextTime *int64, err error) {
	videoList, err := operation.GetVideoList(ctx, latestTime)
	if err != nil {
		return nil, nil, err
	}
	list, err = ConvertAllVideoInfo(ctx, userId, videoList)
	if err != nil {
		return
	}
	length := len(videoList)
	if length == 0 {
		nextTime = nil
	} else {
		timestamp := videoList[length-1].PublishTime.UnixMilli()
		nextTime = &timestamp
	}
	return
}
func GetVideoListByIdHandler(ctx context.Context, request *feed.GetByIDRequest) (resp *feed.FeedResponse, err error) {
	var userId, queryId int64
	var list []*feed.VideoInfo
	userId = request.UserId
	queryId = request.QueryId
	list, err = GetVideoListById(ctx, userId, queryId)
	if err != nil {
		klog.CtxErrorf(ctx, "An error occurred while fetching the list of videos: %v", err)
		resp = PackResponse(nil, constant.Failed, nil)
		return
	}
	resp = PackResponse(list, constant.Success, nil)
	return
}
func GetVideoListById(ctx context.Context, userId int64, queryId int64) (list []*feed.VideoInfo, err error) {
	videoList, err := operation.GetVideoById(ctx, queryId)
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to get video by id: %v", err)
		return
	}
	list, err = ConvertAllVideoInfo(ctx, userId, videoList)
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to convert video info: %v", err)
		return
	}
	return
}
