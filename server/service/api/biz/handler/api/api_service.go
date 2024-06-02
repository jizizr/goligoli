// Code generated by hertz generator.

package api

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hertz-contrib/sse"
	consts2 "github.com/jizizr/goligoli/server/common/consts"
	"github.com/jizizr/goligoli/server/common/tools"
	base2 "github.com/jizizr/goligoli/server/kitex_gen/base"
	"github.com/jizizr/goligoli/server/kitex_gen/bullet"
	"github.com/jizizr/goligoli/server/kitex_gen/push"
	"github.com/jizizr/goligoli/server/kitex_gen/user"
	"github.com/jizizr/goligoli/server/service/api/biz/errno"
	"github.com/jizizr/goligoli/server/service/api/biz/global"
	"github.com/jizizr/goligoli/server/service/api/biz/model/api"
	"github.com/jizizr/goligoli/server/service/api/biz/model/base"
	"net/http"
	"time"
)

func SuccessBaseResponse() *base.BaseResponse {
	return &base.BaseResponse{
		StatusCode: int32(consts2.Success),
		StatusMsg:  consts2.Success.Msg(),
	}
}

// Register .
// @router /register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.RegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.RegisterResponse)
	res, err := global.UserClient.Register(ctx, &user.RegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		errno.SendResponse(c, consts2.CodeRegisterFailed, err.Error())
		return
	}

	if res.Token == "" {
		errno.SendResponse(c, consts2.CodeUserAlreadyExists, nil)
		return
	}

	resp = &api.RegisterResponse{
		BaseResp: SuccessBaseResponse(),
		Token:    res.Token,
	}
	c.JSON(consts.StatusOK, resp)
}

// Login .
// @router /login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.LoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp := new(api.LoginResponse)
	res, err := global.UserClient.Login(ctx, &user.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		errno.SendResponse(c, consts2.CodeLoginFailed, err.Error())
		return
	}

	if res == nil {
		errno.SendResponse(c, consts2.CodeUserNotFound, nil)
		return
	}
	if res.Token == "" {
		errno.SendResponse(c, consts2.CodeWrongPassword, nil)
		return
	}

	resp = &api.LoginResponse{
		BaseResp: SuccessBaseResponse(),
		Token:    res.Token,
	}
	c.JSON(consts.StatusOK, resp)
}

// SendBullet .
// @router /bullet/live [POST]
func SendBullet(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.AddBulletRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.AddBulletResponse)
	uid, _ := tools.GetUID(c)
	sf, err := snowflake.NewNode(consts2.BulletSnowflakeNode)
	if err != nil {
		klog.Errorf("generate snowflake node failed, %v", err)
		return
	}
	resp.BulletID = sf.Generate().Int64()
	bul := &base2.Bullet{
		BulletId: resp.BulletID,
		UserId:   uid,
		LiveId:   req.LiveID,
		LiveTime: req.LiveTime,
		SendTime: time.Now().Unix(),
		Content:  req.Content,
	}
	err = global.PushClient.PushBullet(ctx, &push.PushBulletRequest{
		Bullet: bul,
	})
	if err != nil {
		errno.SendResponse(c, consts2.CodeSendBulletFailed, err.Error())
		return
	}
	resp.BaseResp = SuccessBaseResponse()
	c.JSON(consts.StatusOK, resp)
}

// GetHistoryBullets .
// @router /bullet/history/multi [GET]
func GetHistoryBullets(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetHistoryBulletsRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.GetHistoryBulletsResponse)
	res, err := global.BulletClient.GetHistoryBullets(ctx, &bullet.GetHistoryBulletsRequest{
		LiveId:    req.LiveID,
		StartTime: req.StartTime,
		Offset:    req.Offset,
	})

	if err != nil {
		errno.SendResponse(c, consts2.CodeGetHistoryBulletsFailed, err.Error())
		return
	}

	resp.BaseResp = SuccessBaseResponse()
	for _, b := range res.Bullets {
		resp.Bullets = append(resp.Bullets, &base.Bullet{
			BulletID: b.BulletId,
			UserID:   b.UserId,
			LiveID:   b.LiveId,
			LiveTime: b.LiveTime,
			SendTime: b.SendTime,
			Content:  b.Content,
		})

	}
	c.JSON(consts.StatusOK, resp)
}

// GetBulletByID .
// @router /bullet/history/single/ [GET]
func GetBulletByID(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetBulletByIDRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.GetBulletByIDResponse)

	res, err := global.BulletClient.GetBullet(ctx, &bullet.GetBulletRequest{
		BulletId: req.BulletID,
	})

	if err != nil {
		errno.SendResponse(c, consts2.CodeGetBulletByIDFailed, err.Error())
		return
	}
	resp.BaseResp = SuccessBaseResponse()
	b := res.Bullet
	if b != nil {
		resp.Bullet = &base.Bullet{
			BulletID: b.BulletId,
			UserID:   b.UserId,
			LiveID:   b.LiveId,
			LiveTime: b.LiveTime,
			SendTime: b.SendTime,
			Content:  b.Content,
		}
	}
	c.JSON(consts.StatusOK, resp)
}

// GetBulletRT .
// @router /bullet/live [GET]
func GetBulletRT(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetBulletRTRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	stream := sse.NewStream(c)

	uid, _ := tools.GetUID(c)
	for bul := range global.Receiver[req.LiveID][uid] {
		payload, err := sonic.Marshal(bul)
		if err != nil {
			klog.Errorf("marshal bullet failed, %v", err)
			return
		}
		hlog.CtxInfof(ctx, "message received: %+v", bul)
		event := &sse.Event{
			Event: "bullet",
			Data:  payload,
		}
		c.SetStatusCode(http.StatusOK)
		err = stream.Publish(event)
		if err != nil {
			return
		}
	}
}
