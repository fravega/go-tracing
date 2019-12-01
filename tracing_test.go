package tracing

import (
	"context"
	"github.com/google/uuid"
	"testing"
)

func TestSetId(t *testing.T) {
	id := uuid.New().String()
	ctx := context.Background()

	ctxWithTrace := SetId(ctx, id)

	if traceId, ok := ctxWithTrace.Value(_traceIdKey).(string); !ok || traceId != id {
		t.Fatalf("traceId from context %s is not equal as previusly set %s", traceId, id)
	}
}

func TestSetId_Empty(t *testing.T) {
	id := ""
	ctx := context.Background()

	ctxWithTrace := SetId(ctx, id)

	if traceId, ok := ctxWithTrace.Value(_traceIdKey).(string); !ok || traceId == "" {
		t.Fatal("traceId from context is empty")
	}
}

func TestSetId_WithValueParent(t *testing.T) {
	id := uuid.New().String()
	parentValue := 1
	parentKey := "parentKey"
	ctx := context.WithValue(context.Background(), parentKey, parentValue)

	ctxWithTrace := SetId(ctx, id)

	if traceId, ok := ctxWithTrace.Value(_traceIdKey).(string); !ok || traceId != id {
		t.Fatalf("traceId from context %s is not equal as previusly set %s", traceId, id)
	}

	if parentValueFromCtx, ok := ctxWithTrace.Value(parentKey).(int); !ok || parentValueFromCtx != parentValue {
		t.Fatalf("parent value from context %d is not equal as expected %d", parentValueFromCtx, parentValue)
	}
}

func TestGetId(t *testing.T) {
	id := uuid.New().String()
	ctx := context.WithValue(context.Background(), _traceIdKey, id)

	idFromCtx := GetId(ctx)

	if id != idFromCtx {
		t.Fatalf("traceId from context %s is not equal as previusly set %s", idFromCtx, id)
	}
}

func TestGetId_Empty(t *testing.T) {
	ctx := context.Background()
	idFromCtx := GetId(ctx)

	if idFromCtx != "" {
		t.Fatalf("traceId from context %s is not empty as expected", idFromCtx)
	}
}