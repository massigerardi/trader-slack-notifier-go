package model

import (
    "encoding/json"
    "reflect"
    "testing"
)

func TestNewText(t *testing.T) {
    type args struct {
        text string
    }
    tests := []struct {
        name string
        args args
        want Text
    }{
        {
            name: "new",
            args: args{text:"test"},
            want: Text{Text:"test"},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := NewText(tt.args.text); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("NewText() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestNewTextMessageRequest(t *testing.T) {
    type args struct {
        Token     string
        Sender    string
        Channel   string
        Receiver  string
        Ephemeral bool
        Message   Text
    }
    tests := []struct {
        name string
        args args
        want *TextMessageRequest
    }{
        {
            name: "full",
            args: args{
                Token:     "my_token",
                Sender:    "user",
                Channel:   "channel",
                Receiver:  "receiver",
                Ephemeral: false,
                Message:   Text{Text: "test"},
            },
            want: &TextMessageRequest{
                Token:     "my_token",
                Sender:    "user",
                Channel:   "channel",
                Receiver:  "receiver",
                Ephemeral: false,
                Message:   Text{Text: "test"},
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := NewTextMessageRequest(tt.args.Token, tt.args.Sender, tt.args.Channel, tt.args.Receiver, tt.args.Ephemeral, tt.args.Message); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("\ngot      = %v, \nexpected = %v", got, tt.want)
            }
        })
    }
}

func TestTextMessageRequest_GetBlocks(t *testing.T) {
    type fields struct {
        token    string
        sender   string
        channel  string
        receiver string
        message  Text
    }
    tests := []struct {
        name   string
        fields fields
        want   string
    }{
        {
            name: "test_1",
            fields: fields{
                token:    "my_token",
                sender:   "user",
                receiver: "receiver",
                channel:  "channel",
                message:  Text{Text: "test"},
            },
            want: `[{"type":"section","text":{"type":"plain_text","text":"test"}}]`,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            obj := &TextMessageRequest{
                Token:    tt.fields.token,
                Sender:   tt.fields.sender,
                Receiver: tt.fields.receiver,
                Message:  tt.fields.message,
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

func TestTextMessageRequest_GetChannel(t *testing.T) {
    type fields struct {
        Token     string
        Sender    string
        Channel   string
        Receiver  string
        Ephemeral bool
        Message   Text
    }
    tests := []struct {
        name   string
        fields fields
        want   string
    }{
        {
            name:   "get",
            fields: fields{
                Channel:"test_channel",
            },
            want:   "test_channel",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            obj := &TextMessageRequest{
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

func TestTextMessageRequest_GetReceiver(t *testing.T) {
    type fields struct {
        Token     string
        Sender    string
        Channel   string
        Receiver  string
        Ephemeral bool
        Message   Text
    }
    tests := []struct {
        name   string
        fields fields
        want   string
    }{
        {
            name:   "get",
            fields: fields{
                Receiver:"test_receiver",
            },
            want:   "test_receiver",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            obj := &TextMessageRequest{
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

func TestTextMessageRequest_GetToken(t *testing.T) {
    type fields struct {
        Token     string
        Sender    string
        Channel   string
        Receiver  string
        Ephemeral bool
        Message   Text
    }
    tests := []struct {
        name   string
        fields fields
        want   string
    }{
        {
            name:   "get",
            fields: fields{
                Token:"test_token",
            },
            want:   "test_token",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            obj := &TextMessageRequest{
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

func TestTextMessageRequest_IsEphemeral(t *testing.T) {
    type fields struct {
        Token     string
        Sender    string
        Channel   string
        Receiver  string
        Ephemeral bool
        Message   Text
    }
    tests := []struct {
        name   string
        fields fields
        want   bool
    }{
        {
            name:   "get",
            fields: fields{
                Ephemeral: false,
            },
            want: false,
        },

    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            obj := &TextMessageRequest{
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

func TestText_getBlocks(t *testing.T) {
    type fields struct {
        text string
    }
    tests := []struct {
        name   string
        fields fields
        want   string
    }{
        {
            name:   "",
            fields: fields{text: "test"},
            want:   `[{"type":"section","text":{"type":"plain_text","text":"test"}}]`,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            obj := NewText(tt.fields.text)
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



func TestTextJson(t *testing.T) {
    tests := []struct {
        name   string
        object Text
        want   string
    }{
        {
            name: "",
            object: Text{
                Text: "sell",
            },
            want: `{"text":"sell"}`,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            text := &Text{}
            err := json.Unmarshal([]byte(tt.want), text)
            if err != nil {
                t.Errorf("Fail:\n  got: %v\n want: %v", tt.want, err)
            }
            if !reflect.DeepEqual(text, &tt.object) {
                t.Errorf("Fail:\n  got: %v\n want: %v", text, &tt.object)
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
