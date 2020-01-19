package model

import "github.com/nlopes/slack"

type TransactionStatus string

const (
    STARTED   TransactionStatus = "Started"
    COMPLETED TransactionStatus = "Completed"
    ABORTED   TransactionStatus = "Aborted"
)

type Transaction struct {
    Command string `json:"command"`
    Status  TransactionStatus `json:"status"`
}

func NewTransaction(Command string, Status TransactionStatus) Transaction {
    return Transaction{Command: Command, Status: Status}
}

type TransactionMessageRequest struct {
    Token    string `json:"token"`
    Sender   string `json:"sender"`
    Channel  string `json:"channel"`
    Receiver string `json:"receiver"`
    Ephemeral bool `json:"ephemeral"`
    Message  Transaction `json:"message"`
}

func NewTransactionMessageRequest(Token, Sender, Channel, Receiver string, ephemeral bool, transaction Transaction) *TransactionMessageRequest {
    return &TransactionMessageRequest{
        Token:    Token,
        Sender:   Sender,
        Channel:  Channel,
        Receiver: Receiver,
        Ephemeral: ephemeral,
        Message:  transaction,
    }
}

func (obj *TransactionMessageRequest) GetBlocks() []slack.Block {
    return obj.Message.getBlocks()
}

func (obj *TransactionMessageRequest) GetReceiver() string {
    return obj.Receiver
}

func (obj *TransactionMessageRequest) GetChannel() string {
    return obj.Channel
}

func (obj *TransactionMessageRequest) GetToken() string {
    return obj.Token
}

func (obj *TransactionMessageRequest) IsEphemeral() bool {
    return obj.Ephemeral
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


