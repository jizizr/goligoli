// Code generated by Kitex v0.9.1. DO NOT EDIT.

package recordservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	record "github.com/jizizr/goligoli/server/kitex_gen/record"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"StartRecord": kitex.NewMethodInfo(
		startRecordHandler,
		newRecordServiceStartRecordArgs,
		newRecordServiceStartRecordResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"StopRecord": kitex.NewMethodInfo(
		stopRecordHandler,
		newRecordServiceStopRecordArgs,
		newRecordServiceStopRecordResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	recordServiceServiceInfo                = NewServiceInfo()
	recordServiceServiceInfoForClient       = NewServiceInfoForClient()
	recordServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return recordServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return recordServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return recordServiceServiceInfoForClient
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
	serviceName := "RecordService"
	handlerType := (*record.RecordService)(nil)
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
		"PackageName": "record",
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

func startRecordHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*record.RecordServiceStartRecordArgs)

	err := handler.(record.RecordService).StartRecord(ctx, realArg.Req)
	if err != nil {
		return err
	}

	return nil
}
func newRecordServiceStartRecordArgs() interface{} {
	return record.NewRecordServiceStartRecordArgs()
}

func newRecordServiceStartRecordResult() interface{} {
	return record.NewRecordServiceStartRecordResult()
}

func stopRecordHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*record.RecordServiceStopRecordArgs)

	err := handler.(record.RecordService).StopRecord(ctx, realArg.Req)
	if err != nil {
		return err
	}

	return nil
}
func newRecordServiceStopRecordArgs() interface{} {
	return record.NewRecordServiceStopRecordArgs()
}

func newRecordServiceStopRecordResult() interface{} {
	return record.NewRecordServiceStopRecordResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) StartRecord(ctx context.Context, req *record.StartRecordRecordRequest) (err error) {
	var _args record.RecordServiceStartRecordArgs
	_args.Req = req
	var _result record.RecordServiceStartRecordResult
	if err = p.c.Call(ctx, "StartRecord", &_args, &_result); err != nil {
		return
	}
	return nil
}

func (p *kClient) StopRecord(ctx context.Context, req *record.StopRecordRecordRequest) (err error) {
	var _args record.RecordServiceStopRecordArgs
	_args.Req = req
	var _result record.RecordServiceStopRecordResult
	if err = p.c.Call(ctx, "StopRecord", &_args, &_result); err != nil {
		return
	}
	return nil
}