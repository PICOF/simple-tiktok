// Code generated by Kitex v0.4.4. DO NOT EDIT.

package relationservice

import (
	"context"
	relation "github.com/PICOF/simple-tiktok/kitex_gen/relation"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	FollowAction(ctx context.Context, request *relation.FollowRequest, callOptions ...callopt.Option) (r *relation.FollowResponse, err error)
	GetFollowList(ctx context.Context, request *relation.RelationListRequest, callOptions ...callopt.Option) (r *relation.RelationListResponse, err error)
	GetFollowerList(ctx context.Context, request *relation.RelationListRequest, callOptions ...callopt.Option) (r *relation.RelationListResponse, err error)
	GetFriendList(ctx context.Context, request *relation.RelationListRequest, callOptions ...callopt.Option) (r *relation.RelationListResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kRelationServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kRelationServiceClient struct {
	*kClient
}

func (p *kRelationServiceClient) FollowAction(ctx context.Context, request *relation.FollowRequest, callOptions ...callopt.Option) (r *relation.FollowResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FollowAction(ctx, request)
}

func (p *kRelationServiceClient) GetFollowList(ctx context.Context, request *relation.RelationListRequest, callOptions ...callopt.Option) (r *relation.RelationListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFollowList(ctx, request)
}

func (p *kRelationServiceClient) GetFollowerList(ctx context.Context, request *relation.RelationListRequest, callOptions ...callopt.Option) (r *relation.RelationListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFollowerList(ctx, request)
}

func (p *kRelationServiceClient) GetFriendList(ctx context.Context, request *relation.RelationListRequest, callOptions ...callopt.Option) (r *relation.RelationListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFriendList(ctx, request)
}
