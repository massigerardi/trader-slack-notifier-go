package services

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"slack-notifier/model"
)

type Endpoints struct {
	MessageEndpoint endpoint.Endpoint
}

func MakeMessageEndpoint(srv Notifier) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.MessageRequest)
		_, err := srv.SendMessage(ctx, req)
		if err != nil {
			return model.MessageResponse{Success: false, Err: err.Error()}, nil
		}
		return model.MessageResponse{Success: true}, nil
	}
}

func (e Endpoints) SendExecutionMessage(ctx context.Context) (model.MessageResponse, error) {
	req := model.MessageRequest{}
	resp, err := e.MessageEndpoint(ctx, req)
	if err != nil {
		return model.MessageResponse{}, err
	}
	response := resp.(model.MessageResponse)
	if response.Err != "" {
		return response, errors.New(response.Err)
	}
	return response, nil
}
