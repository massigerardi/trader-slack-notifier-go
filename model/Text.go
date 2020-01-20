package model

import "github.com/nlopes/slack"

type Text struct {
	Text string `json:"text"`
}

func NewText(text string) Text {
	return Text{Text: text}
}

func NewTextMessageRequest(Token, Sender, Channel, Receiver string, ephemeral bool, message Text) *MessageRequest {
	return &MessageRequest{
		Token:     Token,
		Sender:    Sender,
		Channel:   Channel,
		Receiver:  Receiver,
		Type:      TEXT,
		Ephemeral: ephemeral,
		Message:   &message,
	}
}

func (obj *Text) getBlocks() []slack.Block {
	text := slack.NewSectionBlock(
		slack.NewTextBlockObject("plain_text", obj.Text, false, false),
		nil,
		nil)
	return []slack.Block{text}
}
