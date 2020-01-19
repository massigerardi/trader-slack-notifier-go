package services

import (
	"context"
	"github.com/jarcoal/httpmock"
	"reflect"
	"slack-notifier/model"
	"testing"
)

func TestNewNotifier(t *testing.T) {
	tests := []struct {
		name string
		want Notifier
	}{
		{
			name: "test",
			want: notifier{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNotifier(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNotifier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_notifier_SendMessage(t *testing.T) {
	type args struct {
		ctx     context.Context
		message model.MessageRequest
	}
	tests := []struct {
		name       string
		args       args
		response   string
		wantResult string
		wantErr    bool
	}{
		{
			name:"success",
			args:args{
				ctx:     nil,
				message: &model.TextMessageRequest{
					Token:    "",
					Sender:   "user",
					Channel:  "",
					Receiver: "channel",
					Message: model.Text{Text:"test"},
				},
			},
			response: "{\"ok\":true}",
			wantResult: "",
			wantErr: false,
		},
		{
			name:"error",
			args:args{
				ctx:     nil,
				message: &model.TextMessageRequest{
					Token:    "",
					Sender:   "user",
					Channel:  "",
					Receiver: "channel",
					Message: model.Text{Text:"test"},
				},
			},
			response: "{\"ok\":false, \"error\":\"error\"}",
			wantResult: "",
			wantErr: true,
		},
	}
	httpmock.Activate()
	for _, tt := range tests {
		responder := httpmock.NewStringResponder(200, tt.response).Once()
		httpmock.RegisterResponder("POST", "https://slack.com/api/chat.postMessage", responder)
		t.Run(tt.name, func(t *testing.T) {
			no := notifier{}
			gotResult, err := no.SendMessage(tt.args.ctx, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("sendMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("sendMessage() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_notifier_SendEphemeralMessage(t *testing.T) {
	type args struct {
		ctx     context.Context
		message model.MessageRequest
	}
	tests := []struct {
		name       string
		args       args
		response   string
		wantResult string
		wantErr    bool
	}{
		{
			name:"success",
			args:args{
				ctx:     nil,
				message: &model.TextMessageRequest{
					Token:    "",
					Sender:   "user",
					Channel:  "",
					Receiver: "channel",
					Message: model.Text{Text:"test"},
				},
			},
			response: "{\"ok\":true}",
			wantResult: "",
			wantErr: false,
		},
		{
			name:"error",
			args:args{
				ctx:     nil,
				message: &model.TextMessageRequest{
					Token:    "",
					Sender:   "user",
					Channel:  "",
					Receiver: "channel",
					Message: model.Text{Text:"test"},
				},
			},
			response: "{\"ok\":false, \"error\":\"error\"}",
			wantResult: "",
			wantErr: true,
		},
	}
	httpmock.Activate()
	for _, tt := range tests {
		responder := httpmock.NewStringResponder(200, tt.response).Once()
		httpmock.RegisterResponder("POST", "https://slack.com/api/chat.postEphemeral", responder)
		t.Run(tt.name, func(t *testing.T) {
			no := notifier{}
			gotResult, err := no.SendEphemeralMessage(tt.args.ctx, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendEphemeralMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("SendEphemeralMessage() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}