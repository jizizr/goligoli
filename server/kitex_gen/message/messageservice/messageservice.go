// Code generated by Kitex v0.9.1. DO NOT EDIT.

package messageservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	message "github.com/jizizr/goligoli/server/kitex_gen/message"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"AddMessage": kitex.NewMethodInfo(
		addMessageHandler,
		newMessageServiceAddMessageArgs,
		newMessageServiceAddMessageResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetMessage": kitex.NewMethodInfo(
		getMessageHandler,
		newMessageServiceGetMessageArgs,
		newMessageServiceGetMessageResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetHistoryMessages": kitex.NewMethodInfo(
		getHistoryMessagesHandler,
		newMessageServiceGetHistoryMessagesArgs,
		newMessageServiceGetHistoryMessagesResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	messageServiceServiceInfo                = NewServiceInfo()
	messageServiceServiceInfoForClient       = NewServiceInfoForClient()
	messageServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return messageServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return messageServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return messageServiceServiceInfoForClient
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
	serviceName := "MessageService"
	handlerType := (*message.MessageService)(nil)
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
		"PackageName": "message",
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

func addMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*message.MessageServiceAddMessageArgs)

	err := handler.(message.MessageService).AddMessage(ctx, realArg.Req)
	if err != nil {
		return err
	}

	return nil
}
func newMessageServiceAddMessageArgs() interface{} {
	return message.NewMessageServiceAddMessageArgs()
}

func newMessageServiceAddMessageResult() interface{} {
	return message.NewMessageServiceAddMessageResult()
}

func getMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*message.MessageServiceGetMessageArgs)
	realResult := result.(*message.MessageServiceGetMessageResult)
	success, err := handler.(message.MessageService).GetMessage(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMessageServiceGetMessageArgs() interface{} {
	return message.NewMessageServiceGetMessageArgs()
}

func newMessageServiceGetMessageResult() interface{} {
	return message.NewMessageServiceGetMessageResult()
}

func getHistoryMessagesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*message.MessageServiceGetHistoryMessagesArgs)
	realResult := result.(*message.MessageServiceGetHistoryMessagesResult)
	success, err := handler.(message.MessageService).GetHistoryMessages(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMessageServiceGetHistoryMessagesArgs() interface{} {
	return message.NewMessageServiceGetHistoryMessagesArgs()
}

func newMessageServiceGetHistoryMessagesResult() interface{} {
	return message.NewMessageServiceGetHistoryMessagesResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) AddMessage(ctx context.Context, req *message.AddMessageRequest) (err error) {
	var _args message.MessageServiceAddMessageArgs
	_args.Req = req
	var _result message.MessageServiceAddMessageResult
	if err = p.c.Call(ctx, "AddMessage", &_args, &_result); err != nil {
		return
	}
	return nil
}

func (p *kClient) GetMessage(ctx context.Context, req *message.GetMessageRequest) (r *message.GetMessageResponse, err error) {
	var _args message.MessageServiceGetMessageArgs
	_args.Req = req
	var _result message.MessageServiceGetMessageResult
	if err = p.c.Call(ctx, "GetMessage", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetHistoryMessages(ctx context.Context, req *message.GetHistoryMessagesRequest) (r *message.GetHistoryMessagesResponse, err error) {
	var _args message.MessageServiceGetHistoryMessagesArgs
	_args.Req = req
	var _result message.MessageServiceGetHistoryMessagesResult
	if err = p.c.Call(ctx, "GetHistoryMessages", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
