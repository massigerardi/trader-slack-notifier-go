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
			obj := &MessageRequest{
				Token:     tt.fields.token,
				Sender:    tt.fields.sender,
				Receiver:  tt.fields.receiver,
				Ephemeral: tt.fields.ephemeral,
				Message:   &tt.fields.message,
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
		want *MessageRequest
	}{
		{
			name: "",
			args: args{
				Token:       "my_token",
				Sender:      "user",
				Channel:     "channel",
				Receiver:    "receiver",
				Ephemeral:   false,
				transaction: transactionSellCompleted,
			},
			want: &MessageRequest{
				Token:     "my_token",
				Sender:    "user",
				Channel:   "channel",
				Receiver:  "receiver",
				Ephemeral: false,
				Message:   &transactionSellCompleted,
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

func TestTransactionMessageRequestJson(t *testing.T) {
	tests := []struct {
		name   string
		object Text
		want   string
	}{
		{
			name:   "test",
			object: example_text,
			want:   `{"token":"token","sender":"sender","channel":"channel","receiver":"receiver","ephemeral":false,"type":"Transaction","message":{"command":"sell","status":"Completed"}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := NewTransactionMessageRequest("token", "sender", "channel", "receiver", false, transactionSellCompleted)
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
