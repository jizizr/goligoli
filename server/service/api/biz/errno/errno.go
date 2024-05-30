package errno

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/jizizr/goligoli/server/common/consts"
	"github.com/jizizr/goligoli/server/kitex_gen/base"
	"net/http"
)

type Response struct {
	Base *base.BaseResponse
	Data interface{}
}

func SendResponse(c *app.RequestContext, err consts.RespCode, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Base: &base.BaseResponse{
			StatusCode: int32(err),
			StatusMsg:  err.Msg(),
		},
		Data: data,
	})
}
