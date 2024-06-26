package lsp

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/two-hundred/ls-builder/common"
	"go.uber.org/zap"
)

type TraceServiceTestSuite struct {
	suite.Suite
}

func (s *TraceServiceTestSuite) Test_trace_service() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)
	ts := NewTraceService(logger)
	traceHandler := ts.CreateSetTraceHandler()
	ctx := &common.LSPContext{}
	err = traceHandler(ctx, &SetTraceParams{Value: TraceValueVerbose})
	s.Require().NoError(err)
	s.Equal(TraceValueVerbose, ts.traceValue)

	srvLogger := &serverLogger{}
	err = ts.Trace(srvLogger, MessageTypeLog, "test")
	s.Require().NoError(err)

	s.Len(srvLogger.messages, 1)
	s.Equal(MessageTypeLog, srvLogger.messages[0].Type)
	s.Equal("test", srvLogger.messages[0].Message)

	// message is not logged because trace value is off
	err = traceHandler(ctx, &SetTraceParams{Value: TraceValueOff})
	s.Require().NoError(err)
	err = ts.Trace(srvLogger, MessageTypeLog, "test")
	s.Require().NoError(err)
	// Should have not received a message, because trace value is off.
	s.Len(srvLogger.messages, 1)
}

type serverLogger struct {
	messages []LogMessageParams
}

func (s *serverLogger) LogMessage(params LogMessageParams) error {
	s.messages = append(s.messages, params)
	return nil
}

func TestTraceServiceTestSuite(t *testing.T) {
	suite.Run(t, new(TraceServiceTestSuite))
}
