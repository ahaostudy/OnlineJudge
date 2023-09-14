package ctxt

import (
	"context"
	"time"
)

// WithTimeoutContext 超时上下文
func WithTimeoutContext(second time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), second*time.Second)
}
