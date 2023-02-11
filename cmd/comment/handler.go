package main

import (
	"context"
	"github.com/PICOF/simple-tiktok/kitex_gen/comment"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, request *comment.CommentRequest) (resp *comment.CommentResponse, err error) {
	// TODO: Your code here...
	return
}

// GetCommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) GetCommentList(ctx context.Context, request *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {
	// TODO: Your code here...
	return
}
