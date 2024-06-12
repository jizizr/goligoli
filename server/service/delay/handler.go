package main

import (
	"context"
	delay "github.com/jizizr/goligoli/server/kitex_gen/delay"
)

// DelayTaskServiceImpl implements the last service interface defined in the IDL.
type DelayTaskServiceImpl struct {
	DelayQueueImpl
}

type DelayQueueImpl interface {
	Push(ctx context.Context, req *delay.DelayTaskRequest) error
}

// DelayTask implements the DelayTaskServiceImpl interface.
func (s *DelayTaskServiceImpl) DelayTask(ctx context.Context, req *delay.DelayTaskRequest) (err error) {
	err = s.Push(ctx, req)
	return
}
