package global

import (
	"github.com/jizizr/goligoli/server/kitex_gen/base"
	"github.com/jizizr/goligoli/server/kitex_gen/bullet/bulletservice"
	"github.com/jizizr/goligoli/server/kitex_gen/push/pushservice"
	"github.com/jizizr/goligoli/server/kitex_gen/user/userservice"
)

var (
	UserClient   userservice.Client
	BulletClient bulletservice.Client
	PushClient   pushservice.Client
)

var Receiver map[int64]map[int64]chan base.Bullet
