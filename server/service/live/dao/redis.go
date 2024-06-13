package dao

import (
	"context"
	"fmt"
	"github.com/jizizr/goligoli/server/kitex_gen/base"
	"github.com/redis/go-redis/v9"
	"strconv"
)

type LiveRedis struct {
	rd *redis.Client
}

func (l *LiveRedis) AddLiveRoomCache(ctx context.Context, room *base.Room) error {
	return l.rd.HMSet(ctx, fmt.Sprintf("live:%d", room.LiveId), map[string]interface{}{
		"room_name":    room.RoomName,
		"introduction": room.Introduction,
		"owner":        room.Owner,
		"cover":        room.Cover,
		"start_time":   room.StartTime,
		"is_live":      room.IsLive,
	}).Err()
}

func (l *LiveRedis) DeleteLiveRoomCache(ctx context.Context, id int64) error {
	return l.rd.Del(ctx, fmt.Sprintf("live:%d", id)).Err()
}

func (l *LiveRedis) GetLiveRoomCache(ctx context.Context, id int64) (*base.Room, error) {
	res, err := l.rd.HGetAll(ctx, fmt.Sprintf("live:%d", id)).Result()
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, nil
	}
	owner, _ := strconv.ParseInt(res["owner"], 10, 64)
	startTime, _ := strconv.ParseInt(res["start_time"], 10, 64)
	return &base.Room{
		LiveId:       id,
		RoomName:     res["room_name"],
		Introduction: res["introduction"],
		Owner:        owner,
		Cover:        res["cover"],
		StartTime:    startTime,
		IsLive:       res["is_live"] == "true",
	}, nil
}

func (l *LiveRedis) GetLiveRoomOwnerCache(ctx context.Context, id int64) (int64, error) {
	owner, err := l.rd.HGet(ctx, fmt.Sprintf("live:%d", id), "owner").Int64()
	if err == redis.Nil {
		return 0, nil
	}
	return owner, err
}

func NewLiveRedis(rd *redis.Client) *LiveRedis {
	return &LiveRedis{rd: rd}
}
