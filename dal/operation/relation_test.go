package operation

import (
	"context"
	"fmt"
	"testing"
)

func TestFollow(t *testing.T) {
	err := FollowAction(context.Background(), 24, 22)
	if err != nil {
		return

	}
}
func TestUnfollow(t *testing.T) {
	err := UnfollowAction(context.Background(), 22, 24)
	if err != nil {
		return
	}
}
func TestGetList(t *testing.T) {
	list, err := GetFriendList(context.Background(), 22)
	if err != nil {
		return
	}
	fmt.Println(list)
	list, err = GetFollowList(context.Background(), 22)
	if err != nil {
		return
	}
	fmt.Println(list)
	list, err = GetFollowerList(context.Background(), 22)
	if err != nil {
		return
	}
	fmt.Println(list)
}
