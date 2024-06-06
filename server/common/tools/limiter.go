package tools

import (
	"sort"
	"sync"
	"time"
)

type CircularBuffer struct {
	buffer []int64
	head   int
	tail   int
	size   int
	count  int
	mu     sync.Mutex
}

func NewCircularBuffer(size int) *CircularBuffer {
	return &CircularBuffer{
		buffer: make([]int64, size),
		size:   size,
	}
}

func (cb *CircularBuffer) Add(timestamp int64) {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	cb.buffer[cb.tail] = timestamp
	cb.tail = (cb.tail + 1) % cb.size
	if cb.count < cb.size {
		cb.count++
	} else {
		cb.head = (cb.head + 1) % cb.size
	}
}

func (cb *CircularBuffer) RemoveExpired(expirationTime int64) {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	if cb.count == 0 {
		return
	}

	// 使用二分查找找到第一个未过期的时间戳的位置
	index := sort.Search(cb.count, func(i int) bool {
		return cb.buffer[(cb.head+i)%cb.size] > expirationTime
	})

	// 更新 head 和 count
	cb.head = (cb.head + index) % cb.size
	cb.count -= index
}

func (cb *CircularBuffer) Count() int {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	return cb.count
}

type Limiter struct {
	interval    int64
	limit       int
	wordTracker sync.Map
	bufferSize  int
}

func NewLimiter(interval time.Duration, limit int) *Limiter {
	return &Limiter{
		interval:   interval.Nanoseconds(),
		limit:      limit,
		bufferSize: limit,
	}
}

func (l *Limiter) ShouldSendWord(word string) bool {
	now := time.Now().UnixNano()
	expirationTime := now - l.interval

	bufferIface, _ := l.wordTracker.LoadOrStore(word, NewCircularBuffer(l.bufferSize))
	buffer := bufferIface.(*CircularBuffer)

	// 移除过期时间戳
	buffer.RemoveExpired(expirationTime)

	// 判断是否超过限制
	if buffer.Count() >= l.limit {
		return false
	}

	// 添加当前时间戳
	buffer.Add(now)
	return true
}
