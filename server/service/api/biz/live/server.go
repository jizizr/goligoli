package live

import (
	"github.com/gwuhaolin/livego/protocol/hls"
	"github.com/gwuhaolin/livego/protocol/rtmp"
)

func init() {
	stream := rtmp.NewRtmpStream()
	var hlsServer *hls.Server
	hlsServer = startHls()
	startHTTPFlv(stream)
	go startRtmp(stream, hlsServer)
}
