package helper

import (
	"context"
	"net/http"
)

// GetValueFromHTTPCtx 从请求上下文获取指定的key值
func GetValueFromHTTPCtx(r *http.Request, key interface{}) interface{} {
	return r.Context().Value(key)
}

// SetValueToHTTPCtx 将指定的key/val设置到上下文中
func SetValueToHTTPCtx(r *http.Request, key, val interface{}) *http.Request {
	if val == nil {
		return r
	}

	return r.WithContext(context.WithValue(r.Context(), key, val))
}

// GetStringByCtx 从上下文上获取string的key
func GetStringByCtx(ctx context.Context, key string) string {
	val := GetContextValue(ctx, key)
	if val == nil {
		return ""
	}

	str, ok := val.(string)
	if !ok {
		return ""
	}

	return str
}

// GetContextValue 从ctx上获得key/val
func GetContextValue(ctx context.Context, key interface{}) interface{} {
	return ctx.Value(key)
}

// SetContextValue 设置key/val到标准的上下文中
func SetContextValue(ctx context.Context,
	key interface{}, val interface{}) context.Context {
	return context.WithValue(ctx, key, val)
}
