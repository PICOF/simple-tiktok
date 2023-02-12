package main

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/relation/relationaction"
	relation "github.com/PICOF/simple-tiktok/kitex_gen/relation"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// FollowAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) FollowAction(ctx context.Context, request *relation.FollowRequest) (resp *relation.FollowResponse, err error) {
	// TODO: Your code here...
	return relationaction.FollowAction(ctx, request)
}

// GetFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowList(ctx context.Context, request *relation.RelationListRequest) (resp *relation.RelationListResponse, err error) {
	// TODO: Your code here...
	return relationaction.GetFollowList(ctx, request)
}

// GetFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowerList(ctx context.Context, request *relation.RelationListRequest) (resp *relation.RelationListResponse, err error) {
	// TODO: Your code here...
	return relationaction.GetFollowerList(ctx, request)
}

// GetFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFriendList(ctx context.Context, request *relation.RelationListRequest) (resp *relation.RelationListResponse, err error) {
	// TODO: Your code here...
	return relationaction.GetFriendList(ctx, request)
}
