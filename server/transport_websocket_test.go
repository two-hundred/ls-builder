package server

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/sourcegraph/jsonrpc2"
	wsjsonrpc2 "github.com/sourcegraph/jsonrpc2/websocket"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type WebSocketTransportTestSuite struct {
	suite.Suite
}

func (s *WebSocketTransportTestSuite) Test_websocket_transport() {
	port, err := getFreePort()
	s.Require().NoError(err)
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	handler := createCounterHandler()
	server := NewServer(handler, false, logger, nil)
	httpServer := &http.Server{}
	go RunWebSocketServer(fmt.Sprintf("localhost:%d", port), server, logger, httpServer)
	defer httpServer.Shutdown(context.TODO())

	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://localhost:%d", port), nil)
	s.Require().NoError(err)

	ctx := context.Background()
	clientContainer := createClientHandler()
	clientJSONRPCConn := jsonrpc2.NewConn(
		ctx,
		wsjsonrpc2.NewObjectStream(conn),
		clientContainer.handler,
	)

	testCountRes := testCountResult{}
	err = clientJSONRPCConn.Call(ctx, "increment", testCountParams{Count: 1}, &testCountRes)
	s.Require().NoError(err)
	s.Require().Equal(2, testCountRes.Count)
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(WebSocketTransportTestSuite))
}
