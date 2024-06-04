// Code generated by hertz generator.

package api

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/streaming"
	"github.com/hertz-contrib/sse"
	consts2 "github.com/jizizr/goligoli/server/common/consts"
	"github.com/jizizr/goligoli/server/common/tools"
	base2 "github.com/jizizr/goligoli/server/kitex_gen/base"
	Message "github.com/jizizr/goligoli/server/kitex_gen/message"
	"github.com/jizizr/goligoli/server/kitex_gen/push"
	"github.com/jizizr/goligoli/server/kitex_gen/user"
	"github.com/jizizr/goligoli/server/service/api/biz/errno"
	"github.com/jizizr/goligoli/server/service/api/biz/global"
	"github.com/jizizr/goligoli/server/service/api/biz/model/api"
	"github.com/jizizr/goligoli/server/service/api/biz/model/base"
	"io"
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

// SendMessage .
// @router /message/live [POST]
func SendMessage(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.AddMessageRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.AddMessageResponse)
	uid, _ := tools.GetUID(c)
	sf, err := snowflake.NewNode(consts2.MessageSnowflakeNode)
	if err != nil {
		klog.Errorf("generate snowflake node failed, %v", err)
		return
	}
	resp.ID = sf.Generate().Int64()
	bul := &base2.LiveMessage{
		Type:     req.Type,
		Id:       resp.ID,
		UserId:   uid,
		LiveId:   req.LiveID,
		LiveTime: time.Now().Unix(),
		SendTime: time.Now().Unix(),
		Content:  req.Content,
	}
	err = global.PushClient.PushMessage(ctx, &push.PushMessageRequest{
		Message: bul,
	})
	if err != nil {
		errno.SendResponse(c, consts2.CodeSendMessageFailed, err.Error())
		return
	}
	resp.BaseResp = SuccessBaseResponse()
	c.JSON(consts.StatusOK, resp)
}

// GetHistoryMessages .
// @router /message/history/multi [GET]
func GetHistoryMessages(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetHistoryMessagesRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.GetHistoryMessagesResponse)
	res, err := global.MessageClient.GetHistoryMessages(ctx, &Message.GetHistoryMessagesRequest{
		LiveId:    req.LiveID,
		StartTime: req.StartTime,
		Offset:    req.Offset,
	})

	if err != nil {
		errno.SendResponse(c, consts2.CodeGetHistoryMessagesFailed, err.Error())
		return
	}

	resp.BaseResp = SuccessBaseResponse()
	for _, b := range res.Messages {
		resp.Messages = append(resp.Messages, &base.LiveMessage{
			ID:       b.Id,
			UserID:   b.UserId,
			LiveID:   b.LiveId,
			LiveTime: b.LiveTime,
			SendTime: b.SendTime,
			Content:  b.Content,
		})

	}
	c.JSON(consts.StatusOK, resp)
}

// GetMessageByID .
// @router /message/history/single/ [GET]
func GetMessageByID(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetMessageByIDRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.GetMessageByIDResponse)

	res, err := global.MessageClient.GetMessage(ctx, &Message.GetMessageRequest{
		Id: req.ID,
	})

	if err != nil {
		errno.SendResponse(c, consts2.CodeGetMessageByIDFailed, err.Error())
		return
	}
	resp.BaseResp = SuccessBaseResponse()
	b := res.Message
	if b != nil {
		resp.Message = &base.LiveMessage{
			ID:       b.Id,
			UserID:   b.UserId,
			LiveID:   b.LiveId,
			LiveTime: b.LiveTime,
			SendTime: b.SendTime,
			Content:  b.Content,
		}
	}
	c.JSON(consts.StatusOK, resp)
}

// GetMessageRT .
// @router /message/live [GET]
func GetMessageRT(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.GetMessageRTRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	c.SetStatusCode(http.StatusOK)
	clientStream := sse.NewStream(c)
	uid, _ := tools.GetUID(c)
	serverStream, err := global.ReceiveStreamClient.ReceiveMessage(ctx)
	if err != nil {
		klog.Errorf("receive message failed, %v", err)
		return
	}
	defer serverStream.Close()
	err = serverStream.Send(&push.ReceiveMessageRequest{
		LiveId: req.LiveID,
		UserId: uid,
	})
	for {
		bul, err := serverStream.Recv()
		if err == io.EOF {
			klog.CtxInfof(ctx, "stream closed")
			break
		} else if err != nil {
			klog.CtxErrorf(ctx, "stream error: %v", err)
			streaming.FinishStream(serverStream, err)
			break
		}
		payload, err := sonic.Marshal(bul.Message)
		if err != nil {
			klog.Errorf("marshal error: %v", err)
			break
		}
		err = clientStream.Publish(&sse.Event{
			Event: "message",
			Data:  payload,
		})
		if err != nil {
			streaming.FinishStream(serverStream, err)
			break
		}
	}
}
