package coroutine

import "sync"

type parallel struct {
	callbacks []func() interface{}
	ch        chan struct{}
	limit     int
}

// NewParallel 创建一个并发执行
//
// 	limit 是启用协程的数量，若为0，则不限制协程数
func NewParallel(limit int) *parallel {
	return &parallel{
		callbacks: make([]func() interface{}, 0),
		ch:        make(chan struct{}, limit),
		limit:     limit,
	}
}

func (p *parallel) Add(fn func() interface{}) *parallel {
	p.callbacks = append(p.callbacks, fn)

	return p
}

func (p *parallel) Wait() []interface{} {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	result := make([]interface{}, len(p.callbacks))

	for idx, fn := range p.callbacks {
		wg.Add(1)

		if p.limit > 0 {
			p.ch <- struct{}{}
		}

		go func(fn func() interface{}, idx int) {
			defer wg.Done()

			result[idx] = fn()

			if p.limit > 0 {
				<-p.ch
			}
		}(fn, idx)
	}

	return result
}
