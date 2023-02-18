package operation

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetCommentList(t *testing.T) {
	ret, err := GetCommentList(context.Background(), 12)
	if err != nil {
		return
	}
	fmt.Println(ret)
}

func TestDeleteComment(t *testing.T) {
	err := DeleteComment(context.Background(), 11, 45, 14)
	if err != nil {
		return
	}
}

func TestPublishComment(t *testing.T) {
	ret, err := PublishComment(context.Background(), TComment{
		UserId:     1,
		VideoId:    1,
		Content:    "11",
		GmtCreated: time.Now(),
	})
	if err != nil {
		return
	}
	fmt.Println(ret)
}
