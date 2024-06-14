package live

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gwuhaolin/livego/configure"
	"github.com/gwuhaolin/livego/protocol/httpflv"
	"github.com/gwuhaolin/livego/protocol/rtmp"
	"net"
)

func startHTTPFlv(stream *rtmp.RtmpStream) {
	httpflvAddr := configure.Config.GetString("httpflv_addr")

	flvListen, err := net.Listen("tcp", httpflvAddr)
	if err != nil {
		klog.Fatal(err)
	}

	hdlServer := httpflv.NewServer(stream)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				klog.Error("HTTP-FLV server panic: ", r)
			}
		}()
		klog.Info("HTTP-FLV listen On ", httpflvAddr)
		hdlServer.Serve(flvListen)
	}()
}
