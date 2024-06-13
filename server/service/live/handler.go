package main

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jizizr/goligoli/server/common/consts"
	"github.com/jizizr/goligoli/server/kitex_gen/base"
	live "github.com/jizizr/goligoli/server/kitex_gen/live"
	"time"
)

// LiveServiceImpl implements the last service interface defined in the IDL.
type LiveServiceImpl struct {
	MySqlServiceImpl
	RedisServiceImpl
}

type MySqlServiceImpl interface {
	AddLiveRoom(room *base.Room) (int64, error)
	DeleteLiveRoom(id int64) error
	GetLiveRoomByID(id int64) (*base.Room, error)
	GetLiveRoomOwnerByID(id int64) (int64, error)
}

type RedisServiceImpl interface {
	AddLiveRoomCache(ctx context.Context, room *base.Room) error
	DeleteLiveRoomCache(ctx context.Context, id int64) error
	GetLiveRoomCache(ctx context.Context, id int64) (*base.Room, error)
	GetLiveRoomOwnerCache(ctx context.Context, id int64) (int64, error)
}

// CreateLiveRoom implements the LiveServiceImpl interface.
func (s *LiveServiceImpl) CreateLiveRoom(ctx context.Context, req *live.CreateLiveRoomRequest) (resp *live.CreateLiveRoomResponse, err error) {
	resp = new(live.CreateLiveRoomResponse)
	sf, err := snowflake.NewNode(consts.LiveSnowflakeNode)
	if err != nil {
		return
	}
	req.Room.LiveId = sf.Generate().Int64()
	req.Room.StartTime = time.Now().Unix()
	id, err := s.AddLiveRoom(req.Room)
	if err != nil {
		return
	}
	if id != 0 {
		return
	}
	resp.LiveId = req.Room.LiveId
	if err := s.AddLiveRoomCache(ctx, req.Room); err != nil {
		klog.Errorf("AddLiveRoomCache failed: %v", err)
	}
	return
}

// GetLiveRoomOwner implements the LiveServiceImpl interface.
func (s *LiveServiceImpl) GetLiveRoomOwner(ctx context.Context, req *live.GetLiveRoomOwnerRequest) (resp *live.GetLiveRoomOwnerResponse, err error) {
	resp = new(live.GetLiveRoomOwnerResponse)
	resp.Owner, err = s.GetLiveRoomOwnerCache(ctx, req.LiveId)
	if err != nil {
		klog.Errorf("GetLiveRoomOwnerCache failed: %v", err)
		return
	}
	if resp.Owner != 0 {
		return
	}
	resp.Owner, err = s.GetLiveRoomOwnerByID(req.LiveId)
	if err != nil {
		klog.Errorf("GetLiveRoomOwnerByID failed: %v", err)
		return
	}
	return
}

// GetLiveRoom implements the LiveServiceImpl interface.
func (s *LiveServiceImpl) GetLiveRoom(ctx context.Context, req *live.GetLiveRoomRequest) (resp *live.GetLiveRoomResponse, err error) {
	resp = new(live.GetLiveRoomResponse)
	resp.Room, err = s.GetLiveRoomCache(ctx, req.LiveId)
	if err != nil {
		klog.Errorf("GetLiveRoomCache failed: %v", err)
		return
	}
	if resp.Room != nil {
		return
	}
	resp.Room, err = s.GetLiveRoomByID(req.LiveId)
	if err != nil {
		klog.Errorf("GetLiveRoomByID failed: %v", err)
	}
	if resp.Room != nil {
		if err := s.AddLiveRoomCache(ctx, resp.Room); err != nil {
			klog.Errorf("AddLiveRoomCache failed: %v", err)
		}
	}
	return
}
