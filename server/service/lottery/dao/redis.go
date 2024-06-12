package dao

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jizizr/goligoli/server/kitex_gen/base"
	"github.com/redis/go-redis/v9"
	"strconv"
)

type LotteryRedis struct {
	client *redis.Client
}

func (l *LotteryRedis) SetLotteryCache(ctx context.Context, info *base.Gift) error {
	err := l.client.SAdd(ctx, fmt.Sprintf("lo:%d", info.LiveId), info.Id).Err()
	if err != nil {
		klog.Errorf("failed to set lottery: %v", err)
		return err
	}
	err = l.client.HMSet(ctx, fmt.Sprintf("lo:%d", info.Id), map[string]interface{}{
		"live_id":  info.LiveId,
		"gift":     info.Gift,
		"count":    info.Count,
		"end_time": info.EndTime,
	}).Err()
	if err != nil {
		klog.Errorf("failed to set lottery: %v", err)
		// rollback
		defer func() {
			_ = l.client.SRem(ctx, fmt.Sprintf("lo:%d", info.LiveId), info.Id).Err()
		}()
	}
	return err
}

func (l *LotteryRedis) GetLotteryCache(ctx context.Context, lotteryID int64) (info *base.Gift, err error) {
	res, err := l.client.HGetAll(ctx, fmt.Sprintf("lo:%d", lotteryID)).Result()
	if err != nil {
		klog.Errorf("failed to get lottery: %v", err)
		return
	}
	if len(res) == 0 {
		return nil, nil
	}
	liveId, _ := strconv.ParseInt(res["live_id"], 10, 64)
	count, _ := strconv.Atoi(res["count"])
	endTime, _ := strconv.ParseInt(res["end_time"], 10, 64)
	info = &base.Gift{
		Id:      lotteryID,
		LiveId:  liveId,
		Gift:    res["gift"],
		Count:   int32(count),
		EndTime: endTime,
	}
	return
}

func (l *LotteryRedis) JoinLotteryCache(ctx context.Context, lotteryID int64, userID int64) error {
	err := l.client.SAdd(ctx, fmt.Sprintf("lo:j:%d", lotteryID), userID).Err()
	if err != nil {
		klog.Errorf("failed to join lottery: %v", err)
	}
	return err
}

func (l *LotteryRedis) GetLiveRoomLotteryCache(ctx context.Context, liveRoomID int64) (gifts []*base.Gift, err error) {
	lotteryIdRawSet := l.client.SMembers(ctx, fmt.Sprintf("lo:%d", liveRoomID)).Val()
	for _, lotteryIdRaw := range lotteryIdRawSet {
		lotteryId, _ := strconv.ParseInt(lotteryIdRaw, 10, 64)
		info, err := l.GetLotteryCache(ctx, lotteryId)
		if err != nil {
			klog.Errorf("failed to get lottery: %v", err)
			return nil, err
		}
		if info == nil {
			return nil, nil
		}
		gifts = append(gifts, info)
	}
	return
}

func (l *LotteryRedis) DrawLotteryCache(ctx context.Context, lotteryID int64, count int32) (winners []int64, err error) {
	winnersRaw, err := l.client.SRandMemberN(ctx, fmt.Sprintf("lo:j:%d", lotteryID), int64(count)).Result()
	if err != nil {
		klog.Errorf("failed to draw lottery: %v", err)
		return
	}
	for _, winnerRaw := range winnersRaw {
		winner, _ := strconv.ParseInt(winnerRaw, 10, 64)
		winners = append(winners, winner)
	}
	return
}

func (l *LotteryRedis) AddWinnersCache(ctx context.Context, lotteryID int64, winners []int64) error {
	winnersRaw := make([]interface{}, len(winners))
	for i, winner := range winners {
		winnersRaw[i] = winner
	}
	err := l.client.SAdd(ctx, fmt.Sprintf("lo:w:%d", lotteryID), winnersRaw...).Err()
	if err != nil {
		klog.Errorf("failed to add winners: %v", err)
	}
	return err
}

func NewLotteryRedis(client *redis.Client) *LotteryRedis {
	return &LotteryRedis{client: client}
}
