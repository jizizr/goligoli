// Code generated by hertz generator.

package main

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/jizizr/goligoli/server/common/consts"
	"github.com/jizizr/goligoli/server/service/api/initialize"
	"github.com/jizizr/goligoli/server/service/api/initialize/rpc"
)

func main() {
	initialize.InitConfig()
	r, info := initialize.InitEtcd()
	tracer, cfg := hertztracing.NewServerTracer()
	rpc.Init()
	h := server.New(
		tracer,
		server.WithHostPorts(fmt.Sprintf(":%d", consts.ApiServerPort)),
		server.WithRegistry(r, info),
		server.WithHandleMethodNotAllowed(true),
	)
	h.Use(hertztracing.ServerMiddleware(cfg))
	register(h)
	h.Spin()
}
