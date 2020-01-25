package slack

import (
	"github.com/jarcoal/httpmock"
	"github.com/massigerardi/trader-slack-notifier-go/model"
	slackLib "github.com/nlopes/slack"
	"reflect"
	"testing"
)

func Test_get(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name     string
		args     args
		want     *slackLib.Client
		wantSize int
	}{
		{
			name:     "found",
			args:     args{token: "my_token"},
			want:     slackLib.New("my_token"),
			wantSize: 1,
		},
		{
			name:     "notFound",
			args:     args{token: "my_token2"},
			want:     slackLib.New("my_token2"),
			wantSize: 2,
		},
	}
	get("my_token") //pre-populate the clients
	if len(clients) != 1 {
		t.Errorf("Error while inserting first client")
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := get(tt.args.token); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
			if len(clients) != tt.wantSize {
				t.Errorf("len(clients) = %v, want %v", len(clients), tt.wantSize)
			}
		})
	}
}

func Test_SendMessage(t *testing.T) {
	type fields struct {
		client *slackLib.Client
	}
	type args struct {
		messageRequest model.MessageRequest
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		url        string
		response   string
		wantResult string
		wantErr    bool
	}{
		{
			name: "ephemeral",
			fields: fields{
				client: slackLib.New("my_token"),
			},
			args:       args{*model.NewTextMessageRequest("my_token", "sender", "channel", "receiver", true, model.NewText("test"))},
			response:   `{"ok":true}`,
			url:        "https://slack.com/api/chat.postEphemeral",
			wantErr:    false,
			wantResult: "",
		},
		{
			name: "channel",
			fields: fields{
				client: slackLib.New("my_token"),
			},
			args:       args{*model.NewTextMessageRequest("my_token", "sender", "channel", "receiver", false, model.NewText("test"))},
			response:   "{\"ok\":true}",
			url:        "https://slack.com/api/chat.postMessage",
			wantResult: "",
			wantErr:    false,
		},
		{
			name: "ephemeral_error",
			fields: fields{
				client: slackLib.New("my_token"),
			},
			args:       args{*model.NewTextMessageRequest("my_token", "sender", "channel", "receiver", true, model.NewText("test"))},
			response:   `{"ok":false, "error":"error"}`,
			url:        "https://slack.com/api/chat.postEphemeral",
			wantErr:    true,
			wantResult: "",
		},
		{
			name: "channel_error",
			fields: fields{
				client: slackLib.New("my_token"),
			},
			args:       args{*model.NewTextMessageRequest("my_token", "sender", "channel", "receiver", false, model.NewText("test"))},
			response:   `{"ok":false, "error":"error"}`,
			url:        "https://slack.com/api/chat.postMessage",
			wantResult: "",
			wantErr:    true,
		},
	}
	httpmock.Activate()
	for _, tt := range tests {
		responder := httpmock.NewStringResponder(200, tt.response).Once()
		httpmock.RegisterResponder("POST", tt.url, responder)
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := SendMessage(tt.args.messageRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("SendMessage() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
