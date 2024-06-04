// Code generated by Kitex v0.9.1. DO NOT EDIT.

package liveservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	live "github.com/jizizr/goligoli/server/kitex_gen/live"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"CreateLiveRoom": kitex.NewMethodInfo(
		createLiveRoomHandler,
		newLiveServiceCreateLiveRoomArgs,
		newLiveServiceCreateLiveRoomResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetLiveRoomOwner": kitex.NewMethodInfo(
		getLiveRoomOwnerHandler,
		newLiveServiceGetLiveRoomOwnerArgs,
		newLiveServiceGetLiveRoomOwnerResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	liveServiceServiceInfo                = NewServiceInfo()
	liveServiceServiceInfoForClient       = NewServiceInfoForClient()
	liveServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return liveServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return liveServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return liveServiceServiceInfoForClient
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
	serviceName := "LiveService"
	handlerType := (*live.LiveService)(nil)
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
		"PackageName": "live",
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

func createLiveRoomHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*live.LiveServiceCreateLiveRoomArgs)
	realResult := result.(*live.LiveServiceCreateLiveRoomResult)
	success, err := handler.(live.LiveService).CreateLiveRoom(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLiveServiceCreateLiveRoomArgs() interface{} {
	return live.NewLiveServiceCreateLiveRoomArgs()
}

func newLiveServiceCreateLiveRoomResult() interface{} {
	return live.NewLiveServiceCreateLiveRoomResult()
}

func getLiveRoomOwnerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*live.LiveServiceGetLiveRoomOwnerArgs)
	realResult := result.(*live.LiveServiceGetLiveRoomOwnerResult)
	success, err := handler.(live.LiveService).GetLiveRoomOwner(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLiveServiceGetLiveRoomOwnerArgs() interface{} {
	return live.NewLiveServiceGetLiveRoomOwnerArgs()
}

func newLiveServiceGetLiveRoomOwnerResult() interface{} {
	return live.NewLiveServiceGetLiveRoomOwnerResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateLiveRoom(ctx context.Context, req *live.CreateLiveRoomRequest) (r *live.CreateLiveRoomResponse, err error) {
	var _args live.LiveServiceCreateLiveRoomArgs
	_args.Req = req
	var _result live.LiveServiceCreateLiveRoomResult
	if err = p.c.Call(ctx, "CreateLiveRoom", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetLiveRoomOwner(ctx context.Context, req *live.GetLiveRoomOwnerRequest) (r *live.GetLiveRoomOwnerResponse, err error) {
	var _args live.LiveServiceGetLiveRoomOwnerArgs
	_args.Req = req
	var _result live.LiveServiceGetLiveRoomOwnerResult
	if err = p.c.Call(ctx, "GetLiveRoomOwner", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
