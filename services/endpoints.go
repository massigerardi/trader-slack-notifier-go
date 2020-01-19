package services

import (
    "context"
    "errors"
    "github.com/go-kit/kit/endpoint"
    "slack-notifier/model"
)

type Endpoints struct {
    ExecutionMessageEndpoint   endpoint.Endpoint
    TransactionMessageEndpoint endpoint.Endpoint
    TextMessageEndpoint        endpoint.Endpoint
}

func MakeExecutionMessageEndpoint(srv Notifier) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(model.ExecutionMessageRequest)
        _, err := srv.SendMessage(ctx, &req)
        if err != nil {
            return model.MessageResponse{Success:false, Err:err.Error()}, nil
        }
        return model.MessageResponse{Success:true}, nil
    }
}

func MakeTransactionMessageEndpoint(srv Notifier) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(model.TransactionMessageRequest)
        _, err := srv.SendMessage(ctx, &req)
        if err != nil {
            return model.MessageResponse{Success:false, Err:err.Error()}, nil
        }
        return model.MessageResponse{Success:true}, nil
    }
}

func MakeTextMessageEndpoint(srv Notifier) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(model.TextMessageRequest)
        _, err := srv.SendMessage(ctx, &req)
        if err != nil {
            return model.MessageResponse{Success:false, Err:err.Error()}, nil
        }
        return model.MessageResponse{Success:true}, nil
    }
}

func (e Endpoints) SendExecutionMessage(ctx context.Context) (model.MessageResponse, error) {
    req := model.ExecutionMessageRequest{}
    resp, err := e.ExecutionMessageEndpoint(ctx, req)
    if err != nil {
        return model.MessageResponse{}, err
    }
    response := resp.(model.MessageResponse)
    if response.Err != "" {
        return response, errors.New(response.Err)
    }
    return response, nil
}

func (e Endpoints) SendTransactionMessage(ctx context.Context) (model.MessageResponse, error) {
    req := model.TransactionMessageRequest{}
    resp, err := e.TransactionMessageEndpoint(ctx, req)
    if err != nil {
        return model.MessageResponse{}, err
    }
    response := resp.(model.MessageResponse)
    if response.Err != "" {
        return response, errors.New(response.Err)
    }
    return response, nil
}

func (e Endpoints) SendTextMessage(ctx context.Context) (model.MessageResponse, error) {
    req := model.TextMessageRequest{}
    resp, err := e.TextMessageEndpoint(ctx, req)
    if err != nil {
        return model.MessageResponse{}, err
    }
    response := resp.(model.MessageResponse)
    if response.Err != "" {
        return response, errors.New(response.Err)
    }
    return response, nil
}

