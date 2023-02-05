// Code generated by Kitex v0.4.4. DO NOT EDIT.

package feedservice

import (
	"context"
	feed "github.com/PICOF/simple-tiktok/kitex_gen/feed"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return feedServiceServiceInfo
}

var feedServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "FeedService"
	handlerType := (*feed.FeedService)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetVideoList": kitex.NewMethodInfo(getVideoListHandler, newFeedServiceGetVideoListArgs, newFeedServiceGetVideoListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "feed",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func getVideoListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*feed.FeedServiceGetVideoListArgs)
	realResult := result.(*feed.FeedServiceGetVideoListResult)
	success, err := handler.(feed.FeedService).GetVideoList(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFeedServiceGetVideoListArgs() interface{} {
	return feed.NewFeedServiceGetVideoListArgs()
}

func newFeedServiceGetVideoListResult() interface{} {
	return feed.NewFeedServiceGetVideoListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetVideoList(ctx context.Context, request *feed.FeedRequest) (r *feed.FeedResponse, err error) {
	var _args feed.FeedServiceGetVideoListArgs
	_args.Request = request
	var _result feed.FeedServiceGetVideoListResult
	if err = p.c.Call(ctx, "GetVideoList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
