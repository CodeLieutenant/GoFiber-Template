package constants

type ContextKey string

const (
	RequestIDContextKey          ContextKey = "request_id"
	CancelFuncContextKey         ContextKey = "cancel"
	CancelWillBeCalledContextKey ContextKey = "cancelFnWillBeCalled"
	ContainerContextKey          ContextKey = "container"
	CancelContextKey             ContextKey = "cancel"
)
