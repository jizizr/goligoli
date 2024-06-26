// Code generated by Kitex v0.9.1. DO NOT EDIT.

package delaytaskservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	delay "github.com/jizizr/goligoli/server/kitex_gen/delay"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"delayTask": kitex.NewMethodInfo(
		delayTaskHandler,
		newDelayTaskServiceDelayTaskArgs,
		newDelayTaskServiceDelayTaskResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	delayTaskServiceServiceInfo                = NewServiceInfo()
	delayTaskServiceServiceInfoForClient       = NewServiceInfoForClient()
	delayTaskServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return delayTaskServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return delayTaskServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return delayTaskServiceServiceInfoForClient
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
	serviceName := "DelayTaskService"
	handlerType := (*delay.DelayTaskService)(nil)
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
		"PackageName": "delay",
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

func delayTaskHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*delay.DelayTaskServiceDelayTaskArgs)

	err := handler.(delay.DelayTaskService).DelayTask(ctx, realArg.Req)
	if err != nil {
		return err
	}

	return nil
}
func newDelayTaskServiceDelayTaskArgs() interface{} {
	return delay.NewDelayTaskServiceDelayTaskArgs()
}

func newDelayTaskServiceDelayTaskResult() interface{} {
	return delay.NewDelayTaskServiceDelayTaskResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) DelayTask(ctx context.Context, req *delay.DelayTaskRequest) (err error) {
	var _args delay.DelayTaskServiceDelayTaskArgs
	_args.Req = req
	var _result delay.DelayTaskServiceDelayTaskResult
	if err = p.c.Call(ctx, "delayTask", &_args, &_result); err != nil {
		return
	}
	return nil
}
