// Code generated by Kitex v0.9.1. DO NOT EDIT.

package bulletservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	bullet "github.com/jizizr/goligoli/server/kitex_gen/bullet"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"AddBullet": kitex.NewMethodInfo(
		addBulletHandler,
		newBulletServiceAddBulletArgs,
		newBulletServiceAddBulletResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetBullet": kitex.NewMethodInfo(
		getBulletHandler,
		newBulletServiceGetBulletArgs,
		newBulletServiceGetBulletResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetHistoryBullets": kitex.NewMethodInfo(
		getHistoryBulletsHandler,
		newBulletServiceGetHistoryBulletsArgs,
		newBulletServiceGetHistoryBulletsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	bulletServiceServiceInfo                = NewServiceInfo()
	bulletServiceServiceInfoForClient       = NewServiceInfoForClient()
	bulletServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return bulletServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return bulletServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return bulletServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "BulletService"
	handlerType := (*bullet.BulletService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "bullet",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func addBulletHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*bullet.BulletServiceAddBulletArgs)

	err := handler.(bullet.BulletService).AddBullet(ctx, realArg.Req)
	if err != nil {
		return err
	}

	return nil
}
func newBulletServiceAddBulletArgs() interface{} {
	return bullet.NewBulletServiceAddBulletArgs()
}

func newBulletServiceAddBulletResult() interface{} {
	return bullet.NewBulletServiceAddBulletResult()
}

func getBulletHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*bullet.BulletServiceGetBulletArgs)
	realResult := result.(*bullet.BulletServiceGetBulletResult)
	success, err := handler.(bullet.BulletService).GetBullet(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBulletServiceGetBulletArgs() interface{} {
	return bullet.NewBulletServiceGetBulletArgs()
}

func newBulletServiceGetBulletResult() interface{} {
	return bullet.NewBulletServiceGetBulletResult()
}

func getHistoryBulletsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*bullet.BulletServiceGetHistoryBulletsArgs)
	realResult := result.(*bullet.BulletServiceGetHistoryBulletsResult)
	success, err := handler.(bullet.BulletService).GetHistoryBullets(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBulletServiceGetHistoryBulletsArgs() interface{} {
	return bullet.NewBulletServiceGetHistoryBulletsArgs()
}

func newBulletServiceGetHistoryBulletsResult() interface{} {
	return bullet.NewBulletServiceGetHistoryBulletsResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) AddBullet(ctx context.Context, req *bullet.AddBulletRequest) (err error) {
	var _args bullet.BulletServiceAddBulletArgs
	_args.Req = req
	var _result bullet.BulletServiceAddBulletResult
	if err = p.c.Call(ctx, "AddBullet", &_args, &_result); err != nil {
		return
	}
	return nil
}

func (p *kClient) GetBullet(ctx context.Context, req *bullet.GetBulletRequest) (r *bullet.GetBulletResponse, err error) {
	var _args bullet.BulletServiceGetBulletArgs
	_args.Req = req
	var _result bullet.BulletServiceGetBulletResult
	if err = p.c.Call(ctx, "GetBullet", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetHistoryBullets(ctx context.Context, req *bullet.GetHistoryBulletsRequest) (r *bullet.GetHistoryBulletsResponse, err error) {
	var _args bullet.BulletServiceGetHistoryBulletsArgs
	_args.Req = req
	var _result bullet.BulletServiceGetHistoryBulletsResult
	if err = p.c.Call(ctx, "GetHistoryBullets", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
