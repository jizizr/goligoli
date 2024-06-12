package dqueue

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jizizr/goligoli/server/kitex_gen/delay"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type Handler func(taskID int64) error

type DelayQueue struct {
	Name    string
	rd      *redis.Client
	ticker  *time.Ticker
	handler Handler
	stop    chan struct{}
}

func (dq *DelayQueue) Push(ctx context.Context, req *delay.DelayTaskRequest) error {
	err := dq.rd.ZAdd(ctx, dq.Name, redis.Z{
		Score:  float64(req.EndTime),
		Member: req.Id,
	}).Err()
	if err != nil {
		klog.Errorf("push task error: %v", err)
	}
	return err
}

func NewDelayQueue(name string, interval time.Duration, rd *redis.Client, handler Handler) *DelayQueue {
	return &DelayQueue{
		Name:    name + ":dq",
		rd:      rd,
		ticker:  time.NewTicker(interval),
		handler: handler,
		stop:    make(chan struct{}),
	}
}

func (dq *DelayQueue) Start() {
	go func() {
		for {
			select {
			case <-dq.ticker.C:
				dq.consume()
			case <-dq.stop:
				dq.ticker.Stop()
				return
			}
		}
	}()
}

func (dq *DelayQueue) consume() {
	now := time.Now().Unix()
	// 查找并移除到期的任务
	tasks, err := dq.rd.ZRangeByScoreWithScores(context.Background(), dq.Name, &redis.ZRangeBy{
		Min:    "-inf",
		Max:    fmt.Sprintf("%d", now),
		Offset: 0,
		Count:  1,
	}).Result()

	if err != nil {
		klog.Errorf("consume error: %v", err)
		return
	}

	for _, task := range tasks {
		taskID, _ := strconv.ParseInt(task.Member.(string), 10, 64)
		// 处理任务
		if err := dq.handler(taskID); err != nil {
			klog.Errorf("handler task error: %v", err)
			continue
		}
		// 从有序集合中移除任务
		_, err := dq.rd.ZRem(context.Background(), dq.Name, taskID).Result()
		if err != nil {
			klog.Errorf("remove task error: %v", err)
		}
	}
}

func (dq *DelayQueue) Stop() {
	close(dq.stop)
}
