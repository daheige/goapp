package logic

import (
	"context"
)

// BaseLogic 基本logic
// Ctx 标准上下文，如果是gin.Context就是ctx.Request.Context(),如果是grpc就是ctx
type BaseLogic struct {
	Ctx context.Context
}

// SetCtx set ctx
func (b *BaseLogic) SetCtx(ctx context.Context) {
	b.Ctx = ctx
}
