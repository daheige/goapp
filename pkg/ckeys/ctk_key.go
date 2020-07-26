package ckeys

import "context"

// CtxKey ctx key type.
type CtxKey struct {
	name string
}

// String return name string.
func (c CtxKey) String() string {
	return c.name
}

// GetStringByCtxKey get string by CtxKey
func GetStringByCtxKey(ctx context.Context, key CtxKey) string {
	val := ctx.Value(key)
	if val == nil {
		return ""
	}

	str, ok := val.(string)
	if !ok {
		return ""
	}

	return str
}
