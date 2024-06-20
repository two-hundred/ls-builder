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

func (s *HandlerTestSuite) Test_calls_cancel_request_notification_handler() {
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

func (s *HandlerTestSuite) Test_calls_progress_notification_handler() {
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

func (s *HandlerTestSuite) Test_calls_initialized_notification_handler() {
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

func (s *HandlerTestSuite) Test_calls_set_trace_notification_handler() {
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

func (s *HandlerTestSuite) Test_calls_exit_notification_handler() {
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

func (s *HandlerTestSuite) Test_calls_text_document_did_open_notification_handler() {
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

func (s *HandlerTestSuite) Test_calls_text_document_did_change_notification_handler() {
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

func (s *HandlerTestSuite) Test_calls_text_document_will_save_notification_handler() {
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

func (s *HandlerTestSuite) Test_calls_text_document_did_save_notification_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	callChan := make(chan *DidSaveTextDocumentParams, 1)
	serverHandler := NewHandler(
		WithTextDocumentDidSaveHandler(
			func(ctx *common.LSPContext, params *DidSaveTextDocumentParams) error {
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

	content := "test text file contents"
	textDocumentDidSaveParams := DidSaveTextDocumentParams{
		TextDocument: TextDocumentIdentifier{
			URI: "file:///test.txt",
		},
		Text: &content,
	}

	err = clientLSPContext.Notify(MethodTextDocumentDidSave, textDocumentDidSaveParams)
	s.Require().NoError(err)

	select {
	case <-ctx.Done():
		s.Fail("timeout")
	case receivedParams := <-callChan:
		s.Require().Equal(textDocumentDidSaveParams, *receivedParams)
	}
}

func (s *HandlerTestSuite) Test_calls_text_document_did_close_notification_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	callChan := make(chan *DidCloseTextDocumentParams, 1)
	serverHandler := NewHandler(
		WithTextDocumentDidCloseHandler(
			func(ctx *common.LSPContext, params *DidCloseTextDocumentParams) error {
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

	textDocumentDidCloseParams := DidCloseTextDocumentParams{
		TextDocument: TextDocumentIdentifier{
			URI: "file:///test.txt",
		},
	}

	err = clientLSPContext.Notify(MethodTextDocumentDidClose, textDocumentDidCloseParams)
	s.Require().NoError(err)

	select {
	case <-ctx.Done():
		s.Fail("timeout")
	case receivedParams := <-callChan:
		s.Require().Equal(textDocumentDidCloseParams, *receivedParams)
	}
}

func (s *HandlerTestSuite) Test_calls_notebook_document_did_open_notification_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	callChan := make(chan *DidOpenNotebookDocumentParams, 1)
	serverHandler := NewHandler(
		WithNotebookDocumentDidOpenHandler(
			func(ctx *common.LSPContext, params *DidOpenNotebookDocumentParams) error {
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

	notebookDocumentDidOpenParams := DidOpenNotebookDocumentParams{
		Notebook: NotebookDocument{
			URI: "file:///test.ipynb",
		},
		CellTextDocuments: []TextDocumentItem{
			{
				URI:        "file:///test.ipynb",
				LanguageID: "python",
				Version:    1,
				Text:       "print('hello world')",
			},
		},
	}

	err = clientLSPContext.Notify(MethodNotebookDocumentDidOpen, notebookDocumentDidOpenParams)
	s.Require().NoError(err)

	select {
	case <-ctx.Done():
		s.Fail("timeout")
	case receivedParams := <-callChan:
		s.Require().Equal(notebookDocumentDidOpenParams, *receivedParams)
	}
}

func (s *HandlerTestSuite) Test_calls_notebook_document_did_change_notification_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	callChan := make(chan *DidChangeNotebookDocumentParams, 1)
	serverHandler := NewHandler(
		WithNotebookDocumentDidChangeHandler(
			func(ctx *common.LSPContext, params *DidChangeNotebookDocumentParams) error {
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

	cellExecutionSuccess := true
	notebookDocumentDidChangeParams := DidChangeNotebookDocumentParams{
		NotebookDocument: VersionedNotebookDocumentIdentifier{
			URI:     "file:///test.ipynb",
			Version: 2,
		},
		Change: NotebookDocumentChangeEvent{
			Cells: &NotebookCellChanges{
				Structure: &NotebookCellChangesStructure{
					Array: NotebookCellArrayChange{
						Start:       3,
						DeleteCount: 1,
						Cells: []NotebookCell{
							{
								Kind:     NotebookCellKindMarkup,
								Document: "file:///test.ipynb",
								ExecutionSummary: &NotebookCellExecutionSummary{
									ExecutionOrder: 4,
									Success:        &cellExecutionSuccess,
								},
							},
						},
					},
					DidOpen: []TextDocumentItem{
						{
							URI:        "file:///test.ipynb",
							LanguageID: "python",
							Version:    1,
							Text:       "print('hello world')",
						},
					},
					DidClose: []TextDocumentIdentifier{
						{
							URI: "file:///test.ipynb",
						},
					},
				},
			},
		},
	}

	err = clientLSPContext.Notify(MethodNotebookDocumentDidChange, notebookDocumentDidChangeParams)
	s.Require().NoError(err)

	select {
	case <-ctx.Done():
		s.Fail("timeout")
	case receivedParams := <-callChan:
		s.Require().Equal(notebookDocumentDidChangeParams, *receivedParams)
	}
}

func (s *HandlerTestSuite) Test_calls_notebook_document_did_save_notification_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	callChan := make(chan *DidSaveNotebookDocumentParams, 1)
	serverHandler := NewHandler(
		WithNotebookDocumentDidSaveHandler(
			func(ctx *common.LSPContext, params *DidSaveNotebookDocumentParams) error {
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

	notebookDocumentDidSaveParams := DidSaveNotebookDocumentParams{
		NotebookDocument: NotebookDocumentIdentifier{
			URI: "file:///test.ipynb",
		},
	}

	err = clientLSPContext.Notify(MethodNotebookDocumentDidSave, notebookDocumentDidSaveParams)
	s.Require().NoError(err)

	select {
	case <-ctx.Done():
		s.Fail("timeout")
	case receivedParams := <-callChan:
		s.Require().Equal(notebookDocumentDidSaveParams, *receivedParams)
	}
}

func (s *HandlerTestSuite) Test_calls_notebook_document_did_close_notification_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	callChan := make(chan *DidCloseNotebookDocumentParams, 1)
	serverHandler := NewHandler(
		WithNotebookDocumentDidCloseHandler(
			func(ctx *common.LSPContext, params *DidCloseNotebookDocumentParams) error {
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

	notebookDocumentDidCloseParams := DidCloseNotebookDocumentParams{
		NotebookDocument: NotebookDocumentIdentifier{
			URI: "file:///test.ipynb",
		},
		CellTextDocuments: []TextDocumentIdentifier{
			{
				URI: "file:///test.ipynb",
			},
		},
	}

	err = clientLSPContext.Notify(MethodNotebookDocumentDidClose, notebookDocumentDidCloseParams)
	s.Require().NoError(err)

	select {
	case <-ctx.Done():
		s.Fail("timeout")
	case receivedParams := <-callChan:
		s.Require().Equal(notebookDocumentDidCloseParams, *receivedParams)
	}
}

func (s *HandlerTestSuite) Test_calls_goto_declaration_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	links := []LocationLink{
		{
			TargetURI: "file:///test.go",
			OriginSelectionRange: &Range{
				Start: Position{
					Line:      1,
					Character: 5,
				},
				End: Position{
					Line:      1,
					Character: 10,
				},
			},
			TargetRange: Range{
				Start: Position{
					Line:      3,
					Character: 2,
				},
				End: Position{
					Line:      3,
					Character: 14,
				},
			},
			TargetSelectionRange: Range{
				Start: Position{
					Line:      3,
					Character: 2,
				},
				End: Position{
					Line:      3,
					Character: 14,
				},
			},
		},
	}
	serverHandler := NewHandler(
		WithGotoDeclarationHandler(
			func(ctx *common.LSPContext, params *DeclarationParams) (any, error) {
				return links, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	declarationParams := DeclarationParams{
		TextDocumentPositionParams: TextDocumentPositionParams{
			TextDocument: TextDocumentIdentifier{
				URI: "file:///test.go",
			},
		},
	}

	returnedLinks := []LocationLink{}
	err = clientLSPContext.Call(MethodGotoDeclaration, declarationParams, &returnedLinks)
	s.Require().NoError(err)
	s.Require().Equal(links, returnedLinks)
}

func (s *HandlerTestSuite) Test_calls_goto_definition_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	links := []LocationLink{
		{
			TargetURI: "file:///test_definition.go",
			OriginSelectionRange: &Range{
				Start: Position{
					Line:      10,
					Character: 15,
				},
				End: Position{
					Line:      11,
					Character: 20,
				},
			},
			TargetRange: Range{
				Start: Position{
					Line:      13,
					Character: 12,
				},
				End: Position{
					Line:      13,
					Character: 24,
				},
			},
			TargetSelectionRange: Range{
				Start: Position{
					Line:      13,
					Character: 12,
				},
				End: Position{
					Line:      13,
					Character: 24,
				},
			},
		},
	}
	serverHandler := NewHandler(
		WithGotoDefinitionHandler(
			func(ctx *common.LSPContext, params *DefinitionParams) (any, error) {
				return links, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	definitionParams := DefinitionParams{
		TextDocumentPositionParams: TextDocumentPositionParams{
			TextDocument: TextDocumentIdentifier{
				URI: "file:///test_definition.go",
			},
		},
	}

	returnedLinks := []LocationLink{}
	err = clientLSPContext.Call(MethodGotoDefinition, definitionParams, &returnedLinks)
	s.Require().NoError(err)
	s.Require().Equal(links, returnedLinks)
}

func (s *HandlerTestSuite) Test_calls_goto_type_definition_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	links := []LocationLink{
		{
			TargetURI: "file:///test_type_definition.go",
			OriginSelectionRange: &Range{
				Start: Position{
					Line:      110,
					Character: 115,
				},
				End: Position{
					Line:      111,
					Character: 120,
				},
			},
			TargetRange: Range{
				Start: Position{
					Line:      113,
					Character: 112,
				},
				End: Position{
					Line:      113,
					Character: 124,
				},
			},
			TargetSelectionRange: Range{
				Start: Position{
					Line:      113,
					Character: 112,
				},
				End: Position{
					Line:      113,
					Character: 124,
				},
			},
		},
	}
	serverHandler := NewHandler(
		WithGotoTypeDefinitionHandler(
			func(ctx *common.LSPContext, params *TypeDefinitionParams) (any, error) {
				return links, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	typeDefinitionParams := TypeDefinitionParams{
		TextDocumentPositionParams: TextDocumentPositionParams{
			TextDocument: TextDocumentIdentifier{
				URI: "file:///test_type_definition.go",
			},
		},
	}

	returnedLinks := []LocationLink{}
	err = clientLSPContext.Call(MethodGotoTypeDefinition, typeDefinitionParams, &returnedLinks)
	s.Require().NoError(err)
	s.Require().Equal(links, returnedLinks)
}

func (s *HandlerTestSuite) Test_calls_goto_implementation_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	links := []LocationLink{
		{
			TargetURI: "file:///test_implementation.go",
			OriginSelectionRange: &Range{
				Start: Position{
					Line:      210,
					Character: 215,
				},
				End: Position{
					Line:      211,
					Character: 220,
				},
			},
			TargetRange: Range{
				Start: Position{
					Line:      213,
					Character: 212,
				},
				End: Position{
					Line:      213,
					Character: 224,
				},
			},
			TargetSelectionRange: Range{
				Start: Position{
					Line:      213,
					Character: 212,
				},
				End: Position{
					Line:      213,
					Character: 224,
				},
			},
		},
	}
	serverHandler := NewHandler(
		WithGotoImplementationHandler(
			func(ctx *common.LSPContext, params *ImplementationParams) (any, error) {
				return links, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	implementationParams := ImplementationParams{
		TextDocumentPositionParams: TextDocumentPositionParams{
			TextDocument: TextDocumentIdentifier{
				URI: "file:///test_implementation.go",
			},
		},
	}

	returnedLinks := []LocationLink{}
	err = clientLSPContext.Call(MethodGotoImplementation, implementationParams, &returnedLinks)
	s.Require().NoError(err)
	s.Require().Equal(links, returnedLinks)
}

func (s *HandlerTestSuite) Test_calls_find_references_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	locations := []Location{
		{
			URI: "file:///test_references.go",
			Range: Range{
				Start: Position{
					Line:      310,
					Character: 315,
				},
				End: Position{
					Line:      311,
					Character: 320,
				},
			},
		},
	}
	serverHandler := NewHandler(
		WithFindReferencesHandler(
			func(ctx *common.LSPContext, params *ReferencesParams) ([]Location, error) {
				return locations, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	referencesParams := ReferencesParams{
		TextDocumentPositionParams: TextDocumentPositionParams{
			TextDocument: TextDocumentIdentifier{
				URI: "file:///test_references.go",
			},
		},
		Context: ReferenceContext{
			IncludeDeclaration: true,
		},
	}

	returnedLocations := []Location{}
	err = clientLSPContext.Call(MethodFindReferences, referencesParams, &returnedLocations)
	s.Require().NoError(err)
	s.Require().Equal(locations, returnedLocations)
}

func (s *HandlerTestSuite) Test_calls_prepare_call_hierarchy_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	items := []CallHierarchyItem{
		{
			URI: "file:///test_prepare_call_hierarchy.go",
			Range: Range{
				Start: Position{
					Line:      410,
					Character: 415,
				},
				End: Position{
					Line:      411,
					Character: 420,
				},
			},
		},
	}
	serverHandler := NewHandler(
		WithPrepareCallHierarchyHandler(
			func(ctx *common.LSPContext, params *CallHierarchyPrepareParams) ([]CallHierarchyItem, error) {
				return items, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	prepareCallHierarchyParams := CallHierarchyPrepareParams{
		TextDocumentPositionParams: TextDocumentPositionParams{
			TextDocument: TextDocumentIdentifier{
				URI: "file:///test_prepare_call_hierarchy.go",
			},
		},
	}

	returnedItems := []CallHierarchyItem{}
	err = clientLSPContext.Call(MethodPrepareCallHierarchy, prepareCallHierarchyParams, &returnedItems)
	s.Require().NoError(err)
	s.Require().Equal(items, returnedItems)
}

func (s *HandlerTestSuite) Test_calls_call_hierarchy_incoming_calls_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	calls := []CallHierarchyIncomingCall{
		{
			From: CallHierarchyItem{
				URI: "file:///test_call_hierarchy_incoming_calls.go",
			},
		},
	}
	serverHandler := NewHandler(
		WithCallHierarchyIncomingCallsHandler(
			func(ctx *common.LSPContext, params *CallHierarchyIncomingCallsParams) ([]CallHierarchyIncomingCall, error) {
				return calls, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	callHierarchyIncomingCallsParams := CallHierarchyIncomingCallsParams{
		Item: CallHierarchyItem{
			URI: "file:///test_call_hierarchy_incoming_calls.go",
		},
	}

	returnedCalls := []CallHierarchyIncomingCall{}
	err = clientLSPContext.Call(
		MethodCallHierarchyIncomingCalls,
		callHierarchyIncomingCallsParams,
		&returnedCalls,
	)
	s.Require().NoError(err)
	s.Require().Equal(calls, returnedCalls)
}

func (s *HandlerTestSuite) Test_calls_call_hierarchy_outgoing_calls_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	calls := []CallHierarchyOutgoingCall{
		{
			To: CallHierarchyItem{
				URI: "file:///test_call_hierarchy_outgoing_calls.go",
			},
			FromRanges: []Range{
				{
					Start: Position{
						Line:      1,
						Character: 5,
					},
					End: Position{
						Line:      1,
						Character: 10,
					},
				},
			},
		},
	}
	serverHandler := NewHandler(
		WithCallHierarchyOutgoingCallsHandler(
			func(ctx *common.LSPContext, params *CallHierarchyOutgoingCallsParams) ([]CallHierarchyOutgoingCall, error) {
				return calls, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	callHierarchyOutgoingCallsParams := CallHierarchyOutgoingCallsParams{
		Item: CallHierarchyItem{
			URI: "file:///test_call_hierarchy_outgoing_calls.go",
		},
	}

	returnedCalls := []CallHierarchyOutgoingCall{}
	err = clientLSPContext.Call(
		MethodCallHierarchyOutgoingCalls,
		callHierarchyOutgoingCallsParams,
		&returnedCalls,
	)
	s.Require().NoError(err)
	s.Require().Equal(calls, returnedCalls)
}

func (s *HandlerTestSuite) Test_calls_prepare_type_hierarchy_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	items := []TypeHierarchyItem{
		{
			Name: "TestType",
			Kind: SymbolKindArray,
			URI:  "file:///test_prepare_type_hierarchy.go",
			Range: Range{
				Start: Position{
					Line:      410,
					Character: 415,
				},
				End: Position{
					Line:      411,
					Character: 420,
				},
			},
		},
	}
	serverHandler := NewHandler(
		WithPrepareTypeHierarchyHandler(
			func(ctx *common.LSPContext, params *TypeHierarchyPrepareParams) ([]TypeHierarchyItem, error) {
				return items, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	workDoneToken := "test-token"
	prepareTypeHierarchyParams := TypeHierarchyPrepareParams{
		TextDocumentPositionParams: TextDocumentPositionParams{
			TextDocument: TextDocumentIdentifier{
				URI: "file:///test_prepare_type_hierarchy.go",
			},
			Position: Position{
				Line:      1,
				Character: 5,
			},
		},
		WorkDoneProgressParams: WorkDoneProgressParams{
			WorkDoneToken: &IntOrString{
				StrVal: &workDoneToken,
			},
		},
	}

	returnedItems := []TypeHierarchyItem{}
	err = clientLSPContext.Call(
		MethodPrepareTypeHierarchy,
		prepareTypeHierarchyParams,
		&returnedItems,
	)
	s.Require().NoError(err)
	s.Require().Equal(items, returnedItems)
}

func (s *HandlerTestSuite) Test_calls_type_hierarchy_supertypes_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	items := []TypeHierarchyItem{
		{
			Name: "TestType",
			Kind: SymbolKindArray,
			URI:  "file:///test_type_hierarchy_supertypes.go",
			Range: Range{
				Start: Position{
					Line:      420,
					Character: 140,
				},
				End: Position{
					Line:      420,
					Character: 170,
				},
			},
		},
	}
	serverHandler := NewHandler(
		WithTypeHierarchySupertypesHandler(
			func(ctx *common.LSPContext, params *TypeHierarchySupertypesParams) ([]TypeHierarchyItem, error) {
				return items, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	workDoneToken := "test-token-supertypes"
	typeHierarchySupertypesParams := TypeHierarchySupertypesParams{
		WorkDoneProgressParams: WorkDoneProgressParams{
			WorkDoneToken: &IntOrString{
				StrVal: &workDoneToken,
			},
		},
	}

	returnedItems := []TypeHierarchyItem{}
	err = clientLSPContext.Call(
		MethodTypeHierarchySupertypes,
		typeHierarchySupertypesParams,
		&returnedItems,
	)
	s.Require().NoError(err)
	s.Require().Equal(items, returnedItems)
}

func (s *HandlerTestSuite) Test_calls_type_hierarchy_subtypes_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	items := []TypeHierarchyItem{
		{
			Name: "TestType",
			Kind: SymbolKindArray,
			URI:  "file:///test_type_hierarchy_subtypes.go",
			Range: Range{
				Start: Position{
					Line:      320,
					Character: 140,
				},
				End: Position{
					Line:      320,
					Character: 170,
				},
			},
		},
	}
	serverHandler := NewHandler(
		WithTypeHierarchySubtypesHandler(
			func(ctx *common.LSPContext, params *TypeHierarchySubtypesParams) ([]TypeHierarchyItem, error) {
				return items, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	workDoneToken := "test-token-subtypes"
	typeHierarchySubtypesParams := TypeHierarchySubtypesParams{
		WorkDoneProgressParams: WorkDoneProgressParams{
			WorkDoneToken: &IntOrString{
				StrVal: &workDoneToken,
			},
		},
	}

	returnedItems := []TypeHierarchyItem{}
	err = clientLSPContext.Call(
		MethodTypeHierarchySubtypes,
		typeHierarchySubtypesParams,
		&returnedItems,
	)
	s.Require().NoError(err)
	s.Require().Equal(items, returnedItems)
}

func (s *HandlerTestSuite) Test_calls_document_highlight_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	highlights := []DocumentHighlight{
		{
			Range: Range{
				Start: Position{
					Line:      10,
					Character: 5,
				},
				End: Position{
					Line:      10,
					Character: 15,
				},
			},
			Kind: &DocumentHighlightKindText,
		},
	}
	serverHandler := NewHandler(
		WithDocumentHighlightHandler(
			func(ctx *common.LSPContext, params *DocumentHighlightParams) ([]DocumentHighlight, error) {
				return highlights, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	workDoneToken := "test-token-document-highlight"
	documentHighlightParams := DocumentHighlightParams{
		TextDocumentPositionParams: TextDocumentPositionParams{
			TextDocument: TextDocumentIdentifier{
				URI: "file:///test_document_highlight.go",
			},
			Position: Position{
				Line:      1,
				Character: 5,
			},
		},
		WorkDoneProgressParams: WorkDoneProgressParams{
			WorkDoneToken: &IntOrString{
				StrVal: &workDoneToken,
			},
		},
	}

	returnedHighlights := []DocumentHighlight{}
	err = clientLSPContext.Call(
		MethodDocumentHighlight,
		documentHighlightParams,
		&returnedHighlights,
	)
	s.Require().NoError(err)
	s.Require().Equal(highlights, returnedHighlights)
}

func (s *HandlerTestSuite) Test_calls_document_link_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	target := "file:///test_document_link.go"
	links := []DocumentLink{
		{
			Range: Range{
				Start: Position{
					Line:      10,
					Character: 5,
				},
				End: Position{
					Line:      10,
					Character: 15,
				},
			},
			Target: &target,
		},
	}
	serverHandler := NewHandler(
		WithDocumentLinkHandler(
			func(ctx *common.LSPContext, params *DocumentLinkParams) ([]DocumentLink, error) {
				return links, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	workDoneToken := "test-token-document-link"
	documentLinkParams := DocumentLinkParams{
		WorkDoneProgressParams: WorkDoneProgressParams{
			WorkDoneToken: &IntOrString{
				StrVal: &workDoneToken,
			},
		},
	}

	returnedLinks := []DocumentLink{}
	err = clientLSPContext.Call(
		MethodDocumentLink,
		documentLinkParams,
		&returnedLinks,
	)
	s.Require().NoError(err)
	s.Require().Equal(links, returnedLinks)
}

func (s *HandlerTestSuite) Test_calls_document_link_resolve_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	target := "file:///test_document_link_resolve.go"
	resolvedLink := DocumentLink{
		Range: Range{
			Start: Position{
				Line:      10,
				Character: 5,
			},
			End: Position{
				Line:      10,
				Character: 15,
			},
		},
		Target: &target,
	}
	serverHandler := NewHandler(
		WithDocumentLinkResolveHandler(
			func(ctx *common.LSPContext, params *DocumentLink) (*DocumentLink, error) {
				return &resolvedLink, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	documentLinkResolveParams := DocumentLink{
		Range: Range{
			Start: Position{
				Line:      15,
				Character: 5,
			},
			End: Position{
				Line:      15,
				Character: 22,
			},
		},
		Target: &target,
	}

	returnedLink := DocumentLink{}
	err = clientLSPContext.Call(
		MethodDocumentLinkResolve,
		documentLinkResolveParams,
		&returnedLink,
	)
	s.Require().NoError(err)
	s.Require().Equal(resolvedLink, returnedLink)
}

func (s *HandlerTestSuite) Test_calls_hover_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	hover := Hover{
		Contents: MarkedString{
			Value: MarkedStringLanguage{
				Language: "go",
				Value:    "package main",
			},
		},
	}
	serverHandler := NewHandler(
		WithHoverHandler(
			func(ctx *common.LSPContext, params *HoverParams) (*Hover, error) {
				return &hover, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	workDoneToken := "test-token-hover"
	hoverParams := HoverParams{
		TextDocumentPositionParams: TextDocumentPositionParams{
			TextDocument: TextDocumentIdentifier{
				URI: "file:///test_hover.go",
			},
		},
		WorkDoneProgressParams: WorkDoneProgressParams{
			WorkDoneToken: &IntOrString{StrVal: &workDoneToken},
		},
	}

	returnedHover := Hover{}
	err = clientLSPContext.Call(
		MethodHover,
		hoverParams,
		&returnedHover,
	)
	s.Require().NoError(err)
	s.Require().Equal(hover, returnedHover)
}

func (s *HandlerTestSuite) Test_calls_code_lens_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	codeLensResults := []CodeLens{
		{
			Range: Range{
				Start: Position{
					Line:      10,
					Character: 5,
				},
				End: Position{
					Line:      10,
					Character: 15,
				},
			},
			Command: &Command{
				Title:     "Run test",
				Command:   "test.run",
				Arguments: []interface{}{"test.go"},
			},
		},
	}
	serverHandler := NewHandler(
		WithCodeLensHandler(
			func(ctx *common.LSPContext, params *CodeLensParams) ([]CodeLens, error) {
				return codeLensResults, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	workDoneToken := "test-token-code-lens"
	codeLensParams := CodeLensParams{
		WorkDoneProgressParams: WorkDoneProgressParams{
			WorkDoneToken: &IntOrString{StrVal: &workDoneToken},
		},
	}

	returnedCodeLensResults := []CodeLens{}
	err = clientLSPContext.Call(
		MethodCodeLens,
		codeLensParams,
		&returnedCodeLensResults,
	)
	s.Require().NoError(err)
	s.Require().Equal(codeLensResults, returnedCodeLensResults)
}

func (s *HandlerTestSuite) Test_calls_code_lens_resolve_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	codeLensResult := CodeLens{
		Range: Range{
			Start: Position{
				Line:      10,
				Character: 5,
			},
			End: Position{
				Line:      10,
				Character: 15,
			},
		},
		Command: &Command{
			Title:     "Run test",
			Command:   "test.run",
			Arguments: []interface{}{"test.go"},
		},
	}
	serverHandler := NewHandler(
		WithCodeLensResolveHandler(
			func(ctx *common.LSPContext, params *CodeLens) (*CodeLens, error) {
				return &codeLensResult, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	codeLensParams := CodeLens{
		Range: Range{
			Start: Position{
				Line:      100,
				Character: 11,
			},
			End: Position{
				Line:      100,
				Character: 25,
			},
		},
		Command: &Command{
			Title:     "Run test",
			Command:   "test.run",
			Arguments: []interface{}{"test.go"},
		},
	}

	returnedCodeLensResult := CodeLens{}
	err = clientLSPContext.Call(
		MethodCodeLensResolve,
		codeLensParams,
		&returnedCodeLensResult,
	)
	s.Require().NoError(err)
	s.Require().Equal(codeLensResult, returnedCodeLensResult)
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}
