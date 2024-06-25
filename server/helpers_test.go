package server

import (
	"context"
	"encoding/json"
	"net"
	"sync"

	"github.com/sourcegraph/jsonrpc2"
	"github.com/two-hundred/ls-builder/common"
)

func getFreePort() (port int, err error) {
	var a *net.TCPAddr
	if a, err = net.ResolveTCPAddr("tcp", "localhost:0"); err == nil {
		var l *net.TCPListener
		if l, err = net.ListenTCP("tcp", a); err == nil {
			defer l.Close()
			return l.Addr().(*net.TCPAddr).Port, nil
		}
	}
	return
}

func createCounterHandler() common.Handler {
	return common.HandlerFunc(
		func(ctx *common.LSPContext) (r any, validMethod bool, validParams bool, err error) {
			if ctx.Method == "increment" {
				validMethod = true
				countParams := testCountParams{}
				err = json.Unmarshal(ctx.Params, &countParams)
				if err != nil {
					return
				}

				validParams = true
				countResult := testCountResult{
					Count:     countParams.Count + 1,
					PrevCount: countParams.Count,
				}
				r = countResult
			}
			return
		},
	)
}

type clientContainer struct {
	handler                *jsonrpc2.HandlerWithErrorConfigurer
	clientReceivedMessages []*json.RawMessage
	clientReceivedMethods  []string
	mu                     sync.Mutex
}

func createClientHandler() *clientContainer {
	container := &clientContainer{}
	container.handler = jsonrpc2.HandlerWithError(
		func(
			ctx context.Context,
			conn *jsonrpc2.Conn,
			req *jsonrpc2.Request,
		) (interface{}, error) {
			container.mu.Lock()
			defer container.mu.Unlock()
			container.clientReceivedMessages = append(container.clientReceivedMessages, req.Params)
			container.clientReceivedMethods = append(container.clientReceivedMethods, req.Method)
			return nil, nil
		},
	)
	return container
}

type testCountParams struct {
	Count int `json:"count"`
}

type testCountResult struct {
	Count     int `json:"count"`
	PrevCount int `json:"prevCount"`
}
