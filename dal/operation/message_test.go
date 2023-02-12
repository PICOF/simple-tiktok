package operation

import (
	"context"
	"testing"
	"time"
)

func TestSendMessage(t *testing.T) {
	err := SendMessage(context.Background(), 22, 24, "hello", time.Now())
	if err != nil {
		return
	}
}
