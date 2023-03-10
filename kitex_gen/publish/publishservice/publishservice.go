// Code generated by Kitex v0.4.4. DO NOT EDIT.

package publishservice

import (
	"context"
	publish "github.com/PICOF/simple-tiktok/kitex_gen/publish"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return publishServiceServiceInfo
}

var publishServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "PublishService"
	handlerType := (*publish.PublishService)(nil)
	methods := map[string]kitex.MethodInfo{
		"PublishAction":  kitex.NewMethodInfo(publishActionHandler, newPublishServicePublishActionArgs, newPublishServicePublishActionResult, false),
		"GetPublishList": kitex.NewMethodInfo(getPublishListHandler, newPublishServiceGetPublishListArgs, newPublishServiceGetPublishListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "publish",
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

func publishActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*publish.PublishServicePublishActionArgs)
	realResult := result.(*publish.PublishServicePublishActionResult)
	success, err := handler.(publish.PublishService).PublishAction(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPublishServicePublishActionArgs() interface{} {
	return publish.NewPublishServicePublishActionArgs()
}

func newPublishServicePublishActionResult() interface{} {
	return publish.NewPublishServicePublishActionResult()
}

func getPublishListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*publish.PublishServiceGetPublishListArgs)
	realResult := result.(*publish.PublishServiceGetPublishListResult)
	success, err := handler.(publish.PublishService).GetPublishList(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPublishServiceGetPublishListArgs() interface{} {
	return publish.NewPublishServiceGetPublishListArgs()
}

func newPublishServiceGetPublishListResult() interface{} {
	return publish.NewPublishServiceGetPublishListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) PublishAction(ctx context.Context, request *publish.PublishRequest) (r *publish.PublishResponse, err error) {
	var _args publish.PublishServicePublishActionArgs
	_args.Request = request
	var _result publish.PublishServicePublishActionResult
	if err = p.c.Call(ctx, "PublishAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetPublishList(ctx context.Context, request *publish.PublishListRequest) (r *publish.PublishListResponse, err error) {
	var _args publish.PublishServiceGetPublishListArgs
	_args.Request = request
	var _result publish.PublishServiceGetPublishListResult
	if err = p.c.Call(ctx, "GetPublishList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
