package live

import (
	"crypto/tls"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gwuhaolin/livego/configure"
	"github.com/gwuhaolin/livego/protocol/hls"
	"github.com/gwuhaolin/livego/protocol/rtmp"
	"net"
)

func startRtmp(stream *rtmp.RtmpStream, hlsServer *hls.Server) {
	rtmpAddr := configure.Config.GetString("rtmp_addr")
	isRtmps := configure.Config.GetBool("enable_rtmps")

	var rtmpListen net.Listener
	if isRtmps {
		certPath := configure.Config.GetString("rtmps_cert")
		keyPath := configure.Config.GetString("rtmps_key")
		cert, err := tls.LoadX509KeyPair(certPath, keyPath)
		if err != nil {
			klog.Fatal(err)
		}

		rtmpListen, err = tls.Listen("tcp", rtmpAddr, &tls.Config{
			Certificates: []tls.Certificate{cert},
		})
		if err != nil {
			klog.Fatal(err)
		}
	} else {
		var err error
		rtmpListen, err = net.Listen("tcp", rtmpAddr)
		if err != nil {
			klog.Fatal(err)
		}
	}

	var rtmpServer *rtmp.Server

	if hlsServer == nil {
		rtmpServer = rtmp.NewRtmpServer(stream, nil)
		klog.Info("HLS server disable....")
	} else {
		rtmpServer = rtmp.NewRtmpServer(stream, hlsServer)
		klog.Info("HLS server enable....")
	}

	defer func() {
		if r := recover(); r != nil {
			klog.Error("RTMP server panic: ", r)
		}
	}()
	if isRtmps {
		klog.Info("RTMPS Listen On ", rtmpAddr)
	} else {
		klog.Info("RTMP Listen On ", rtmpAddr)
	}
	rtmpServer.Serve(rtmpListen)
}
