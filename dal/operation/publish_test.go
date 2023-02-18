package operation

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetVideoById(t *testing.T) {
	ret, err := GetMessage(context.Background(), 11, 45, 14)
	if err != nil {
		return
	}
	fmt.Println(ret)
}

func TestPublishVideo(t *testing.T) {
	// create a test context
	ctx := context.Background()

	// create dummy data for the video
	info := &TVideoInfo{
		Title:       "Test Video",
		PublishTime: time.Now(),
	}
	data := []byte("Test video data")
	objName := "test-video"

	// call the PublishVideo function with the test data
	err := PublishVideo(ctx, info, data, objName)
	if err != nil {
		t.Errorf("PublishVideo returned error: %v", err)
		return
	}
}
