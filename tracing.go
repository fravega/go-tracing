package tracing

import (
	"context"
	"github.com/google/uuid"
)

type traceId string
var _traceIdKey = traceId("traceId")

func SetId(ctx context.Context, id string) context.Context {
	if id == "" {
		id = uuid.New().String()
	}
	return context.WithValue(ctx, _traceIdKey, id)
}

func GetId(ctx context.Context) string {
	if id, ok := ctx.Value(_traceIdKey).(string); ok {
		return id
	}
	return ""
}