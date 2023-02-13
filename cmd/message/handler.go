package main

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/message/messageaction"
	message "github.com/PICOF/simple-tiktok/kitex_gen/message"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// SendMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) SendMessage(ctx context.Context, request *message.MessageRequest) (resp *message.MessageResponse, err error) {
	// TODO: Your code here...
	return messageaction.SendMessage(ctx, request)
}

// GetChatRecord implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) GetChatRecord(ctx context.Context, request *message.ChatRecordRequest) (resp *message.ChatRecordResponse, err error) {
	// TODO: Your code here...
	return messageaction.GetRecord(ctx, request)
}
