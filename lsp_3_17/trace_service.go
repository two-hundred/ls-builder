package lsp

import (
	"fmt"
	"sync"

	"github.com/two-hundred/ls-builder/common"
	"go.uber.org/zap"
)

// TraceServerLogger is an interface that provides a way to log messages
// to the client taking into account the current trace configuration.
type TraceServerLogger interface {
	// LogMessage logs a message to the client.
	LogMessage(params LogMessageParams) error
}

// TraceService provides a convenient way to store trace configuration
// and log trace messages to the client where the current shared trace
// configuration is taken into account.
type TraceService struct {
	logger     *zap.Logger
	traceValue TraceValue
	mu         sync.Mutex
}

// NewTraceService creates a new trace service.
func NewTraceService(logger *zap.Logger) *TraceService {
	return &TraceService{
		logger:     logger,
		traceValue: TraceValueOff,
	}
}

// CreateSetTraceHandler creates a new set trace handler function
// that can be configured with a `Handler` instance to handle `$/setTrace`
// notifications from the client and store the trace level in the trace service.
func (s *TraceService) CreateSetTraceHandler() SetTraceHandlerFunc {
	return func(ctx *common.LSPContext, params *SetTraceParams) error {
		s.mu.Lock()
		defer s.mu.Unlock()
		s.traceValue = params.Value
		return nil
	}
}

// Trace logs a message to the client if the provided message type
// is enabled in the current trace configuration.
func (s *TraceService) Trace(serverLogger TraceServerLogger, mType MessageType, message string) error {
	if s.hasTraceMessageType(mType) {
		return serverLogger.LogMessage(LogMessageParams{
			Type:    mType,
			Message: message,
		})
	}
	return nil
}

// GetTraceValue returns the current trace value
// shared between the client and the server.
func (s *TraceService) GetTraceValue() TraceValue {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.traceValue
}

func (s *TraceService) hasTraceMessageType(mType MessageType) bool {
	switch mType {
	case MessageTypeError, MessageTypeWarning, MessageTypeInfo:
		return s.hasTraceLevel(TraceValueMessage)

	case MessageTypeLog:
		return s.hasTraceLevel(TraceValueVerbose)

	default:
		s.logger.Fatal(fmt.Sprintf("unsupported message type: %d", mType))
		return false
	}
}

func (s *TraceService) hasTraceLevel(value TraceValue) bool {
	current := s.GetTraceValue()
	switch current {
	case TraceValueOff:
		return false

	case TraceValueMessage:
		return value == TraceValueMessage

	case TraceValueVerbose:
		return true

	default:
		s.logger.Fatal(fmt.Sprintf("unsupported trace level: %s", current))
		return false
	}
}
