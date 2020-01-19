package model

import (
    "encoding/json"
    "reflect"
    "testing"
)

var transactionSellCompleted = Transaction{
    Command: "sell",
    Status:  COMPLETED,
}

var transactionBuyAborted = Transaction{
    Command: "buy",
    Status:  ABORTED,
}

func TestNewTransaction(t *testing.T) {
    tests := []struct {
        name string
        args Transaction
        want Transaction
    }{
        {
            name: "new",
            args: transactionSellCompleted,
            want: transactionSellCompleted,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := NewTransaction(tt.args.Command, tt.args.Status); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("NewTransaction() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestTransaction_getBlocks(t *testing.T) {
    type fields struct {
        Command string
        Status  TransactionStatus
    }
    tests := []struct {
        name   string
        fields fields
        want   string
    }{
        {
            name: "Completed",
            fields: fields{
                Command: "sell",
                Status:  COMPLETED,
            },
            want: `[{"type":"section","text":{"type":"plain_text","text":"sell"}},{"type":"section","text":{"type":"plain_text","text":"Completed"}}]`,
        },
        {
            name: "Aborted",
            fields: fields{
                Command: "buy",
                Status:  ABORTED,
            },
            want: `[{"type":"section","text":{"type":"plain_text","text":"buy"}},{"type":"section","text":{"type":"plain_text","text":"Aborted"}}]`,
        },
        {
            name: "Started",
            fields: fields{
                Command: "buy",
                Status:  STARTED,
            },
            want: `[{"type":"section","text":{"type":"plain_text","text":"buy"}},{"type":"section","text":{"type":"plain_text","text":"Started"}}]`,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            obj := &Transaction{
                Command: tt.fields.Command,
                Status:  tt.fields.Status,
            }
            got, err := json.Marshal(obj.getBlocks())
            if err != nil {
                t.Errorf("getBlocks() error = %v", err)
            }
            if !reflect.DeepEqual(string(got), tt.want) {
                t.Errorf("\ngot      = %v, \nexpected = %v", string(got), tt.want)
            }
        })
    }
}

func TestTransactionMessageRequest_GetBlocks(t *testing.T) {
    type fields struct {
        token     string
        sender    string
        receiver  string
        ephemeral bool
        message   Transaction
    }
    tests := []struct {
        name   string
        fields fields
        want   string
    }{
        {
            name: "",
            fields: fields{
                token:     "my_token",
                sender:    "user",
                receiver:  "channel",
                ephemeral: false,
                message: Transaction{
                    Command: "buy",
                    Status:  ABORTED,
                },
            },
            want: `[{"type":"section","text":{"type":"plain_text","text":"buy"}},{"type":"section","text":{"type":"plain_text","text":"Aborted"}}]`,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            obj := &TransactionMessageRequest{
                Token:     tt.fields.token,
                Sender:    tt.fields.sender,
                Receiver:  tt.fields.receiver,
                Ephemeral: tt.fields.ephemeral,
                Message:   tt.fields.message,
            }
            got, err := json.Marshal(obj.GetBlocks())
            if err != nil {
                t.Errorf("GetBlocks() error = %v", err)
            }
            if !reflect.DeepEqual(string(got), tt.want) {
                t.Errorf("GetBlocks() = %v, want %v", string(got), tt.want)
            }
        })
    }
}

func TestNewTransactionMessageRequest(t *testing.T) {
    type args struct {
        Token       string
        Sender      string
        Channel     string
        Receiver    string
        Ephemeral   bool
        transaction Transaction
    }
    tests := []struct {
        name string
        args args
        want *TransactionMessageRequest
    }{
        {
            name: "",
            args: args{
                Token:     "my_token",
                Sender:    "user",
                Channel:   "channel",
                Receiver:  "receiver",
                Ephemeral: false,
                transaction: Transaction{
                    Command: "buy",
                    Status:  ABORTED,
                },
            },
            want: &TransactionMessageRequest{
                Token:     "my_token",
                Sender:    "user",
                Channel:   "channel",
                Receiver:  "receiver",
                Ephemeral: false,
                Message: Transaction{
                    Command: "buy",
                    Status:  ABORTED,
                },
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := NewTransactionMessageRequest(tt.args.Token, tt.args.Sender, tt.args.Channel, tt.args.Receiver, tt.args.Ephemeral, tt.args.transaction); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("NewTransactionMessageRequest() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestTransactionJson(t *testing.T) {
    tests := []struct {
        name   string
        object Transaction
        want   string
    }{
        {
            name: "",
            object: Transaction{
                Command: "sell",
                Status:  COMPLETED,
            },
            want: `{"command":"sell","status":"Completed"}`,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            transaction := &Transaction{}
            err := json.Unmarshal([]byte(tt.want), transaction)
            if err != nil {
                t.Errorf("Fail:\n  got: %v\n want: %v", tt.want, err)
            }
            if !reflect.DeepEqual(transaction, &tt.object) {
                t.Errorf("Fail:\n  got: %v\n want: %v", transaction, &tt.object)
            }
            data, err := json.Marshal(&tt.object)
            if err != nil {
                t.Errorf("json.Marshal(%v) gave error %v", tt.object, err)
            }
            if !reflect.DeepEqual(string(data), tt.want) {
                t.Errorf("Fail:\n  got: %v\n want: %v", string(data), tt.want)
            }

        })
    }
}
func TestTransactionMessageRequest_GetChannel(t *testing.T) {
    type fields struct {
        Token     string
        Sender    string
        Channel   string
        Receiver  string
        Ephemeral bool
        Message   Transaction
    }
    tests := []struct {
        name   string
        fields fields
        want   string
    }{
        {
            name: "get",
            fields: fields{
                Channel: "test_value",
            },
            want: "test_value",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            obj := &TransactionMessageRequest{
                Token:     tt.fields.Token,
                Sender:    tt.fields.Sender,
                Channel:   tt.fields.Channel,
                Receiver:  tt.fields.Receiver,
                Ephemeral: tt.fields.Ephemeral,
                Message:   tt.fields.Message,
            }
            if got := obj.GetChannel(); got != tt.want {
                t.Errorf("GetChannel() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestTransactionMessageRequest_GetReceiver(t *testing.T) {
    type fields struct {
        Token     string
        Sender    string
        Channel   string
        Receiver  string
        Ephemeral bool
        Message   Transaction
    }
    tests := []struct {
        name   string
        fields fields
        want   string
    }{
        {
            name: "get",
            fields: fields{
                Receiver: "test_value",
            },
            want: "test_value",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            obj := &TransactionMessageRequest{
                Token:     tt.fields.Token,
                Sender:    tt.fields.Sender,
                Channel:   tt.fields.Channel,
                Receiver:  tt.fields.Receiver,
                Ephemeral: tt.fields.Ephemeral,
                Message:   tt.fields.Message,
            }
            if got := obj.GetReceiver(); got != tt.want {
                t.Errorf("GetReceiver() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestTransactionMessageRequest_GetToken(t *testing.T) {
    type fields struct {
        Token     string
        Sender    string
        Channel   string
        Receiver  string
        Ephemeral bool
        Message   Transaction
    }
    tests := []struct {
        name   string
        fields fields
        want   string
    }{
        {
            name: "get",
            fields: fields{
                Token: "test_value",
            },
            want: "test_value",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            obj := &TransactionMessageRequest{
                Token:     tt.fields.Token,
                Sender:    tt.fields.Sender,
                Channel:   tt.fields.Channel,
                Receiver:  tt.fields.Receiver,
                Ephemeral: tt.fields.Ephemeral,
                Message:   tt.fields.Message,
            }
            if got := obj.GetToken(); got != tt.want {
                t.Errorf("GetToken() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestTransactionMessageRequest_IsEphemeral(t *testing.T) {
    type fields struct {
        Token     string
        Sender    string
        Channel   string
        Receiver  string
        Ephemeral bool
        Message   Transaction
    }
    tests := []struct {
        name   string
        fields fields
        want   bool
    }{
        {
            name: "get",
            fields: fields{
                Ephemeral: false,
            },
            want: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            obj := &TransactionMessageRequest{
                Token:     tt.fields.Token,
                Sender:    tt.fields.Sender,
                Channel:   tt.fields.Channel,
                Receiver:  tt.fields.Receiver,
                Ephemeral: tt.fields.Ephemeral,
                Message:   tt.fields.Message,
            }
            if got := obj.IsEphemeral(); got != tt.want {
                t.Errorf("IsEphemeral() = %v, want %v", got, tt.want)
            }
        })
    }
}
