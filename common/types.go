package common

import (
	"context"
	"encoding/json"
)

type NotifyFunc func(method string, params any) error
type CallFunc func(method string, params any, result any) error

type LSPContext struct {
	Method  string
	Params  json.RawMessage
	Notify  NotifyFunc
	Call    CallFunc
	Context context.Context
}

type Handler interface {
	Handle(ctx *LSPContext) (r any, validMethod bool, validParams bool, err error)
}
