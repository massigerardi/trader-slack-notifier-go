package model

import (
    "github.com/nlopes/slack"
)

type MessageRequest interface {
    GetBlocks() []slack.Block
    GetChannel() string
    GetReceiver() string
    GetToken() string
    IsEphemeral() bool
}

type MessageResponse struct {
    Success bool `json:"success"`
    Err     string `json:"err,omitempty"`
}



