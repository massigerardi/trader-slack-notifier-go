package services

import (
	"context"
	"slack-notifier/model"
	"slack-notifier/slack"
)

type Notifier interface {
	SendMessage(ctx context.Context, message model.MessageRequest) (result string, err error)
}

type Receiver interface {
	sendStream(ctx context.Context, message model.MessageRequest) (result string, err error)
}

type notifier struct {}

func NewNotifier() Notifier  {
	return notifier{}
}

func (notifier) SendMessage(ctx context.Context, message model.MessageRequest) (result string, err error)  {
	result, err = slack.SendMessage(message)
	return result, err
}
