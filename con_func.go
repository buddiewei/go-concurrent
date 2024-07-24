package go_concurrent

import "sync"

type conFunc struct {
	fs []func()
	wg *sync.WaitGroup
}

func ConcurrentFunc(f ...func()) *conFunc {
	return &conFunc{
		fs: f,
		wg: new(sync.WaitGroup),
	}
}

func (cf *conFunc) Add(f func()) {
	cf.fs = append(cf.fs, f)
}

func (cf *conFunc) Aggregate(rf func()) {
	n := len(cf.fs)
	cf.wg.Add(n)
	for _, f := range cf.fs {
		go func(f func()) {
			defer cf.wg.Done()
			f()
		}(f)
	}
	cf.wg.Wait()
	if rf == nil {
		return
	}
	rf()
}

func (cf *conFunc) AggregateWithLimit(rf func(), conLimit int) {
	if conLimit <= 0 {
		cf.Aggregate(rf)
	}
	limiter := make(chan any, conLimit)
	n := len(cf.fs)
	cf.wg.Add(n)
	for _, f := range cf.fs {
		go func(f func()) {
			limiter <- 1
			defer cf.wg.Done()
			f()
			<-limiter
		}(f)
	}
	cf.wg.Wait()
	if rf == nil {
		return
	}
	rf()
}
