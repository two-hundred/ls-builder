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
