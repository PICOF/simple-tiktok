package main

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/publish/publishaction"
	publish "github.com/PICOF/simple-tiktok/kitex_gen/publish"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishAction implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishAction(ctx context.Context, request *publish.PublishRequest) (resp *publish.PublishResponse, err error) {
	// TODO: Your code here...
	return publishaction.PublishAction(ctx, request)
}

// GetPublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) GetPublishList(ctx context.Context, request *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
	// TODO: Your code here...
	return publishaction.GetPublishList(ctx, request)
}
