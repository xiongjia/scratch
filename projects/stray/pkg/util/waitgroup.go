package util

import "sync"

type (
	WaitGroup struct {
		wg sync.WaitGroup
	}
)

func NewWaitGroup() *WaitGroup {
	return &WaitGroup{}
}

func (h *WaitGroup) Go(f func()) {
	h.wg.Add(1)
	go func() {
		defer h.wg.Done()
		f()
	}()
}

func (h *WaitGroup) Wait() {
	h.wg.Wait()
}
