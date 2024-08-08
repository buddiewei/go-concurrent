package go_concurrent

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
)

type conErrFunc struct {
	fs []func() error
}

func ConcurrentErrFunc(f ...func() error) *conErrFunc {
	return &conErrFunc{
		fs: f,
	}
}

func (cf *conErrFunc) Add(f func() error) {
	cf.fs = append(cf.fs, f)
}

func (cf *conErrFunc) Aggregate(ctx context.Context, rf func() error) error {
	g, _ := errgroup.WithContext(ctx)
	for _, f := range cf.fs {
		g.Go(f)
	}
	if err := g.Wait(); err != nil {
		return err
	}
	if rf == nil {
		return nil
	}
	return rf()
}

func (cf *conErrFunc) AggregateWithLimit(ctx context.Context, rf func() error, conLimit int) error {
	if conLimit <= 0 {
		return cf.Aggregate(ctx, rf)
	}
	limiter := make(chan any, conLimit)
	g, cancel := errgroup.WithContext(ctx)
	for _, f := range cf.fs {
		tf := f
		g.Go(func() (err error) {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("panic occurred: %v", r)
				}
			}()
			select {
			case limiter <- 1:
			case <-cancel.Done():
				return cancel.Err()
			}
			defer func() { <-limiter }()
			return tf()
		})
	}
	if err := g.Wait(); err != nil {
		return err
	}
	if rf == nil {
		return nil
	}
	return rf()
}
