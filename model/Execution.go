package model

import "github.com/nlopes/slack"

type Execution struct {
    Strategy     string `json:"strategy"`
    Transactions []Transaction `json:"transactions,omitempty"`
}

func NewExecution(Strategy string, Transactions ...Transaction) Execution {
    return Execution{Strategy: Strategy, Transactions:Transactions}
}

type ExecutionMessageRequest struct {
    Token    string `json:"token"`
    Sender   string `json:"sender"`
    Channel  string `json:"channel"`
    Receiver string `json:"receiver"`
    Ephemeral bool `json:"ephemeral"`
    Message  Execution `json:"message"`
}

func NewExecutionMessageRequest(Token, Sender, Channel, Receiver string, ephemeral bool, execution Execution) *ExecutionMessageRequest {
    return &ExecutionMessageRequest{
        Token:    Token,
        Sender:   Sender,
        Channel:  Channel,
        Receiver: Receiver,
        Ephemeral: ephemeral,
        Message:  execution,
    }
}

func (obj *ExecutionMessageRequest) GetBlocks() []slack.Block {
    return obj.Message.getBlocks()
}

func (obj *ExecutionMessageRequest) GetReceiver() string {
    return obj.Receiver
}

func (obj *ExecutionMessageRequest) GetChannel() string {
    return obj.Channel
}

func (obj *ExecutionMessageRequest) IsEphemeral() bool {
    return obj.Ephemeral
}

func (obj *ExecutionMessageRequest) GetToken() string {
    return obj.Token
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



