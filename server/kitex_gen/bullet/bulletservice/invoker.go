// Code generated by Kitex v0.9.1. DO NOT EDIT.

package bulletservice

import (
	server "github.com/cloudwego/kitex/server"
	bullet "github.com/jizizr/goligoli/server/kitex_gen/bullet"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler bullet.BulletService, opts ...server.Option) server.Invoker {
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
