package main

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/jizizr/goligoli/server/common/consts"
	"github.com/jizizr/goligoli/server/kitex_gen/base"
	live "github.com/jizizr/goligoli/server/kitex_gen/live"
)

// LiveServiceImpl implements the last service interface defined in the IDL.
type LiveServiceImpl struct {
	MySqlServiceImpl
}

type MySqlServiceImpl interface {
	AddLiveRoom(room *base.Room) (int64, error)
	DeleteLiveRoom(id int64) error
	GetLiveRoomByID(id int64) (*base.Room, error)
}

// CreateLiveRoom implements the LiveServiceImpl interface.
func (s *LiveServiceImpl) CreateLiveRoom(ctx context.Context, req *live.CreateLiveRoomRequest) (resp *live.CreateLiveRoomResponse, err error) {
	resp = new(live.CreateLiveRoomResponse)
	sf, err := snowflake.NewNode(consts.UserSnowflakeNode)
	if err != nil {
		return
	}
	req.Room.LiveId = sf.Generate().Int64()
	id, err := s.AddLiveRoom(req.Room)
	if err != nil {
		return
	}
	if id != 0 {
		return
	}
	resp.LiveId = req.Room.LiveId
	return
}

// GetLiveRoomOwner implements the LiveServiceImpl interface.
func (s *LiveServiceImpl) GetLiveRoomOwner(ctx context.Context, req *live.GetLiveRoomOwnerRequest) (resp *live.GetLiveRoomOwnerResponse, err error) {
	resp = new(live.GetLiveRoomOwnerResponse)
	room, err := s.GetLiveRoomByID(req.LiveId)
	if err != nil {
		return
	}
	resp.Owner = room.Owner
	return
}
