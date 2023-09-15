package waiter

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

// WaitFunc is a function that the Waiter will wait for before shutting down.
type WaitFunc func(ctx context.Context) error

// CleanupFunc is a function that the Waiter will run at the end of the shutdown process.
type CleanupFunc func()

// Waiter is a helper for graceful shutdown.
type Waiter struct {
	ctx          context.Context
	waitFuncs    []WaitFunc
	cleanupFuncs []CleanupFunc
	cancel       context.CancelFunc
}

type waiterCfg struct {
	parentCtx    context.Context
	catchSignals bool
}

// New creates a new Waiter instance.
func New(options ...WaiterOption) *Waiter {
	cfg := waiterCfg{
		parentCtx:    context.Background(),
		catchSignals: false,
	}
	for _, opt := range options {
		opt(&cfg)
	}

	w := &Waiter{
		waitFuncs:    []WaitFunc{},
		cleanupFuncs: []CleanupFunc{},
	}
	w.ctx, w.cancel = context.WithCancel(cfg.parentCtx)
	if cfg.catchSignals {
		w.ctx, w.cancel = signal.NotifyContext(w.ctx, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	}

	return w
}

// Add adds a WaitFunc to the Waiter.
func (w *Waiter) Add(fns ...WaitFunc) {
	w.waitFuncs = append(w.waitFuncs, fns...)
}

// CLeanup adds a CleanupFunc to the Waiter.
func (w *Waiter) Cleanup(fns ...CleanupFunc) {
	w.cleanupFuncs = append(w.cleanupFuncs, fns...)
}

func (w *Waiter) Wait() error {
	g, ctx := errgroup.WithContext(w.ctx)
	g.Go(func() error {
		<-ctx.Done()
		w.cancel()
		return nil
	})

	for _, fn := range w.waitFuncs {
		waitFunc := fn
		g.Go(func() error { return waitFunc(ctx) })
	}

	for _, fn := range w.cleanupFuncs {
		cleanupFunc := fn
		defer cleanupFunc()
	}

	return g.Wait()
}
