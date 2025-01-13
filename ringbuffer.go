package traceerr

import (
	"sync"
)

type ringBuffer struct {
	entries []string
	size    int
	pos     int
	mu      sync.Mutex
}

func newRingBuffer(size int) *ringBuffer {
	return &ringBuffer{
		entries: make([]string, size),
		size:    size,
	}
}

func (r *ringBuffer) add(msg string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.entries[r.pos] = msg
	r.pos = (r.pos + 1) % r.size
}

func (r *ringBuffer) dump() []string {
	r.mu.Lock()
	defer r.mu.Unlock()

	res := make([]string, 0, r.size)
	for i := r.pos; i < r.pos+r.size; i++ {
		idx := i % r.size
		if r.entries[idx] != "" {
			res = append(res, r.entries[idx])
		}
	}
	return res
}
