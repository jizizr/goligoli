package initialize

import (
	"context"
	"github.com/jizizr/goligoli/server/service/push/config"
	"sync"
)

func InitReciver() {
	liveIds, err := config.LiveClient.GetAllOnlineLiveRoom(context.Background())
	if err != nil {
		panic(err)
	}
	for _, id := range liveIds.LiveIds {
		config.Receiver.Store(id, &sync.Map{})
	}
}
