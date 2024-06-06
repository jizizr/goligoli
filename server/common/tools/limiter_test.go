package tools

import (
	"sync"
	"testing"
	"time"
)

import (
	"sync/atomic"
)

func TestLimiter(t *testing.T) {
	// 设置 1 秒 (1 秒 = 1e9 纳秒) 内最多 5 次
	limiter := NewLimiter(time.Second, 5)
	word := "test"
	var wg sync.WaitGroup
	var successCount int32
	var failCount int32

	start := make(chan struct{})

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			if limiter.ShouldSendWord(word) {
				atomic.AddInt32(&successCount, 1)
			} else {
				atomic.AddInt32(&failCount, 1)
			}
		}()
	}

	// 同时启动所有 Goroutine
	close(start)
	wg.Wait()

	if successCount != 5 {
		t.Errorf("Expected 5 successful sends, but got %d", successCount)
	}

	if failCount != 5 {
		t.Errorf("Expected 5 failed sends, but got %d", failCount)
	}
}
