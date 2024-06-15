package server

import (
	"context"
	"fmt"

	"github.com/sourcegraph/jsonrpc2"
)

// NewHandler creates a handler from the server to handle
// JSON-RPC requests.
func (s *Server) NewHandler() jsonrpc2.Handler {
	return jsonrpc2.HandlerWithError(s.handle)
}

func (s *Server) handle(ctx context.Context, connection *jsonrpc2.Conn, request *jsonrpc2.Request) (any, error) {
	lspContext := NewLSPContext(ctx, connection, request)

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
