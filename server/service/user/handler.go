package main

import (
	"context"
	"errors"
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jizizr/goligoli/server/common/consts"
	"github.com/jizizr/goligoli/server/common/tools"
	user "github.com/jizizr/goligoli/server/kitex_gen/user"
	"github.com/jizizr/goligoli/server/service/user/model"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	MySqlServiceImpl
}

type MySqlServiceImpl interface {
	GetUserByName(username string) (*model.User, error)
	CreateUser(user *model.User) error
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	resp = new(user.LoginResponse)
	usr, err := s.GetUserByName(req.Username)
	if err != nil {
		klog.Error("get user by name failed, %v", err)
		return
	}
	if usr.Password != tools.Md5(req.Password, "") {
		err = errors.New(consts.ErrPassword)
		return
	}
	resp.Token, err = tools.GenToken(usr.ID)
	return
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	resp = new(user.RegisterResponse)
	sf, err := snowflake.NewNode(consts.UserSnowflakeNode)
	if err != nil {
		klog.Error("generate snowflake node failed, %v", err)
		return
	}
	usr := &model.User{
		ID:       sf.Generate().Int64(),
		Username: req.Username,
		Password: tools.Md5(req.Password, ""),
	}
	err = s.CreateUser(usr)
	if err != nil {
		klog.Error("create user failed, %v", err)
		return
	}
	resp.Token, err = tools.GenToken(usr.ID)
	if err != nil {
		klog.Error("generate jwt token failed, %v", err)
		return
	}
	return
}
