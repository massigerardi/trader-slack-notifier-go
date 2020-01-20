package model

import (
	"encoding/json"
	"reflect"
	"testing"
)

var example_text = Text{Text: "test"}

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
			args: args{text: "test"},
			want: Text{Text: "test"},
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
		want *MessageRequest
	}{
		{
			name: "full",
			args: args{
				Token:     "my_token",
				Sender:    "user",
				Channel:   "channel",
				Receiver:  "receiver",
				Ephemeral: false,
				Message:   example_text,
			},
			want: &MessageRequest{
				Token:     "my_token",
				Sender:    "user",
				Channel:   "channel",
				Receiver:  "receiver",
				Ephemeral: false,
				Message:   &example_text,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTextMessageRequest(tt.args.Token, tt.args.Sender, tt.args.Channel, tt.args.Receiver, tt.args.Ephemeral, tt.args.Message)
			if !reflect.DeepEqual(got.Message, tt.want.Message) {
				t.Errorf("Fail Marshal:\n  got: %v\n want: %v", got, tt.want)
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
			obj := &MessageRequest{
				Token:    tt.fields.token,
				Sender:   tt.fields.sender,
				Receiver: tt.fields.receiver,
				Message:  &tt.fields.message,
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

func TestTextMessageRequestJson(t *testing.T) {
	tests := []struct {
		name   string
		object Text
		want   string
	}{
		{
			name:   "test",
			object: example_text,
			want:   `{"token":"token","sender":"sender","channel":"channel","receiver":"receiver","ephemeral":false,"type":"Text","message":{"text":"test"}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := NewTextMessageRequest("token", "sender", "channel", "receiver", false, example_text)
			data, err := json.Marshal(&request)
			if err != nil {
				t.Errorf("json.Marshal(%v) gave error %v", tt.object, err)
			}
			if !reflect.DeepEqual(string(data), tt.want) {
				t.Errorf("Fail:\n  got: %v\n want: %v", string(data), tt.want)
			}
			req := &MessageRequest{}
			err = json.Unmarshal([]byte(tt.want), req)
			if err != nil {
				t.Errorf("Fail:\n  got: %v\n want: %v", tt.want, err)
			}
			if !reflect.DeepEqual(req, request) {
				t.Errorf("Fail:\n  got: %v\n want: %v", req, request)
			}

		})
	}
}
