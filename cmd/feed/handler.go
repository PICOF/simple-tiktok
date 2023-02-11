package main

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/feed/feedaction"
	feed "github.com/PICOF/simple-tiktok/kitex_gen/feed"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// GetVideoList implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) GetVideoList(ctx context.Context, request *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	// TODO: Your code here...
	return feedaction.GetVideoListHandler(ctx, request)
}

// GetVideoListById implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) GetVideoListById(ctx context.Context, request *feed.GetByIDRequest) (resp *feed.FeedResponse, err error) {
	// TODO: Your code here...
	return feedaction.GetVideoListByIdHandler(ctx, request)
}

// ConvertVideoList implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) ConvertVideoList(ctx context.Context, request *feed.GetByIDRequest) (resp *feed.FeedResponse, err error) {
	// TODO: Your code here...
	return
}
