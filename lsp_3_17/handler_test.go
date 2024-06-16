package lsp

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/two-hundred/ls-builder/common"
	"github.com/two-hundred/ls-builder/server"
	"go.uber.org/zap"
)

type HandlerTestSuite struct {
	suite.Suite
}

func (s *HandlerTestSuite) Test_calls_cancel_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	callChan := make(chan *CancelParams, 1)
	serverHandler := NewHandler(
		WithCancelRequestHandler(
			func(ctx *common.LSPContext, params *CancelParams) error {
				callChan <- params
				return nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)
	clientDispatcher := NewDispatcher(clientLSPContext)

	cancelID := "test-request-id"
	cancelParams := CancelParams{
		ID: &IntOrString{StrVal: &cancelID},
	}
	err = clientDispatcher.CancelRequest(cancelParams)
	s.Require().NoError(err)

	select {
	case <-ctx.Done():
		s.Fail("timeout")
	case receivedParams := <-callChan:
		s.Require().Equal(cancelParams, *receivedParams)
	}
}

func (s *HandlerTestSuite) Test_calls_progress_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	callChan := make(chan *ProgressParams, 1)
	serverHandler := NewHandler(
		WithProgressHandler(
			func(ctx *common.LSPContext, params *ProgressParams) error {
				callChan <- params
				return nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)
	clientDispatcher := NewDispatcher(clientLSPContext)

	progressToken := "test-progress-token"
	progressParams := ProgressParams{
		Token: &ProgressToken{StrVal: &progressToken},
	}
	err = clientDispatcher.Progress(progressParams)
	s.Require().NoError(err)

	select {
	case <-ctx.Done():
		s.Fail("timeout")
	case receivedParams := <-callChan:
		s.Require().Equal(progressParams, *receivedParams)
	}
}

func (s *HandlerTestSuite) Test_calls_initialize_request_handler_and_sets_initialized_state() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	version := "0.0.1"
	openClose := true
	initializeResult := InitializeResult{
		Capabilities: ServerCapabilities{
			PositionEncoding: PositionEncodingKindUTF8,
			TextDocumentSync: TextDocumentSyncOptions{
				OpenClose: &openClose,
				Change:    &TextDocumentSyncKindIncremental,
			},
		},
		ServerInfo: &InitializeResultServerInfo{
			Name:    "test-language-server",
			Version: &version,
		},
	}

	serverHandler := NewHandler(
		WithInitializeHandler(
			func(ctx *common.LSPContext, params *InitializeParams) (any, error) {
				return initializeResult, nil
			},
		),
	)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	// Assert before the LSP initialisation process.
	s.Require().False(serverHandler.IsInitialized())

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	willSave := true
	didSave := true
	initializeParams := InitializeParams{
		ClientInfo: &InitializeClientInfo{
			Name:    "test-client",
			Version: "0.1.0",
		},
		Capabilities: ClientCapabilities{
			TextDocument: &TextDocumentClientCapabilities{
				Synchronization: &TextDocumentSyncClientCapabilities{
					WillSave: &willSave,
					DidSave:  &didSave,
				},
			},
			General: &GeneralClientCapabilities{
				PositionEncodings: []PositionEncodingKind{PositionEncodingKindUTF8},
			},
		},
	}

	returnedResult := InitializeResult{}
	err = clientLSPContext.Call(MethodInitialize, initializeParams, &returnedResult)
	s.Require().NoError(err)
	s.Require().Equal(initializeResult, returnedResult)
	s.Require().True(serverHandler.IsInitialized())
}

func (s *HandlerTestSuite) Test_calls_initialized_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	callChan := make(chan *InitializedParams, 1)
	serverHandler := NewHandler(
		WithInitializedHandler(
			func(ctx *common.LSPContext, params *InitializedParams) error {
				callChan <- params
				return nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	initializedParams := InitializedParams{}
	err = clientLSPContext.Notify(MethodInitialized, initializedParams)
	s.Require().NoError(err)

	select {
	case <-ctx.Done():
		s.Fail("timeout")
	case receivedParams := <-callChan:
		s.Require().Equal(initializedParams, *receivedParams)
	}
}

func (s *HandlerTestSuite) Test_calls_set_trace_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	callChan := make(chan *SetTraceParams, 1)
	serverHandler := NewHandler(
		WithSetTraceHandler(
			func(ctx *common.LSPContext, params *SetTraceParams) error {
				callChan <- params
				return nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	setTraceParams := SetTraceParams{
		Value: TraceValueVerbose,
	}
	err = clientLSPContext.Notify(MethodSetTrace, setTraceParams)
	s.Require().NoError(err)

	select {
	case <-ctx.Done():
		s.Fail("timeout")
	case receivedParams := <-callChan:
		s.Require().Equal(setTraceParams, *receivedParams)
	}
}

func (s *HandlerTestSuite) Test_calls_shutdown_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	callChan := make(chan bool, 1)
	serverHandler := NewHandler(
		WithShutdownHandler(
			func(ctx *common.LSPContext) error {
				callChan <- true
				return nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	err = clientLSPContext.Call(MethodShutdown, nil, nil)
	s.Require().NoError(err)

	select {
	case <-ctx.Done():
		s.Fail("timeout")
	case receivedCall := <-callChan:
		s.Require().True(receivedCall)
	}

	// Assert that the server is no longer initialised.
	s.Require().False(serverHandler.IsInitialized())
}

func (s *HandlerTestSuite) Test_calls_exit_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	callChan := make(chan bool, 1)
	serverHandler := NewHandler(
		WithExitHandler(
			func(ctx *common.LSPContext) error {
				callChan <- true
				return nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	err = clientLSPContext.Notify(MethodExit, nil)
	s.Require().NoError(err)

	select {
	case <-ctx.Done():
		s.Fail("timeout")
	case receivedCall := <-callChan:
		s.Require().True(receivedCall)
	}
}

func (s *HandlerTestSuite) Test_calls_text_document_did_open_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	callChan := make(chan *DidOpenTextDocumentParams, 1)
	serverHandler := NewHandler(
		WithTextDocumentDidOpenHandler(
			func(ctx *common.LSPContext, params *DidOpenTextDocumentParams) error {
				callChan <- params
				return nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	textDocumentDidOpenParams := DidOpenTextDocumentParams{
		TextDocument: TextDocumentItem{
			URI:        "file:///test.txt",
			LanguageID: "plaintext",
			Text:       "test text file contents",
			Version:    1,
		},
	}

	err = clientLSPContext.Notify(MethodTextDocumentDidOpen, textDocumentDidOpenParams)
	s.Require().NoError(err)

	select {
	case <-ctx.Done():
		s.Fail("timeout")
	case receivedParams := <-callChan:
		s.Require().Equal(textDocumentDidOpenParams, *receivedParams)
	}
}

func (s *HandlerTestSuite) Test_calls_text_document_did_change_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	callChan := make(chan *DidChangeTextDocumentParams, 1)
	serverHandler := NewHandler(
		WithTextDocumentDidChangeHandler(
			func(ctx *common.LSPContext, params *DidChangeTextDocumentParams) error {
				callChan <- params
				return nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	textDocumentDidChangeParams := DidChangeTextDocumentParams{
		TextDocument: VersionedTextDocumentIdentifier{
			TextDocumentIdentifier: TextDocumentIdentifier{
				URI: "file:///test.txt",
			},
			Version: 1,
		},
		ContentChanges: []interface{}{
			TextDocumentContentChangeEventWhole{
				Text: "new content",
			},
		},
	}

	err = clientLSPContext.Notify(MethodTextDocumentDidChange, textDocumentDidChangeParams)
	s.Require().NoError(err)

	select {
	case <-ctx.Done():
		s.Fail("timeout")
	case receivedParams := <-callChan:
		s.Require().Equal(textDocumentDidChangeParams, *receivedParams)
	}
}

func (s *HandlerTestSuite) Test_calls_text_document_will_save_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	callChan := make(chan *WillSaveTextDocumentParams, 1)
	serverHandler := NewHandler(
		WithTextDocumentWillSaveHandler(
			func(ctx *common.LSPContext, params *WillSaveTextDocumentParams) error {
				callChan <- params
				return nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	textDocumentWillSaveParams := WillSaveTextDocumentParams{
		TextDocument: TextDocumentIdentifier{
			URI: "file:///test.txt",
		},
		Reason: TextDocumentSaveReasonManual,
	}

	err = clientLSPContext.Notify(MethodTextDocumentWillSave, textDocumentWillSaveParams)
	s.Require().NoError(err)

	select {
	case <-ctx.Done():
		s.Fail("timeout")
	case receivedParams := <-callChan:
		s.Require().Equal(textDocumentWillSaveParams, *receivedParams)
	}
}

func (s *HandlerTestSuite) Test_calls_text_document_will_save_wait_until_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	textEdits := []TextEdit{
		{
			Range: Range{
				Start: Position{
					Line:      1,
					Character: 5,
				},
				End: Position{
					Line:      1,
					Character: 10,
				},
			},
			NewText: "new text",
		},
	}
	serverHandler := NewHandler(
		WithTextDocumentWillSaveWaitUntilHandler(
			func(ctx *common.LSPContext, params *WillSaveTextDocumentParams) ([]TextEdit, error) {
				return textEdits, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	textDocumentWillSaveParams := WillSaveTextDocumentParams{
		TextDocument: TextDocumentIdentifier{
			URI: "file:///test.txt",
		},
		Reason: TextDocumentSaveReasonManual,
	}

	returnedResultTextEdits := []TextEdit{}
	err = clientLSPContext.Call(MethodTextDocumentWillSaveWaitUntil, textDocumentWillSaveParams, &returnedResultTextEdits)
	s.Require().NoError(err)
	s.Require().Equal(textEdits, returnedResultTextEdits)
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}
