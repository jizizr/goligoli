// Code generated by hertz generator. DO NOT EDIT.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	router "github.com/jizizr/goligoli/server/service/api/biz/router"
)

// register registers all routers.
func register(r *server.Hertz) {

	router.GeneratedRegister(r)

	customizedRegister(r)
}
