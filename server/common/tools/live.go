package tools

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jizizr/goligoli/server/kitex_gen/live"
	"github.com/jizizr/goligoli/server/kitex_gen/live/liveservice"
)

func CheckLiveRoom(liveID int64, l *liveservice.Client) bool {
	resp, err := (*l).GetLiveRoomOwner(context.Background(), &live.GetLiveRoomOwnerRequest{LiveId: liveID})
	if err != nil {
		klog.Errorf("get live room owner failed, %v", err)
		return false
	}
	if resp.Owner != 0 {
		return true
	}
	return false
}
