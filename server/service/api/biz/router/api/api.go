// Code generated by hertz generator. DO NOT EDIT.

package api

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	api "github.com/jizizr/goligoli/server/service/api/biz/handler/api"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	root.POST("/login", append(_loginMw(), api.Login)...)
	root.POST("/register", append(_registerMw(), api.Register)...)
	{
		_bullet := root.Group("/bullet", _bulletMw()...)
		_bullet.GET("/live", append(_getbulletrtMw(), api.GetBulletRT)...)
		_bullet.POST("/live", append(_sendbulletMw(), api.SendBullet)...)
		{
			_history := _bullet.Group("/history", _historyMw()...)
			_history.GET("/multi", append(_gethistorybulletsMw(), api.GetHistoryBullets)...)
			_history.GET("/single", append(_getbulletbyidMw(), api.GetBulletByID)...)
		}
	}
}
