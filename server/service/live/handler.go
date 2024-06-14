package main

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gwuhaolin/livego/configure"
	"github.com/jizizr/goligoli/server/common/consts"
	"github.com/jizizr/goligoli/server/kitex_gen/base"
	live "github.com/jizizr/goligoli/server/kitex_gen/live"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

// LiveServiceImpl implements the last service interface defined in the IDL.
type LiveServiceImpl struct {
	MySqlServiceImpl
	RedisServiceImpl
}

type MySqlServiceImpl interface {
	AddLiveRoom(room *base.Room) (int64, error)
	TagStopLiveRoom(id int64) error
	GetLiveRoomByID(id int64) (*base.Room, error)
	GetOnlineLiveRooms() ([]int64, error)
}

type RedisServiceImpl interface {
	AddLiveRoomCache(ctx context.Context, room *base.Room) error
	DeleteLiveRoomCache(ctx context.Context, id int64) error
	StopLive(ctx context.Context, id int64) error
	GetLiveRoomCache(ctx context.Context, id int64) (*base.Room, error)
	GetLiveRoomOwnerCache(ctx context.Context, id int64) (int64, error)
	GetLiveRoomIsLiveCache(ctx context.Context, id int64) (bool, error)
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
	req.Room.IsLive = true
	id, err := s.AddLiveRoom(req.Room)
	if err != nil {
		return
	}
	if id != 0 {
		return
	}
	resp.LiveId = req.Room.LiveId
	res, err := s.GetLiveRoomKey(ctx, &live.GetLiveRoomKeyRequest{LiveId: req.Room.LiveId})
	if err != nil {
		klog.Errorf("GetLiveRoomKey failed: %v", err)
	}
	resp.Key = res.Key
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
	room, err := s.GetLiveRoomByID(req.LiveId)
	if err != nil {
		klog.Errorf("GetLiveRoomOwnerByID failed: %v", err)
		return
	}
	if room == nil {
		return
	}
	_ = s.AddLiveRoomCache(ctx, room)
	resp.Owner = room.Owner
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

// StopLiveRoom implements the LiveServiceImpl interface.
func (s *LiveServiceImpl) StopLiveRoom(ctx context.Context, req *live.StopLiveRoomRequest) (err error) {
	// 先更改数据库状态，再删除缓存
	// 保证数据一致性
	err = s.TagStopLiveRoom(req.LiveId)
	if err != nil {
		klog.Errorf("TagStopLiveRoom failed: %v", err)
		return
	}
	if err := s.DeleteLiveRoomCache(ctx, req.LiveId); err != nil {
		klog.Errorf("DeleteLiveRoomCache failed: %v", err)
	}
	if err := s.StopLive(ctx, req.LiveId); err != nil {
		klog.Errorf("StopLive failed: %v", err)
	}
	return
}

// GetLiveRoomKey implements the LiveServiceImpl interface.
func (s *LiveServiceImpl) GetLiveRoomKey(ctx context.Context, req *live.GetLiveRoomKeyRequest) (resp *live.GetLiveRoomKeyResponse, err error) {
	msg, err := configure.RoomKeys.GetKey(strconv.FormatInt(req.LiveId, 10))
	if err != nil {
		klog.Errorf("GetLiveRoomKey failed: %v", err)
	}
	resp = &live.GetLiveRoomKeyResponse{
		Key: msg,
	}
	return
}

// GetLiveRoomStatus implements the LiveServiceImpl interface.
func (s *LiveServiceImpl) GetLiveRoomStatus(ctx context.Context, req *live.GetLiveRoomStatusRequest) (resp *live.GetLiveRoomStatusResponse, err error) {
	resp = new(live.GetLiveRoomStatusResponse)
	resp.IsLive, err = s.GetLiveRoomIsLiveCache(ctx, req.LiveId)
	if err == nil {
		return
	}
	if err != redis.Nil {
		klog.Errorf("GetLiveRoomIsLiveCache failed: %v", err)
		return
	}
	room, err := s.GetLiveRoomByID(req.LiveId)
	if err != nil {
		klog.Errorf("GetLiveRoomByID failed: %v", err)
		return
	}
	if room == nil {
		return nil, nil
	}
	resp.IsLive = room.IsLive
	_ = s.AddLiveRoomCache(ctx, room)
	return
}

// GetAllOnlineLiveRoom implements the LiveServiceImpl interface.
func (s *LiveServiceImpl) GetAllOnlineLiveRoom(ctx context.Context) (resp *live.GetAllOnlineLiveRoomResponse, err error) {
	resp = new(live.GetAllOnlineLiveRoomResponse)
	resp.LiveIds, err = s.GetOnlineLiveRooms()
	return
}
