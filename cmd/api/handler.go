package main

import (
	"context"
	feed "github.com/PICOF/simple-tiktok/kitex_gen/feed"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// GetVideoList implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) GetVideoList(ctx context.Context, request *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	// TODO: Your code here...
	return
}
