package global

import (
	"github.com/jizizr/goligoli/server/kitex_gen/live/liveservice"
	"github.com/jizizr/goligoli/server/kitex_gen/lottery/lotteryservice"
	"github.com/jizizr/goligoli/server/kitex_gen/message/messageservice"
	"github.com/jizizr/goligoli/server/kitex_gen/push/pushservice"
	"github.com/jizizr/goligoli/server/kitex_gen/record/recordservice"
	"github.com/jizizr/goligoli/server/kitex_gen/user/userservice"
)

var (
	UserClient          userservice.Client
	MessageClient       messageservice.Client
	PushClient          pushservice.Client
	ReceiveStreamClient pushservice.StreamClient
	LiveClient          liveservice.Client
	LotteryClient       lotteryservice.Client
	RecordClient        recordservice.Client
)
