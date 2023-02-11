package main

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/favorite/favoriteaction"
	favorite "github.com/PICOF/simple-tiktok/kitex_gen/favorite"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// LikeAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) LikeAction(ctx context.Context, request *favorite.LikeRequest) (resp *favorite.LikeResponse, err error) {
	// TODO: Your code here...
	return favoriteaction.FavoriteAction(ctx, request)
}

// GetLikeList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) GetLikeList(ctx context.Context, request *favorite.LikeListRequest) (resp *favorite.LikeListResponse, err error) {
	// TODO: Your code here...
	return favoriteaction.FavoriteList(ctx, request)
}
