package server

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type TCPTransportTestSuite struct {
	suite.Suite
}

func (s *TCPTransportTestSuite) Test_tcp_transport() {
	port, err := getFreePort()
	s.Require().NoError(err)
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	handler := createCounterHandler()
	server := NewServer(handler, false, logger, nil)
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	go RunTCP(ctx, fmt.Sprintf("localhost:%d", port), server, logger)

	conn, err := waitForTCPServer(fmt.Sprintf("localhost:%d", port), 5*time.Second)
	s.Require().NoError(err)

	clientContainer := createClientHandler()
	clientJSONRPCConn := NewStreamConnection(clientContainer.handler, conn)

	testCountRes := testCountResult{}
	err = clientJSONRPCConn.Call(ctx, "increment", testCountParams{Count: 1}, &testCountRes)
	s.Require().NoError(err)
	s.Require().Equal(2, testCountRes.Count)
}

func TestTCPTransportTestSuite(t *testing.T) {
	suite.Run(t, new(TCPTransportTestSuite))
}
