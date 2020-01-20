package services

import (
	"context"
	"encoding/json"
	"net/http"
	"slack-notifier/model"
)

func decodeMessageRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.MessageRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return model.MessageRequest{}, err
	}
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
