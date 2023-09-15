package waiter

import "context"

// WaiterOption configures how we set up the Waiter.
type WaiterOption func(c *waiterCfg)

// ParentContext sets the parent context for the Waiter.
//
// The default is context.Background().
func ParentContext(ctx context.Context) WaiterOption {
	return func(c *waiterCfg) {
		c.parentCtx = ctx
	}
}

// CatchSignals sets up the Waiter to catch SIGINT, SIGTERM, and SIGQUIT.
func CatchSignals() WaiterOption {
	return func(c *waiterCfg) {
		c.catchSignals = true
	}
}
