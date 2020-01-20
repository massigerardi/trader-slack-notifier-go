package model

import "github.com/nlopes/slack"

type Execution struct {
	Strategy     string        `json:"strategy"`
	Transactions []Transaction `json:"transactions,omitempty"`
}

func NewExecution(Strategy string, Transactions ...Transaction) Execution {
	return Execution{Strategy: Strategy, Transactions: Transactions}
}

func NewExecutionMessageRequest(Token, Sender, Channel, Receiver string, ephemeral bool, execution Execution) *MessageRequest {
	return &MessageRequest{
		Token:     Token,
		Sender:    Sender,
		Channel:   Channel,
		Receiver:  Receiver,
		Ephemeral: ephemeral,
		Type:      EXECUTION,
		Message:   &execution,
	}
}

func (obj *Execution) getBlocks() []slack.Block {
	headerText := slack.NewTextBlockObject("plain_text", obj.Strategy, false, false)
	header := slack.NewSectionBlock(headerText, nil, nil)
	transactionText := slack.NewTextBlockObject("plain_text", "Transactions", false, false)
	transactionHeader := slack.NewSectionBlock(transactionText, nil, nil)
	blocks := []slack.Block{header, transactionHeader}
	for _, v := range obj.Transactions {
		blocks = append(blocks, slack.NewDividerBlock())
		for _, b := range v.getBlocks() {
			blocks = append(blocks, b)
		}
	}
	return blocks
}
