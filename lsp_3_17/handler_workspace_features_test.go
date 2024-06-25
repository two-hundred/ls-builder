package lsp

import (
	"context"

	"github.com/two-hundred/ls-builder/common"
	"github.com/two-hundred/ls-builder/server"
	"go.uber.org/zap"
)

func (s *HandlerTestSuite) Test_calls_workspace_symbol_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	symbols := []WorkspaceSymbol{
		{
			Name: "testArr",
			Kind: SymbolKindArray,
			Tags: []SymbolTag{SymbolTagDeprecated},
			Location: Location{
				URI: "file:///test_doc.go",
				Range: &Range{
					Start: Position{
						Line:      1,
						Character: 1,
					},
					End: Position{
						Line:      1,
						Character: 5,
					},
				},
			},
		},
	}
	serverHandler := NewHandler(
		WithWorkspaceSymbolHandler(
			func(ctx *common.LSPContext, params *WorkspaceSymbolParams) (any, error) {
				return symbols, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	workspaceSymbolParams := &WorkspaceSymbolParams{
		Query: "fuzzyMatch(input)",
	}

	returnedWorkspaceSymbols := []WorkspaceSymbol{}
	err = clientLSPContext.Call(
		MethodWorkspaceSymbol,
		workspaceSymbolParams,
		&returnedWorkspaceSymbols,
	)
	s.Require().NoError(err)
	s.Require().Equal(symbols, returnedWorkspaceSymbols)
}

func (s *HandlerTestSuite) Test_calls_workspace_symbol_resolve_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	symbol := WorkspaceSymbol{
		Name: "testArr",
		Kind: SymbolKindArray,
		Tags: []SymbolTag{SymbolTagDeprecated},
		Location: Location{
			URI: "file:///test_doc.go",
			Range: &Range{
				Start: Position{
					Line:      1,
					Character: 1,
				},
				End: Position{
					Line:      1,
					Character: 5,
				},
			},
		},
	}
	serverHandler := NewHandler(
		WithWorkspaceSymbolResolveHandler(
			func(ctx *common.LSPContext, params *WorkspaceSymbol) (*WorkspaceSymbol, error) {
				return &symbol, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	workspaceSymbolResolveParams := &WorkspaceSymbol{
		Name: "testBool",
		Kind: SymbolKindBoolean,
		Tags: []SymbolTag{SymbolTagDeprecated},
		Location: Location{
			URI: "file:///test_doc_resolve.go",
			Range: &Range{
				Start: Position{
					Line:      241,
					Character: 1,
				},
				End: Position{
					Line:      241,
					Character: 5,
				},
			},
		},
	}

	returnedSymbol := WorkspaceSymbol{}
	err = clientLSPContext.Call(
		MethodWorkspaceSymbolResolve,
		workspaceSymbolResolveParams,
		&returnedSymbol,
	)
	s.Require().NoError(err)
	s.Require().Equal(symbol, returnedSymbol)
}

func (s *HandlerTestSuite) Test_calls_workspace_did_change_configuration_notification_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	callChan := make(chan *DidChangeConfigurationParams, 1)
	serverHandler := NewHandler(
		WithWorkspaceDidChangeConfigurationHandler(
			func(ctx *common.LSPContext, params *DidChangeConfigurationParams) error {
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
	didChangeConfigurationParams := DidChangeConfigurationParams{
		Settings: map[string]interface{}{
			"test": "test",
		},
	}
	err = clientLSPContext.Notify(MethodWorkspaceDidChangeConfiguration, didChangeConfigurationParams)
	s.Require().NoError(err)

	select {
	case <-ctx.Done():
		s.Fail("timeout")
	case receivedParams := <-callChan:
		s.Require().Equal(didChangeConfigurationParams, *receivedParams)
	}
}

func (s *HandlerTestSuite) Test_calls_workspace_did_change_folders_notification_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	callChan := make(chan *DidChangeWorkspaceFoldersParams, 1)
	serverHandler := NewHandler(
		WithWorkspaceDidChangeFoldersHandler(
			func(ctx *common.LSPContext, params *DidChangeWorkspaceFoldersParams) error {
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
	didChangeFoldersParams := DidChangeWorkspaceFoldersParams{
		Event: WorkspaceFoldersChangeEvent{
			Added: []WorkspaceFolder{
				{
					URI:  "file:///test_folder_2",
					Name: "test_folder_2",
				},
			},
			Removed: []WorkspaceFolder{
				{
					URI:  "file:///test_folder_1",
					Name: "test_folder_1",
				},
			},
		},
	}
	err = clientLSPContext.Notify(MethodWorkspaceDidChangeFolders, didChangeFoldersParams)
	s.Require().NoError(err)

	select {
	case <-ctx.Done():
		s.Fail("timeout")
	case receivedParams := <-callChan:
		s.Require().Equal(didChangeFoldersParams, *receivedParams)
	}
}

func (s *HandlerTestSuite) Test_calls_workspace_will_create_files_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	workspaceEdit := WorkspaceEdit{
		Changes: map[string][]TextEdit{
			"edit1": {
				{
					Range: &Range{
						Start: Position{
							Line:      1,
							Character: 1,
						},
						End: Position{
							Line:      1,
							Character: 5,
						},
					},
					NewText: "help",
				},
			},
		},
	}
	serverHandler := NewHandler(
		WithWorkspaceWillCreateFilesHandler(
			func(ctx *common.LSPContext, params *CreateFilesParams) (*WorkspaceEdit, error) {
				return &workspaceEdit, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	createFilesParams := &CreateFilesParams{
		Files: []FileCreate{
			{
				URI: "file:///test_doc.go",
			},
		},
	}

	returnedEdit := WorkspaceEdit{}
	err = clientLSPContext.Call(
		MethodWorkspaceWillCreateFiles,
		createFilesParams,
		&returnedEdit,
	)
	s.Require().NoError(err)
	s.Require().Equal(workspaceEdit, returnedEdit)
}

func (s *HandlerTestSuite) Test_calls_workspace_did_create_files_notification_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	callChan := make(chan *CreateFilesParams, 1)
	serverHandler := NewHandler(
		WithWorkspaceDidCreateFilesHandler(
			func(ctx *common.LSPContext, params *CreateFilesParams) error {
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
	createFileParams := CreateFilesParams{
		Files: []FileCreate{
			{
				URI: "file:///test_doc.go",
			},
		},
	}
	err = clientLSPContext.Notify(MethodWorkspaceDidCreateFiles, createFileParams)
	s.Require().NoError(err)

	select {
	case <-ctx.Done():
		s.Fail("timeout")
	case receivedParams := <-callChan:
		s.Require().Equal(createFileParams, *receivedParams)
	}
}

func (s *HandlerTestSuite) Test_calls_workspace_will_rename_files_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	workspaceEdit := WorkspaceEdit{
		Changes: map[string][]TextEdit{
			"edit2": {
				{
					Range: &Range{
						Start: Position{
							Line:      1,
							Character: 1,
						},
						End: Position{
							Line:      1,
							Character: 7,
						},
					},
					NewText: "helper",
				},
			},
		},
	}
	serverHandler := NewHandler(
		WithWorkspaceWillRenameFilesHandler(
			func(ctx *common.LSPContext, params *RenameFilesParams) (*WorkspaceEdit, error) {
				return &workspaceEdit, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	renameFilesParams := &RenameFilesParams{
		Files: []FileRename{
			{
				OldURI: "file:///test_doc.go",
				NewURI: "file:///test_doc_renamed.go",
			},
		},
	}

	returnedEdit := WorkspaceEdit{}
	err = clientLSPContext.Call(
		MethodWorkspaceWillRenameFiles,
		renameFilesParams,
		&returnedEdit,
	)
	s.Require().NoError(err)
	s.Require().Equal(workspaceEdit, returnedEdit)
}

func (s *HandlerTestSuite) Test_calls_workspace_did_rename_files_notification_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	callChan := make(chan *RenameFilesParams, 1)
	serverHandler := NewHandler(
		WithWorkspaceDidRenameFilesHandler(
			func(ctx *common.LSPContext, params *RenameFilesParams) error {
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
	renameFilesParams := RenameFilesParams{
		Files: []FileRename{
			{
				OldURI: "file:///test_doc_1.go",
				NewURI: "file:///test_doc_renamed_1.go",
			},
		},
	}
	err = clientLSPContext.Notify(MethodWorkspaceDidRenameFiles, renameFilesParams)
	s.Require().NoError(err)

	select {
	case <-ctx.Done():
		s.Fail("timeout")
	case receivedParams := <-callChan:
		s.Require().Equal(renameFilesParams, *receivedParams)
	}
}

func (s *HandlerTestSuite) Test_calls_workspace_will_delete_files_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	needsConfirmation := true
	ChangeAnnotationDescription := "Change Annotation Description"
	workspaceEdit := WorkspaceEdit{
		ChangeAnnotations: map[string]ChangeAnnotation{
			"changeAnnotation1": {
				Label:             "Change Annotation 1",
				NeedsConfirmation: &needsConfirmation,
				Description:       &ChangeAnnotationDescription,
			},
		},
	}
	serverHandler := NewHandler(
		WithWorkspaceWillDeleteFilesHandler(
			func(ctx *common.LSPContext, params *DeleteFilesParams) (*WorkspaceEdit, error) {
				return &workspaceEdit, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	deleteFilesParams := &DeleteFilesParams{
		Files: []FileDelete{
			{
				URI: "file:///test_doc_to_delete.go",
			},
		},
	}

	returnedEdit := WorkspaceEdit{}
	err = clientLSPContext.Call(
		MethodWorkspaceWillDeleteFiles,
		deleteFilesParams,
		&returnedEdit,
	)
	s.Require().NoError(err)
	s.Require().Equal(workspaceEdit, returnedEdit)
}

func (s *HandlerTestSuite) Test_calls_workspace_did_delete_files_notification_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	callChan := make(chan *DeleteFilesParams, 1)
	serverHandler := NewHandler(
		WithWorkspaceDidDeleteFilesHandler(
			func(ctx *common.LSPContext, params *DeleteFilesParams) error {
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
	deleteFilesParams := DeleteFilesParams{
		Files: []FileDelete{
			{
				URI: "file:///test_doc_deleted.go",
			},
		},
	}
	err = clientLSPContext.Notify(MethodWorkspaceDidDeleteFiles, deleteFilesParams)
	s.Require().NoError(err)

	select {
	case <-ctx.Done():
		s.Fail("timeout")
	case receivedParams := <-callChan:
		s.Require().Equal(deleteFilesParams, *receivedParams)
	}
}

func (s *HandlerTestSuite) Test_calls_workspace_did_change_watched_files_notification_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	callChan := make(chan *DidChangeWatchedFilesParams, 1)
	serverHandler := NewHandler(
		WithWorkspaceDidChangeWatchedFilesHandler(
			func(ctx *common.LSPContext, params *DidChangeWatchedFilesParams) error {
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
	didChangeWatchedFilesParams := DidChangeWatchedFilesParams{
		Changes: []FileEvent{
			{
				URI:  "file:///test_doc_deleted.go",
				Type: FileChangeDeleted,
			},
			{
				URI:  "file:///test_doc_created.go",
				Type: FileChangeCreated,
			},
		},
	}
	err = clientLSPContext.Notify(MethodWorkspaceDidChangeWatchedFiles, didChangeWatchedFilesParams)
	s.Require().NoError(err)

	select {
	case <-ctx.Done():
		s.Fail("timeout")
	case receivedParams := <-callChan:
		s.Require().Equal(didChangeWatchedFilesParams, *receivedParams)
	}
}

func (s *HandlerTestSuite) Test_calls_workspace_execute_command_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	commandExpectedResponse := map[string]interface{}{
		"TestResponseID": "TestResponseValue",
	}
	serverHandler := NewHandler(
		WithWorkspaceExecuteCommandHandler(
			func(ctx *common.LSPContext, params *ExecuteCommandParams) (LSPAny, error) {
				return &commandExpectedResponse, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	commandParams := &ExecuteCommandParams{
		Command: "testCommand",
		Arguments: []interface{}{
			"arg1",
		},
	}

	returnedResponse := map[string]interface{}{}
	err = clientLSPContext.Call(
		MethodWorkspaceExecuteCommand,
		commandParams,
		&returnedResponse,
	)
	s.Require().NoError(err)
	s.Require().Equal(commandExpectedResponse, returnedResponse)
}
