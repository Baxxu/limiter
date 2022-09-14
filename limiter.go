package limiter

type Limiter struct {
	limiter chan struct{}
}

func New(maxGoroutinesCount int) Limiter {
	return Limiter{
		limiter: make(chan struct{}, maxGoroutinesCount),
	}
}

func (limiter *Limiter) Add() {
	limiter.limiter <- struct{}{}
}

func (limiter *Limiter) Remove() {
	select {
	case <-limiter.limiter:
	default:
	}
}
