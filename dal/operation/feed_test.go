package operation

import (
	"context"
	"fmt"
	"testing"
)

func TestGetVideoList(t *testing.T) {
	ret, err := GetFriendList(context.Background(), 1)
	if err != nil {
		return
	}
	fmt.Println(ret)
}
