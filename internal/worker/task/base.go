package task

import "context"

// BaseTask base task.
type BaseTask struct {
	ctx context.Context
}

// SetCtx set ctx for base task.
func (b *BaseTask) SetCtx(ctx context.Context) {
	b.ctx = ctx
}
