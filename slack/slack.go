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
	if message.Ephemeral {
		return sendEphemeralMessage(message)
	}
	return sendMessage(message)
}

func sendMessage(message model.MessageRequest) (string, error) {
	client := get(message.Token)
	blocks := message.GetBlocks()
	comment := slackLib.NewTextBlockObject("plain_text", fmt.Sprintf("%v sent by golang", message.ID), false, false)
	context := slackLib.NewContextBlock("context"+message.ID, comment)
	blocks = append(blocks, context)
	fmt.Printf("sending message %v to %v channel\n", message.ID, message.Channel)
	_, timestamp, err := client.PostMessage(message.Channel, slackLib.MsgOptionBlocks(blocks...))
	return timestamp, err
}

func sendEphemeralMessage(message model.MessageRequest) (string, error) {
	client := get(message.Token)
	blocks := message.GetBlocks()
	comment := slackLib.NewTextBlockObject("plain_text", fmt.Sprintf("%v sent by golang", message.ID), false, false)
	context := slackLib.NewContextBlock("context"+message.ID, comment)
	blocks = append(blocks, context)
	fmt.Printf("sending ephemeral message to %v channel\n", message.Channel)
	timestamp, err := client.PostEphemeral(message.Channel, message.Receiver, slackLib.MsgOptionBlocks(blocks...))
	return timestamp, err
}
