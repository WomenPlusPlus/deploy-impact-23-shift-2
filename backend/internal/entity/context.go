package entity

type ContextKey string

var (
	ContextKeyKind   ContextKey = "X-Kind"
	ContextKeyRole   ContextKey = "X-Role"
	ContextKeyEmail  ContextKey = "X-Email"
	ContextKeyUserId ContextKey = "X-UserID"
	ContextKeyToken  ContextKey = "X-Token"
)
