package go_concurrent

import (
	"context"
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
	return rf()
}
