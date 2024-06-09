package main

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jizizr/goligoli/server/common/consts"
	"github.com/jizizr/goligoli/server/kitex_gen/base"
	"github.com/jizizr/goligoli/server/kitex_gen/lottery"
)

// LotteryServiceImpl implements the last service interface defined in the IDL.
type LotteryServiceImpl struct {
	WinnerMySqlServiceImpl
	LotteryMySqlServiceImpl
	RedisServiceImpl
	NsqServiceImpl
}

type WinnerMySqlServiceImpl interface {
	AddWinners(lotteryId int64, users []int64) error
	GetWinners(lotteryId int64) ([]int64, error)
}

type LotteryMySqlServiceImpl interface {
	SetLotteryDB(info *base.Gift) error
	GetLotteryDB(LiveID int64) ([]*base.Gift, error)
	GetLotteryByID(lotteryID int64) (*base.Gift, error)
}

type RedisServiceImpl interface {
	SetLotteryCache(ctx context.Context, info *base.Gift) error
	GetLotteryCache(ctx context.Context, lotteryID int64) (info *base.Gift, err error)
	JoinLotteryCache(ctx context.Context, lotteryID int64, userID int64) error
	GetLiveRoomLotteryCache(ctx context.Context, liveRoomID int64) ([]*base.Gift, error)
	DrawLotteryCache(ctx context.Context, lotteryID int64, count int32) ([]int64, error)
	AddWinnersCache(ctx context.Context, lotteryID int64, winners []int64) error
}

type NsqServiceImpl interface {
	PushToNsq(result []int64) error
}

// SetLottery implements the LotteryServiceImpl interface.
func (s *LotteryServiceImpl) SetLottery(ctx context.Context, req *lottery.SetLotteryRequest) (resp *lottery.SetLotteryResponse, err error) {
	sf, err := snowflake.NewNode(consts.LotterySnowflakeNode)
	if err != nil {
		return nil, err
	}
	req.Gift.Id = sf.Generate().Int64()
	if err := s.SetLotteryCache(ctx, req.Gift); err != nil {
		return nil, err
	}
	if err := s.SetLotteryDB(req.Gift); err != nil {
		return nil, err
	}
	resp = new(lottery.SetLotteryResponse)
	resp.Id = req.Gift.Id
	return
}

// GetLottery implements the LotteryServiceImpl interface.
func (s *LotteryServiceImpl) GetLottery(ctx context.Context, req *lottery.GetLotteryRequest) (resp *lottery.GetLotteryResponse, err error) {
	resp = new(lottery.GetLotteryResponse)
	resp.Gift, err = s.GetLotteryCache(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if resp.Gift != nil {
		return
	}
	resp.Gift, err = s.GetLotteryByID(req.Id)
	if err != nil {
		return nil, err
	}
	if resp.Gift == nil {
		return nil, nil
	}
	return
}

// JoinLottery implements the LotteryServiceImpl interface.
func (s *LotteryServiceImpl) JoinLottery(ctx context.Context, req *lottery.JoinLotteryRequest) (resp *lottery.JoinLotteryResponse, err error) {
	err = s.JoinLotteryCache(ctx, req.Id, req.Uid)
	if err != nil {
		return nil, err
	}
	resp = new(lottery.JoinLotteryResponse)
	resp.Success = true
	return
}

// GetLiveRoomLottery implements the LotteryServiceImpl interface.
func (s *LotteryServiceImpl) GetLiveRoomLottery(ctx context.Context, req *lottery.GetLiveRoomLotteryRequest) (resp *lottery.GetLiveRoomLotteryResponse, err error) {
	resp = new(lottery.GetLiveRoomLotteryResponse)
	resp.Gifts, err = s.GetLiveRoomLotteryCache(ctx, req.LiveId)
	if err != nil {
		return nil, err
	}
	if resp.Gifts != nil {
		return
	}
	resp.Gifts, err = s.GetLotteryDB(req.LiveId)
	if err != nil {
		klog.Errorf("failed to get lottery: %v", err)
	}
	return
}

// DrawLottery implements the LotteryServiceImpl interface.
func (s *LotteryServiceImpl) DrawLottery(ctx context.Context, req *lottery.DrawLotteryRequest) (resp *lottery.DrawLotteryResponse, err error) {
	winners, err := s.DrawLotteryCache(ctx, req.Id, req.Count)
	if err != nil {
		return
	}
	resp = new(lottery.DrawLotteryResponse)
	if winners == nil {
		resp.Msg = "no joiners"
		return
	}
	if err = s.AddWinnersCache(ctx, req.Id, winners); err != nil {
		klog.Errorf("failed to add winners: %v", err)
		return
	}
	err = s.PushToNsq(append([]int64{req.Id}, winners...))
	if err != nil {
		return nil, err
	}
	return
}
