package server

import (
	"context"
	"fmt"

	"github.com/sourcegraph/jsonrpc2"
	"github.com/two-hundred/ls-builder/common"
)

func (s *Server) newHandler() jsonrpc2.Handler {
	return jsonrpc2.HandlerWithError(s.handle)
}

func (s *Server) handle(ctx context.Context, connection *jsonrpc2.Conn, request *jsonrpc2.Request) (any, error) {
	lspContext := &common.LSPContext{
		Method: request.Method,
		Notify: func(method string, params any) error {
			return connection.Notify(ctx, method, params)
		},
		Call: func(method string, params any, result any) error {
			return connection.Call(ctx, method, params, result)
		},
		Context: ctx,
	}

	if request.Params != nil {
		lspContext.Params = *request.Params
	}

	if request.Method == "exit" {
		// Give the attached handler a chance to handle the request before closing the connection
		// but ignore the result.
		s.handler.Handle(lspContext)
		err := connection.Close()
		return nil, err
	}

	// Note: jsonrpc2 will not get to this point if request.Params is not valid JSON,
	// so there is no need to handle jsonrpc2.CodeParseErrors here.
	result, validMethod, validParams, err := s.handler.Handle(lspContext)
	if !validMethod {
		return nil, &jsonrpc2.Error{
			Code:    jsonrpc2.CodeMethodNotFound,
			Message: fmt.Sprintf("method not supported: %s", request.Method),
		}
	} else if !validParams {
		if err == nil {
			return nil, &jsonrpc2.Error{
				Code: jsonrpc2.CodeInvalidParams,
			}
		} else {
			return nil, &jsonrpc2.Error{
				Code:    jsonrpc2.CodeInvalidParams,
				Message: err.Error(),
			}
		}
	} else if err != nil {
		return nil, &jsonrpc2.Error{
			Code:    jsonrpc2.CodeInvalidRequest,
			Message: err.Error(),
		}
	}

	return result, nil
}
