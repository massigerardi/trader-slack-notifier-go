package model

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

var executionWithNoTransactions = Execution{
	Strategy: "test_strategy",
}

const blocksWithNoTransaction = `[{"type":"section","text":{"type":"plain_text","text":"test_strategy"}},{"type":"section","text":{"type":"plain_text","text":"Transactions"}}]`

const jsonWithNoTransaction = `{"strategy":"test_strategy"}`

var executionWithTransactions = Execution{
	Strategy: "test_strategy",
	Transactions: []Transaction{
		{
			Command: "sell",
			Status:  COMPLETED,
		},
		{
			Command: "buy",
			Status:  ABORTED,
		},
	},
}

const blocksWithTransactions = `[{"type":"section","text":{"type":"plain_text","text":"test_strategy"}},{"type":"section","text":{"type":"plain_text","text":"Transactions"}},{"type":"divider"},{"type":"section","text":{"type":"plain_text","text":"sell"}},{"type":"section","text":{"type":"plain_text","text":"Completed"}},{"type":"divider"},{"type":"section","text":{"type":"plain_text","text":"buy"}},{"type":"section","text":{"type":"plain_text","text":"Aborted"}}]`

const jsonWithTransactions = `{"strategy":"test_strategy","transactions":[{"command":"sell","status":"Completed"},{"command":"buy","status":"Aborted"}]}`

func TestExecution_getBlocks(t *testing.T) {
	tests := []struct {
		name   string
		fields Execution
		want   string
	}{
		{
			name:   "executionWithTransactions",
			fields: executionWithNoTransactions,
			want:   blocksWithNoTransaction,
		},
		{
			name:   "executionWithTransactions",
			fields: executionWithTransactions,
			want:   blocksWithTransactions,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := &Execution{
				Strategy:     tt.fields.Strategy,
				Transactions: tt.fields.Transactions,
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

func TestNewExecution(t *testing.T) {
	tests := []struct {
		name string
		args Execution
		want Execution
	}{
		{
			name: "new",
			args: executionWithTransactions,
			want: executionWithTransactions,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewExecution(tt.args.Strategy, tt.args.Transactions...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewExecution() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestNewExecutionMessageRequest(t *testing.T) {
	type args struct {
		Token     string
		Sender    string
		Channel   string
		Receiver  string
		Ephemeral bool
		Type      BodyType
		Message   Execution
	}
	tests := []struct {
		name string
		args args
		want *MessageRequest
	}{
		{
			name: "new",
			args: args{
				Token:    "my_token",
				Sender:   "sender",
				Channel:  "channel",
				Receiver: "receiver",
				Type:     EXECUTION,
				Message:  executionWithTransactions,
			},
			want: &MessageRequest{
				Token:    "my_token",
				Sender:   "sender",
				Channel:  "channel",
				Receiver: "receiver",
				Type:     EXECUTION,
				Message:  &executionWithTransactions,
			},
		},
	}
	for _, tt := range tests {
		data, _ := json.Marshal(tt.want)
		fmt.Printf("%v\n", string(data))
		t.Run(tt.name, func(t *testing.T) {
			if got := NewExecutionMessageRequest(tt.args.Token, tt.args.Sender, tt.args.Channel, tt.args.Receiver, false, tt.args.Message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewExecutionMessageRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExecutionJson(t *testing.T) {
	tests := []struct {
		name      string
		execution Execution
		want      string
	}{
		{
			name:      "executionWithNoTransactions",
			execution: executionWithNoTransactions,
			want:      jsonWithNoTransaction,
		},
		{
			name:      "executionWithTransactions",
			execution: executionWithTransactions,
			want:      jsonWithTransactions,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			execution := &Execution{}
			err := json.Unmarshal([]byte(tt.want), execution)
			if err != nil {
				t.Errorf("json.Unmarshal(%v) gave error %v", tt.want, err)
			}
			if !reflect.DeepEqual(execution, &tt.execution) {
				t.Errorf("Fail Unmarshal:\n  got: %v\n want: %v", execution, &tt.execution)
			}
			data, err := json.Marshal(&tt.execution)
			if err != nil {
				t.Errorf("json.Marshal(%v) gave error %v", tt.execution, err)
			}
			if !reflect.DeepEqual(string(data), tt.want) {
				t.Errorf("Fail Marshal:\n  got: %v\n want: %v", string(data), tt.want)
			}

		})
	}
}

func TestExecutionMessageRequestJson(t *testing.T) {
	tests := []struct {
		name   string
		object MessageRequest
		want   string
	}{
		{
			name:   "empty",
			object: *NewExecutionMessageRequest("token", "sender", "channel", "receiver", false, executionWithNoTransactions),
			want:   `{"token":"token","sender":"sender","channel":"channel","receiver":"receiver","type":"Execution","message":{"strategy":"test_strategy","transactions":[]}}`,
		},
		{
			name:   "test",
			object: *NewExecutionMessageRequest("token", "sender", "channel", "receiver", false, executionWithTransactions),
			want:   `{"token":"token","sender":"sender","channel":"channel","receiver":"receiver","ephemeral":false,"type":"Execution","message":{"strategy":"test_strategy","transactions":[{"command":"sell","status":"Completed"},{"command":"buy","status":"Aborted"}]}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(&tt.object)
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
			if !reflect.DeepEqual(req, &tt.object) {
				t.Errorf("Fail:\n  got: %v\n want: %v", req, &tt.object)
			}

		})
	}
}
