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
