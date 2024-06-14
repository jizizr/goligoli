package live

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gwuhaolin/livego/configure"
	"github.com/gwuhaolin/livego/protocol/hls"
	"net"
)

func startHls() *hls.Server {
	hlsAddr := configure.Config.GetString("hls_addr")
	hlsListen, err := net.Listen("tcp", hlsAddr)
	if err != nil {
		klog.Fatal(err)
	}

	hlsServer := hls.NewServer()
	go func() {
		defer func() {
			if r := recover(); r != nil {
				klog.Error("HLS server panic: ", r)
			}
		}()
		klog.Info("HLS listen On ", hlsAddr)
		hlsServer.Serve(hlsListen)
	}()
	return hlsServer
}
