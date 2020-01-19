package model

import "github.com/nlopes/slack"

type Text struct {
    Text string `json:"text"`
}

func NewText(text string) Text {
    return Text{Text: text}
}

type TextMessageRequest struct {
    Token    string `json:"token"`
    Sender   string `json:"sender"`
    Channel  string `json:"channel"`
    Receiver string `json:"receiver"`
    Ephemeral bool  `json:"ephemeral"`
    Message  Text   `json:"message"`
}


func (obj *TextMessageRequest) GetBlocks() []slack.Block {
    return obj.Message.getBlocks()
}

func NewTextMessageRequest(Token, Sender, Channel, Receiver string, ephemeral bool, Message Text) *TextMessageRequest {
    return &TextMessageRequest{
        Token:    Token,
        Sender:   Sender,
        Channel:  Channel,
        Receiver: Receiver,
        Ephemeral: ephemeral,
        Message:  Message,
    }
}



func (obj *TextMessageRequest) GetReceiver() string {
    return obj.Receiver
}

func (obj *TextMessageRequest) GetChannel() string {
    return obj.Channel
}

func (obj *TextMessageRequest) GetToken() string {
    return obj.Token
}

func (obj *TextMessageRequest) IsEphemeral() bool {
    return obj.Ephemeral
}

func (obj *Text) getBlocks() []slack.Block {
    text := slack.NewSectionBlock(
        slack.NewTextBlockObject("plain_text", obj.Text, false, false),
        nil,
        nil)
    return []slack.Block{text}
}
