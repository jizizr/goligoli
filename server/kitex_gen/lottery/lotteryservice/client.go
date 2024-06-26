// Code generated by Kitex v0.9.1. DO NOT EDIT.

package lotteryservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	lottery "github.com/jizizr/goligoli/server/kitex_gen/lottery"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	SetLottery(ctx context.Context, req *lottery.SetLotteryRequest, callOptions ...callopt.Option) (r *lottery.SetLotteryResponse, err error)
	GetLottery(ctx context.Context, req *lottery.GetLotteryRequest, callOptions ...callopt.Option) (r *lottery.GetLotteryResponse, err error)
	JoinLottery(ctx context.Context, req *lottery.JoinLotteryRequest, callOptions ...callopt.Option) (r *lottery.JoinLotteryResponse, err error)
	GetLiveRoomLottery(ctx context.Context, req *lottery.GetLiveRoomLotteryRequest, callOptions ...callopt.Option) (r *lottery.GetLiveRoomLotteryResponse, err error)
	DrawLottery(ctx context.Context, req *lottery.DrawLotteryRequest, callOptions ...callopt.Option) (r *lottery.DrawLotteryResponse, err error)
	GetAllUnDrawLottery(ctx context.Context, callOptions ...callopt.Option) (r *lottery.GetAllUnDrawLotteryResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kLotteryServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kLotteryServiceClient struct {
	*kClient
}

func (p *kLotteryServiceClient) SetLottery(ctx context.Context, req *lottery.SetLotteryRequest, callOptions ...callopt.Option) (r *lottery.SetLotteryResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SetLottery(ctx, req)
}

func (p *kLotteryServiceClient) GetLottery(ctx context.Context, req *lottery.GetLotteryRequest, callOptions ...callopt.Option) (r *lottery.GetLotteryResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetLottery(ctx, req)
}

func (p *kLotteryServiceClient) JoinLottery(ctx context.Context, req *lottery.JoinLotteryRequest, callOptions ...callopt.Option) (r *lottery.JoinLotteryResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.JoinLottery(ctx, req)
}

func (p *kLotteryServiceClient) GetLiveRoomLottery(ctx context.Context, req *lottery.GetLiveRoomLotteryRequest, callOptions ...callopt.Option) (r *lottery.GetLiveRoomLotteryResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetLiveRoomLottery(ctx, req)
}

func (p *kLotteryServiceClient) DrawLottery(ctx context.Context, req *lottery.DrawLotteryRequest, callOptions ...callopt.Option) (r *lottery.DrawLotteryResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DrawLottery(ctx, req)
}

func (p *kLotteryServiceClient) GetAllUnDrawLottery(ctx context.Context, callOptions ...callopt.Option) (r *lottery.GetAllUnDrawLotteryResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetAllUnDrawLottery(ctx)
}
