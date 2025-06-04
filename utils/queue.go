package utils

import (
	"fmt"
	"sync"
)

type Queue struct {
	mu       sync.Mutex
	limit    int
	queue    int
	notifyCh chan string
}

func NewQueue(limit int) *Queue {
	return &Queue{
		limit:    limit,
		queue:    0,
		notifyCh: make(chan string),
	}
}

// Enqueue - 控制排队机制
func (q *Queue) Enqueue() error {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.queue >= q.limit {
		return fmt.Errorf("队列已满，请稍后再试")
	}

	q.queue++
	return nil
}

// Dequeue - 完成当前请求并释放队列空间
func (q *Queue) Dequeue() {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.queue > 0 {
		q.queue--
	}
}

// Notify - 通知队列状态
func (q *Queue) Notify(message string) {
	q.notifyCh <- message
}

// CheckStatus - 获取当前队列状态
func (q *Queue) CheckStatus() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.queue
}
