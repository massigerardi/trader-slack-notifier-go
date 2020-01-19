package slack

import (
	"fmt"
	slackLib "github.com/nlopes/slack"
	"slack-notifier/model"
)

var clients = make(map[string]*(slackLib.Client))

func get(token string) *(slackLib.Client) {
	if clients[token] == nil {
		clients[token] = slackLib.New(token)
	}
	return clients[token]
}

func SendMessage(message model.MessageRequest) (string, error) {
	if message.IsEphemeral() {
		return sendEphemeralMessage(message)
	}
	return sendMessage(message)
}

func sendMessage(message model.MessageRequest) (string, error) {
	client := get(message.GetToken())
	blocks := message.GetBlocks();
	fmt.Printf("sending message to %v channel\n", message.GetChannel())
	_, timestamp, err := client.PostMessage(message.GetChannel(), slackLib.MsgOptionBlocks(blocks...))
	return timestamp, err
}

func sendEphemeralMessage(message model.MessageRequest) (string, error) {
	client := get(message.GetToken())
	blocks := message.GetBlocks();
	fmt.Printf("sending ephemeral message to %v channel\n", message.GetChannel())
	timestamp, err := client.PostEphemeral(message.GetChannel(), message.GetReceiver(), slackLib.MsgOptionBlocks(blocks...))
	return timestamp, err
}
