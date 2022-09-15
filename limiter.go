package limiter

type Limiter chan struct{}

func New(maxGoroutinesCount int) Limiter {
	return make(chan struct{}, maxGoroutinesCount)
}

func (limiter *Limiter) Add() {
	*limiter <- struct{}{}
}

func (limiter *Limiter) Remove() {
	select {
	case <-*limiter:
	default:
	}
}

// Queue starts func in goroutine. If goroutines limit reached, wait till some goroutines released
func (limiter *Limiter) Queue(callback func()) {
	limiter.Add()
	go func() {
		defer limiter.Remove()
		callback()
	}()
}
