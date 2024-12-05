package models

var (
	TraceIdFieldName       = "trace_id"
	CorrelationIdFieldName = "correlation_id"
	SpanIdFieldName        = "span_id"
	CorrelationCtxKey      = "correlation_ctx_key"
)

type CorrelationCtx struct {
	TraceId string
	SpanId  string
	Id      string
}
