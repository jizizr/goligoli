// Code generated by hertz generator.

package api

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/jizizr/goligoli/server/common/middleware"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _loginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _registerMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _messageMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _sendmessageMw() []app.HandlerFunc {
	return []app.HandlerFunc{middleware.JWTAuth()}
}

func _getmessagertMw() []app.HandlerFunc {
	return []app.HandlerFunc{middleware.JWTAuth()}
}

func _historyMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _gethistorymessagesMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getmessagebyidMw() []app.HandlerFunc {
	// your code...
	return nil
}
