// Code generated by Kitex v0.9.1. DO NOT EDIT.

package delaytaskservice

import (
	server "github.com/cloudwego/kitex/server"
	delay "github.com/jizizr/goligoli/server/kitex_gen/delay"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler delay.DelayTaskService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
