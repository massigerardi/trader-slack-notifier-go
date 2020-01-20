package model

import "github.com/nlopes/slack"

type TransactionStatus string

const (
	STARTED   TransactionStatus = "Started"
	COMPLETED TransactionStatus = "Completed"
	ABORTED   TransactionStatus = "Aborted"
)

type Transaction struct {
	Command string            `json:"command"`
	Status  TransactionStatus `json:"status"`
}

func NewTransaction(Command string, Status TransactionStatus) Transaction {
	return Transaction{Command: Command, Status: Status}
}

func NewTransactionMessageRequest(Token, Sender, Channel, Receiver string, ephemeral bool, transaction Transaction) *MessageRequest {
	return &MessageRequest{
		Token:     Token,
		Sender:    Sender,
		Channel:   Channel,
		Receiver:  Receiver,
		Ephemeral: ephemeral,
		Type:      TRANSACTION,
		Message:   &transaction,
	}
}

func (obj *Transaction) getBlocks() []slack.Block {
	Command := slack.NewSectionBlock(
		slack.NewTextBlockObject("plain_text", obj.Command, false, false),
		nil,
		nil)
	Status := slack.NewSectionBlock(
		slack.NewTextBlockObject("plain_text", string(obj.Status), false, false),
		nil,
		nil)
	return []slack.Block{Command, Status}
}
