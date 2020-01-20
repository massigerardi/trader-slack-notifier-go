package model

import (
	"encoding/json"
	"github.com/nlopes/slack"
)

type BodyType string

const (
	EXECUTION   BodyType = "Execution"
	TRANSACTION BodyType = "Transaction"
	TEXT        BodyType = "Text"
)

type MessageRequestI interface {
	GetBlocks() []slack.Block
}

type Message interface {
	getBlocks() []slack.Block
}

type MessageRequest struct {
	ID        string   `json:"id"`
	Token     string   `json:"token"`
	Sender    string   `json:"sender,omitempty"`
	Channel   string   `json:"channel,omitempty"`
	Receiver  string   `json:"receiver"`
	Ephemeral bool     `json:"ephemeral"`
	Type      BodyType `json:"type"`
	Message   Message  `json:"message"`
	Comment   string   `json:"comment"`
}

func (request MessageRequest) GetBlocks() []slack.Block {
	return request.Message.getBlocks()
}

type Envelope struct {
	Token     string      `json:"token"`
	Sender    string      `json:"sender,omitempty"`
	Channel   string      `json:"channel,omitempty"`
	Receiver  string      `json:"receiver"`
	Ephemeral bool        `json:"ephemeral"`
	Type      BodyType    `json:"type"`
	Message   interface{} `json:"message"`
}

type MessageResponse struct {
	Success bool   `json:"success"`
	Err     string `json:"err,omitempty"`
}

var bodyTypeHandlers = map[BodyType]func() Message{
	TEXT:        func() Message { return &Text{} },
	EXECUTION:   func() Message { return &Execution{} },
	TRANSACTION: func() Message { return &Transaction{} },
}

func (request *MessageRequest) UnmarshalJSON(data []byte) error {
	var rawMessage json.RawMessage
	req := Envelope{
		Message: &rawMessage,
	}
	if err := json.Unmarshal(data, &req); err != nil {
		return err
	}
	message := bodyTypeHandlers[req.Type]()
	if err := json.Unmarshal(rawMessage, &message); err != nil {
		return err
	}
	request.Token = req.Token
	request.Sender = req.Sender
	request.Channel = req.Channel
	request.Receiver = req.Receiver
	request.Ephemeral = req.Ephemeral
	request.Type = req.Type
	request.Message = message

	return nil
}
