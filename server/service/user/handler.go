package main

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jizizr/goligoli/server/common/consts"
	"github.com/jizizr/goligoli/server/common/tools"
	"github.com/jizizr/goligoli/server/kitex_gen/user"
	"github.com/jizizr/goligoli/server/service/user/model"
	"gorm.io/gorm"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	MySqlServiceImpl
}

type MySqlServiceImpl interface {
	GetUserByName(username string) (*model.User, error)
	CreateUser(user *model.User) (string, error)
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	resp = new(user.LoginResponse)
	usr, err := s.GetUserByName(req.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
			resp = nil
			return
		}
		klog.Errorf("MySQL error in GetUserByName: %v", err)
		return
	}
	if usr.Password != tools.Md5(req.Password, "") {
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
	username, err := s.CreateUser(usr)

	if err != nil {
		klog.Error("create user failed, %v", err)
		return
	}
	// if username isn't empty, but err is nil, it means the user already exists,
	// because it's not a real error, it's just a normal situation,
	// return empty token to indicate that the user already exists
	if username != "" {
		return
	}

	resp.Token, err = tools.GenToken(usr.ID)
	if err != nil {
		klog.Error("generate jwt token failed, %v", err)
		return
	}
	return
}
