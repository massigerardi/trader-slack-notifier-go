package services

import (
    "context"
    "encoding/json"
    "net/http"
    "slack-notifier/model"
)

func decodeExecutionMessageRequest(ctx context.Context, r *http.Request) (interface{}, error) {
    var req model.ExecutionMessageRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        return model.ExecutionMessageRequest{}, err
    }
    return req, nil
}

func decodeTransactionMessageRequest(ctx context.Context, r *http.Request) (interface{}, error) {
    var req model.TransactionMessageRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        return model.TransactionMessageRequest{}, err
    }
    return req, nil
}

func decodeTextMessageRequest(ctx context.Context, r *http.Request) (interface{}, error) {
    var req model.TextMessageRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        return model.TextMessageRequest{}, err
    }
    return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
    return json.NewEncoder(w).Encode(response)
}
