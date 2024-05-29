package global

import (
	"github.com/jizizr/goligoli/server/kitex_gen/bullet/bulletservice"
	"github.com/jizizr/goligoli/server/kitex_gen/user/userservice"
)

var (
	UserClient   userservice.Client
	BulletClient bulletservice.Client
)
