package domain

import "context"

const ContextRequestIdKey = "requestid"
const ContextCorrelationIdKey = "correlationid"

func WithRequestId(ctx context.Context, requestId string) context.Context {
	return context.WithValue(ctx, ContextRequestIdKey, requestId)
}

func RequestId(ctx context.Context) *string {
	requestId, ok := ctx.Value(ContextRequestIdKey).(string)
	if !ok {
		return nil
	}
	return &requestId
}

func WithCorrelationId(ctx context.Context, correlationId string) context.Context {
	return context.WithValue(ctx, ContextCorrelationIdKey, correlationId)
}

func CorrelationId(ctx context.Context) *string {
	correlationId, ok := ctx.Value(ContextCorrelationIdKey).(string)
	if !ok {
		return nil
	}
	return &correlationId
}
