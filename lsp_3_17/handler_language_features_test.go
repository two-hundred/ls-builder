package lsp

import (
	"context"

	"github.com/two-hundred/ls-builder/common"
	"github.com/two-hundred/ls-builder/server"
	"go.uber.org/zap"
)

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
			Range: &Range{
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

func (s *HandlerTestSuite) Test_calls_folding_range_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	startChar := UInteger(5)
	endChar := UInteger(10)
	foldingRanges := []FoldingRange{
		{
			StartLine:      10,
			StartCharacter: &startChar,
			EndLine:        15,
			EndCharacter:   &endChar,
			Kind:           &FoldingRangeKindRegion,
		},
	}
	serverHandler := NewHandler(
		WithFoldingRangeHandler(
			func(ctx *common.LSPContext, params *FoldingRangeParams) ([]FoldingRange, error) {
				return foldingRanges, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	foldingRangeParams := FoldingRangeParams{
		TextDocument: TextDocumentIdentifier{
			URI: "file:///test_folding_range.go",
		},
	}

	returnedFoldingRanges := []FoldingRange{}
	err = clientLSPContext.Call(
		MethodFoldingRange,
		foldingRangeParams,
		&returnedFoldingRanges,
	)
	s.Require().NoError(err)
	s.Require().Equal(foldingRanges, returnedFoldingRanges)
}

func (s *HandlerTestSuite) Test_calls_selection_range_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	selectionRanges := []SelectionRange{
		{
			Range: Range{
				Start: Position{
					Line:      10,
					Character: 5,
				},
				End: Position{
					Line:      10,
					Character: 35,
				},
			},
			Parent: &SelectionRange{
				Range: Range{
					Start: Position{
						Line:      5,
						Character: 2,
					},
					End: Position{
						Line:      5,
						Character: 45,
					},
				},
			},
		},
	}
	serverHandler := NewHandler(
		WithSelectionRangeHandler(
			func(ctx *common.LSPContext, params *SelectionRangeParams) ([]SelectionRange, error) {
				return selectionRanges, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	selectionRangeParams := SelectionRangeParams{
		TextDocument: TextDocumentIdentifier{
			URI: "file:///test_selection_range.go",
		},
		Positions: []Position{
			{
				Line:      10,
				Character: 5,
			},
		},
	}

	returnedSelectionRanges := []SelectionRange{}
	err = clientLSPContext.Call(
		MethodSelectionRange,
		selectionRangeParams,
		&returnedSelectionRanges,
	)
	s.Require().NoError(err)
	s.Require().Equal(selectionRanges, returnedSelectionRanges)
}

func (s *HandlerTestSuite) Test_calls_document_symbol_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	symbols := []DocumentSymbol{
		{
			Name: "TestSymbol",
			Kind: SymbolKindArray,
			Range: Range{
				Start: Position{
					Line:      10,
					Character: 5,
				},
			},
			SelectionRange: Range{
				Start: Position{
					Line:      10,
					Character: 5,
				},
			},
		},
	}
	serverHandler := NewHandler(
		WithDocumentSymbolHandler(
			func(ctx *common.LSPContext, params *DocumentSymbolParams) (any, error) {
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

	documentSymbolParams := DocumentSymbolParams{
		TextDocument: TextDocumentIdentifier{
			URI: "file:///test_selection_range.go",
		},
	}

	returnedDocumentSymbols := []DocumentSymbol{}
	err = clientLSPContext.Call(
		MethodDocumentSymbol,
		documentSymbolParams,
		&returnedDocumentSymbols,
	)
	s.Require().NoError(err)
	s.Require().Equal(symbols, returnedDocumentSymbols)
}

func (s *HandlerTestSuite) Test_calls_semantic_tokens_full_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	resultID := "test-result-id"
	semanticTokens := SemanticTokens{
		ResultID: &resultID,
		Data:     []UInteger{0, 5, 10, 13, 3, 20},
	}
	serverHandler := NewHandler(
		WithSemanticTokensFullHandler(
			func(ctx *common.LSPContext, params *SemanticTokensParams) (*SemanticTokens, error) {
				return &semanticTokens, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	semanticTokensParams := SemanticTokensParams{
		TextDocument: TextDocumentIdentifier{
			URI: "file:///test_semantic_tokens_full.go",
		},
	}

	returnedSemanticTokens := SemanticTokens{}
	err = clientLSPContext.Call(
		MethodSemanticTokensFull,
		semanticTokensParams,
		&returnedSemanticTokens,
	)
	s.Require().NoError(err)
	s.Require().Equal(semanticTokens, returnedSemanticTokens)
}

func (s *HandlerTestSuite) Test_calls_semantic_tokens_full_delta_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	resultID := "test-result-id-delta"
	semanticTokensDelta := SemanticTokensDelta{
		ResultID: &resultID,
		Edits: []SemanticTokensEdit{
			{
				Start:       0,
				DeleteCount: 1,
				Data:        []UInteger{0, 5, 10, 13, 3, 20},
			},
		},
	}
	serverHandler := NewHandler(
		WithSemanticTokensFullDeltaHandler(
			func(ctx *common.LSPContext, params *SemanticTokensDeltaParams) (*SemanticTokensDelta, error) {
				return &semanticTokensDelta, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	semanticTokensDeltaParams := SemanticTokensDeltaParams{
		TextDocument: TextDocumentIdentifier{
			URI: "file:///test_semantic_tokens_full_delta.go",
		},
	}

	returnedSemanticDeltaTokens := SemanticTokensDelta{}
	err = clientLSPContext.Call(
		MethodSemanticTokensFullDelta,
		semanticTokensDeltaParams,
		&returnedSemanticDeltaTokens,
	)
	s.Require().NoError(err)
	s.Require().Equal(semanticTokensDelta, returnedSemanticDeltaTokens)
}

func (s *HandlerTestSuite) Test_calls_semantic_tokens_range_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	resultID := "test-result-id-range"
	semanticTokens := SemanticTokens{
		ResultID: &resultID,
		Data:     []UInteger{0, 5, 10, 13, 3, 20},
	}
	serverHandler := NewHandler(
		WithSemanticTokensRangeHandler(
			func(ctx *common.LSPContext, params *SemanticTokensRangeParams) (*SemanticTokens, error) {
				return &semanticTokens, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	semanticTokensRangeParams := SemanticTokensRangeParams{
		TextDocument: TextDocumentIdentifier{
			URI: "file:///test_semantic_tokens_range.go",
		},
	}

	returnedSemanticTokens := SemanticTokens{}
	err = clientLSPContext.Call(
		MethodSemanticTokensRange,
		semanticTokensRangeParams,
		&returnedSemanticTokens,
	)
	s.Require().NoError(err)
	s.Require().Equal(semanticTokens, returnedSemanticTokens)
}

func (s *HandlerTestSuite) Test_calls_inlay_hint_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	inlayHints := []*InlayHint{
		{
			Position: Position{
				Line:      10,
				Character: 5,
			},
			Label: "TestHint",
			Kind:  &InlayHintKindType,
			Tooltip: MarkupContent{
				Kind:  MarkupKindMarkdown,
				Value: "TestHintTooltip",
			},
			TextEdits: []TextEdit{
				{
					Range: &Range{
						Start: Position{
							Line:      10,
							Character: 5,
						},
						End: Position{
							Line:      10,
							Character: 15,
						},
					},
				},
			},
		},
		{
			Position: Position{
				Line:      15,
				Character: 5,
			},
			Label: []*InlayHintLabelPart{
				{
					Value: "TestHint2",
					Tooltip: MarkupContent{
						Kind:  MarkupKindMarkdown,
						Value: "TestHintTooltip2",
					},
				},
			},
			Kind: &InlayHintKindType,
		},
	}
	serverHandler := NewHandler(
		WithInlayHintHandler(
			func(ctx *common.LSPContext, params *InlayHintParams) ([]*InlayHint, error) {
				return inlayHints, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	inlayHintParams := InlayHintParams{
		TextDocument: TextDocumentIdentifier{
			URI: "file:///test_inlay_hint.go",
		},
	}

	returnedInlayHints := []*InlayHint{}
	err = clientLSPContext.Call(
		MethodInlayHint,
		inlayHintParams,
		&returnedInlayHints,
	)
	s.Require().NoError(err)
	s.Require().Equal(inlayHints, returnedInlayHints)
}

func (s *HandlerTestSuite) Test_calls_inlay_hint_resolve_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	inlayHint := &InlayHint{
		Position: Position{
			Line:      10,
			Character: 5,
		},
		Label: "TestHint",
		Kind:  &InlayHintKindType,
		Tooltip: MarkupContent{
			Kind:  MarkupKindMarkdown,
			Value: "TestHintTooltip",
		},
		TextEdits: []TextEdit{
			{
				Range: &Range{
					Start: Position{
						Line:      10,
						Character: 5,
					},
					End: Position{
						Line:      10,
						Character: 15,
					},
				},
			},
		},
	}
	serverHandler := NewHandler(
		WithInlayHintResolveHandler(
			func(ctx *common.LSPContext, params *InlayHint) (*InlayHint, error) {
				return inlayHint, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	inlayHintParams := InlayHint{
		Position: Position{
			Line:      10,
			Character: 5,
		},
		Label: "TestHint",
		Kind:  &InlayHintKindType,
		Tooltip: MarkupContent{
			Kind:  MarkupKindMarkdown,
			Value: "TestHintTooltip",
		},
	}

	returnedInlayHint := &InlayHint{}
	err = clientLSPContext.Call(
		MethodInlayHintResolve,
		inlayHintParams,
		&returnedInlayHint,
	)
	s.Require().NoError(err)
	s.Require().Equal(inlayHint, returnedInlayHint)
}

func (s *HandlerTestSuite) Test_calls_inline_value_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	variableName := "TestVariable"
	expression := "testExpression(someArgument)"
	inlineValues := []*InlineValue{
		{
			InlineValueText: &InlineValueText{
				Range: Range{
					Start: Position{
						Line:      20,
						Character: 5,
					},
					End: Position{
						Line:      20,
						Character: 13,
					},
				},
				Text: "TestInlineValue",
			},
		},
		{
			InlineValueVariableLookup: &InlineValueVariableLookup{
				Range: Range{
					Start: Position{
						Line:      25,
						Character: 5,
					},
					End: Position{
						Line:      25,
						Character: 13,
					},
				},
				VariableName:        &variableName,
				CaseSensitiveLookup: true,
			},
		},
		{
			InlineValueEvaluatable: &InlineValueEvaluatableExpression{
				Range: Range{
					Start: Position{
						Line:      30,
						Character: 5,
					},
					End: Position{
						Line:      30,
						Character: 14,
					},
				},
				Expression: &expression,
			},
		},
	}
	serverHandler := NewHandler(
		WithInlineValueHandler(
			func(ctx *common.LSPContext, params *InlineValueParams) ([]*InlineValue, error) {
				return inlineValues, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	inlineValueParams := InlineValueParams{
		TextDocument: TextDocumentIdentifier{
			URI: "file:///test_inline_value.go",
		},
		Context: InlineValueContext{
			FrameID: 14,
			StoppedLocation: Range{
				Start: Position{
					Line:      10,
					Character: 5,
				},
				End: Position{
					Line:      10,
					Character: 13,
				},
			},
		},
	}

	returnedInlineValues := []*InlineValue{}
	err = clientLSPContext.Call(
		MethodInlineValue,
		inlineValueParams,
		&returnedInlineValues,
	)
	s.Require().NoError(err)
	s.Require().Equal(inlineValues, returnedInlineValues)
}

func (s *HandlerTestSuite) Test_calls_moniker_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	monikers := []Moniker{
		{
			Scheme:     "tsc",
			Identifier: "test-moniker",
			Unique:     UniquenessLevelGlobal,
			Kind:       &MonikerKindImport,
		},
	}
	serverHandler := NewHandler(
		WithMonikerHandler(
			func(ctx *common.LSPContext, params *MonikerParams) ([]Moniker, error) {
				return monikers, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	monikerParams := MonikerParams{
		TextDocumentPositionParams: TextDocumentPositionParams{
			TextDocument: TextDocumentIdentifier{
				URI: "file:///test_moniker.go",
			},
		},
	}

	returnedMonikers := []Moniker{}
	err = clientLSPContext.Call(
		MethodMoniker,
		monikerParams,
		&returnedMonikers,
	)
	s.Require().NoError(err)
	s.Require().Equal(monikers, returnedMonikers)
}

func (s *HandlerTestSuite) Test_calls_completion_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	labelDetail := "The label for the custom completion item unique to the test language server"
	detail := "A custom completion item unique to the test language server"
	insertText := "Insert custom completion item"
	completionItemKind := CompletionItemKindClass
	completionList := CompletionList{
		IsIncomplete: false,
		ItemDefaults: &CompletionItemDefaults{
			CommitCharacters: []string{"a", "b", "c"},
			EditRange: InsertReplaceRange{
				Insert: &Range{
					Start: Position{
						Line:      145,
						Character: 50,
					},
					End: Position{
						Line:      145,
						Character: 87,
					},
				},
			},
			InsertTextFormat: &InsertTextFormatPlainText,
			InsertTextMode:   &InsertTextModeAsIs,
		},
		Items: []*CompletionItem{
			{
				Label: "CompletionItem",
				LabelDetails: &CompletionItemLabelDetails{
					Detail: &labelDetail,
				},
				Kind: &completionItemKind,
				Tags: []CompletionItemTag{
					CompletionItemTagDeprecated,
				},
				Detail: &detail,
				Documentation: MarkupContent{
					Kind:  MarkupKindMarkdown,
					Value: "# Completion Item Info\nSome additional information about the completion item",
				},
				InsertText:       &insertText,
				InsertTextFormat: &InsertTextFormatPlainText,
				InsertTextMode:   &InsertTextModeAsIs,
				TextEdit: InsertReplaceEdit{
					Insert: &Range{
						Start: Position{
							Line:      1403,
							Character: 50,
						},
						End: Position{
							Line:      1760,
							Character: 100,
						},
					},
					Replace: &Range{
						Start: Position{
							Line:      1203,
							Character: 51,
						},
						End: Position{
							Line:      1657,
							Character: 110,
						},
					},
				},
			},
		},
	}
	serverHandler := NewHandler(
		WithCompletionHandler(
			func(ctx *common.LSPContext, params *CompletionParams) (any, error) {
				return completionList, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	completionParams := CompletionParams{
		TextDocumentPositionParams: TextDocumentPositionParams{
			TextDocument: TextDocumentIdentifier{
				URI: "file:///test_completion.go",
			},
		},
	}

	returnedCompletionList := CompletionList{}
	err = clientLSPContext.Call(
		MethodCompletion,
		completionParams,
		&returnedCompletionList,
	)
	s.Require().NoError(err)
	s.Require().Equal(completionList, returnedCompletionList)
}

func (s *HandlerTestSuite) Test_calls_completion_item_resolve_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	labelDetail := "The label for the custom completion item unique to the test language server"
	detail := "A custom completion item unique to the test language server"
	insertText := "Insert custom completion item"
	completionItemKind := CompletionItemKindClass
	completionItem := &CompletionItem{
		Label: "CompletionItem",
		LabelDetails: &CompletionItemLabelDetails{
			Detail: &labelDetail,
		},
		Kind: &completionItemKind,
		Tags: []CompletionItemTag{
			CompletionItemTagDeprecated,
		},
		Detail: &detail,
		Documentation: MarkupContent{
			Kind:  MarkupKindMarkdown,
			Value: "# Completion Item Info\nSome additional information about the completion item",
		},
		InsertText:       &insertText,
		InsertTextFormat: &InsertTextFormatPlainText,
		InsertTextMode:   &InsertTextModeAsIs,
		TextEdit: InsertReplaceEdit{
			Insert: &Range{
				Start: Position{
					Line:      1403,
					Character: 50,
				},
				End: Position{
					Line:      1760,
					Character: 100,
				},
			},
			Replace: &Range{
				Start: Position{
					Line:      1203,
					Character: 51,
				},
				End: Position{
					Line:      1657,
					Character: 110,
				},
			},
		},
	}
	serverHandler := NewHandler(
		WithCompletionItemResolveHandler(
			func(ctx *common.LSPContext, params *CompletionItem) (*CompletionItem, error) {
				return completionItem, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	completionItemResolveParams := &CompletionItem{
		Label: "CompletionItem",
		// Input params without label details which will be resolved by the handler.
		Kind: &completionItemKind,
		Tags: []CompletionItemTag{
			CompletionItemTagDeprecated,
		},
		Detail: &detail,
		// Input is without documentation, this will be resolved by the handler.
		InsertText:       &insertText,
		InsertTextFormat: &InsertTextFormatPlainText,
		InsertTextMode:   &InsertTextModeAsIs,
		TextEdit: InsertReplaceEdit{
			Insert: &Range{
				Start: Position{
					Line:      1403,
					Character: 50,
				},
				End: Position{
					Line:      1760,
					Character: 100,
				},
			},
			Replace: &Range{
				Start: Position{
					Line:      1203,
					Character: 51,
				},
				End: Position{
					Line:      1657,
					Character: 110,
				},
			},
		},
	}

	returnedCompletionItem := &CompletionItem{}
	err = clientLSPContext.Call(
		MethodCompletionItemResolve,
		completionItemResolveParams,
		&returnedCompletionItem,
	)
	s.Require().NoError(err)
	s.Require().Equal(completionItem, returnedCompletionItem)
}

func (s *HandlerTestSuite) Test_calls_signature_help_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	activeSignature := UInteger(10)
	activeParameter := UInteger(2)
	signatureHelp := &SignatureHelp{
		Signatures: []*SignatureInformation{
			{
				Label:         "SignatureTest1",
				Documentation: "SignatureTest1 Documentation",
				Parameters: []*ParameterInformation{
					{
						Label:         "param1",
						Documentation: "param1 Documentation",
					},
					{
						Label: [2]UInteger{56, 893},
						Documentation: MarkupContent{
							Kind:  MarkupKindMarkdown,
							Value: "# param2 Documentation",
						},
					},
				},
				ActiveParameter: &activeParameter,
			},
			{
				Label: "SignatureTest2",
				Documentation: MarkupContent{
					Kind:  MarkupKindMarkdown,
					Value: "# SignatureTest2 Documentation",
				},
				Parameters: []*ParameterInformation{
					{
						Label:         "param1",
						Documentation: "param1 Documentation",
					},
				},
			},
		},
		ActiveSignature: &activeSignature,
		ActiveParameter: &activeParameter,
	}

	serverHandler := NewHandler(
		WithSignatureHelpHandler(
			func(ctx *common.LSPContext, params *SignatureHelpParams) (*SignatureHelp, error) {
				return signatureHelp, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	triggerChar := "("
	signatureHelpParams := SignatureHelpParams{
		Context: &SignatureHelpContext{
			TriggerKind:      SignatureHelpTriggerKindTriggerCharacter,
			TriggerCharacter: &triggerChar,
			IsRetrigger:      false,
		},
	}

	returnedSignatureHelp := &SignatureHelp{}
	err = clientLSPContext.Call(
		MethodSignatureHelp,
		signatureHelpParams,
		returnedSignatureHelp,
	)
	s.Require().NoError(err)
	s.Require().Equal(signatureHelp, returnedSignatureHelp)
}

func (s *HandlerTestSuite) Test_calls_code_action_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	codeAction1 := CodeActionKindQuickFix
	severity1 := DiagnosticSeverityError
	is1Preferred := true
	diagnosticCode := "ErrorCode"
	codeActions := []*CodeActionOrCommand{
		{
			CodeAction: &CodeAction{
				Title: "TestCodeAction",
				Kind:  &codeAction1,
				Diagnostics: []Diagnostic{
					{
						Range: Range{
							Start: Position{
								Line:      205,
								Character: 5,
							},
							End: Position{
								Line:      205,
								Character: 15,
							},
						},
						Severity: &severity1,
						Code: &IntOrString{
							StrVal: &diagnosticCode,
						},
						Message: "Test Diagnostic Message",
					},
				},
				IsPreferred: &is1Preferred,
			},
		},
		{
			Command: &Command{
				Title:   "save",
				Command: "save.command",
				Arguments: []interface{}{
					"file:///test_code_action.go",
				},
			},
		},
	}
	serverHandler := NewHandler(
		WithCodeActionHandler(
			func(ctx *common.LSPContext, params *CodeActionParams) ([]*CodeActionOrCommand, error) {
				return codeActions, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	codeActionParams := CodeActionParams{
		TextDocument: TextDocumentIdentifier{
			URI: "file:///test_code_action.go",
		},
		Range: Range{
			Start: Position{
				Line:      205,
				Character: 5,
			},
			End: Position{
				Line:      205,
				Character: 15,
			},
		},
		Context: CodeActionContext{
			Diagnostics: []Diagnostic{
				{
					Range: Range{
						Start: Position{
							Line:      105,
							Character: 5,
						},
						End: Position{
							Line:      105,
							Character: 23,
						},
					},
					Severity: &severity1,
					Code: &IntOrString{
						StrVal: &diagnosticCode,
					},
					Message: "Test Diagnostic Message",
				},
			},
			Only: []CodeActionKind{
				CodeActionKindQuickFix,
				CodeActionKindRefactor,
			},
			TriggerKind: &CodeActionTriggerKindAutomatic,
		},
	}

	returnedCodeActions := []*CodeActionOrCommand{}
	err = clientLSPContext.Call(
		MethodCodeAction,
		codeActionParams,
		&returnedCodeActions,
	)
	s.Require().NoError(err)
	s.Require().Equal(codeActions, returnedCodeActions)
}

func (s *HandlerTestSuite) Test_calls_code_action_resolve_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	codeAction1 := CodeActionKindQuickFix
	severity1 := DiagnosticSeverityError
	is1Preferred := true
	diagnosticCode := "ErrorCode"
	codeAction := CodeAction{
		Title: "TestCodeAction",
		Kind:  &codeAction1,
		Diagnostics: []Diagnostic{
			{
				Range: Range{
					Start: Position{
						Line:      305,
						Character: 5,
					},
					End: Position{
						Line:      305,
						Character: 115,
					},
				},
				Severity: &severity1,
				Code: &IntOrString{
					StrVal: &diagnosticCode,
				},
				Message: "Test Diagnostic Message Resolve",
			},
		},
		IsPreferred: &is1Preferred,
	}
	serverHandler := NewHandler(
		WithCodeActionResolveHandler(
			func(ctx *common.LSPContext, params *CodeAction) (*CodeAction, error) {
				return &codeAction, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	codeAction2 := CodeActionKindRefactor
	codeActionParams := &CodeAction{
		Title: "TestCodeActionTrigger",
		Kind:  &codeAction2,
		Diagnostics: []Diagnostic{
			{
				Range: Range{
					Start: Position{
						Line:      202,
						Character: 5,
					},
					End: Position{
						Line:      202,
						Character: 115,
					},
				},
				Severity: &severity1,
				Code: &IntOrString{
					StrVal: &diagnosticCode,
				},
				Message: "Test Diagnostic Message Resolve Trigger",
			},
		},
		IsPreferred: &is1Preferred,
	}

	returnedCodeAction := CodeAction{}
	err = clientLSPContext.Call(
		MethodCodeActionResolve,
		codeActionParams,
		&returnedCodeAction,
	)
	s.Require().NoError(err)
	s.Require().Equal(codeAction, returnedCodeAction)
}

func (s *HandlerTestSuite) Test_calls_document_color_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	colorInfo := []ColorInformation{
		{
			Range: Range{
				Start: Position{
					Line:      300,
					Character: 10,
				},
				End: Position{
					Line:      300,
					Character: 55,
				},
			},
			Color: Color{
				Red:   1.0,
				Green: 0.5,
				Blue:  0.56,
				Alpha: 1.0,
			},
		},
	}
	serverHandler := NewHandler(
		WithDocumentColorHandler(
			func(ctx *common.LSPContext, params *DocumentColorParams) ([]ColorInformation, error) {
				return colorInfo, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	documentColorParams := &DocumentColorParams{
		TextDocument: TextDocumentIdentifier{
			URI: "file:///test_doc_color.go",
		},
	}

	returnedColorInfo := []ColorInformation{}
	err = clientLSPContext.Call(
		MethodDocumentColor,
		documentColorParams,
		&returnedColorInfo,
	)
	s.Require().NoError(err)
	s.Require().Equal(colorInfo, returnedColorInfo)
}

func (s *HandlerTestSuite) Test_calls_document_color_presentation_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	colorPresentation := []ColorPresentation{
		{
			Label: "Test Color Presentation Item",
			TextEdit: &TextEdit{
				Range: &Range{
					Start: Position{
						Line:      1002,
						Character: 58,
					},
					End: Position{
						Line:      1002,
						Character: 122,
					},
				},
				NewText: "Some new text",
			},
		},
	}
	serverHandler := NewHandler(
		WithDocumentColorPresentationHandler(
			func(ctx *common.LSPContext, params *ColorPresentationParams) ([]ColorPresentation, error) {
				return colorPresentation, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	colorPresentationParams := &ColorPresentationParams{
		TextDocument: TextDocumentIdentifier{
			URI: "file:///test_doc_color_presentation.go",
		},
	}

	returnedColorPresentation := []ColorPresentation{}
	err = clientLSPContext.Call(
		MethodDocumentColorPresentation,
		colorPresentationParams,
		&returnedColorPresentation,
	)
	s.Require().NoError(err)
	s.Require().Equal(colorPresentation, returnedColorPresentation)
}

func (s *HandlerTestSuite) Test_calls_document_formatting_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	textEdits := []TextEdit{
		{
			Range: &Range{
				Start: Position{
					Line:      239,
					Character: 45,
				},
				End: Position{
					Line:      239,
					Character: 62,
				},
			},
			NewText: "	label, err := container.GetLabel(ctx, id)",
		},
	}
	serverHandler := NewHandler(
		WithDocumentFormattingHandler(
			func(ctx *common.LSPContext, params *DocumentFormattingParams) ([]TextEdit, error) {
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

	docFormattingParams := &DocumentFormattingParams{
		TextDocument: TextDocumentIdentifier{
			URI: "file:///test_doc_formatting.go",
		},
		Options: map[string]any{
			FormattingOptionTabSize: Integer(4),
		},
	}

	returnedTextEdits := []TextEdit{}
	err = clientLSPContext.Call(
		MethodDocumentFormatting,
		docFormattingParams,
		&returnedTextEdits,
	)
	s.Require().NoError(err)
	s.Require().Equal(textEdits, returnedTextEdits)
}

func (s *HandlerTestSuite) Test_calls_document_range_formatting_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	textEdits := []TextEdit{
		{
			Range: &Range{
				Start: Position{
					Line:      139,
					Character: 15,
				},
				End: Position{
					Line:      205,
					Character: 24,
				},
			},
			NewText: "	user, err := service.GetUser(ctx, id)",
		},
	}
	serverHandler := NewHandler(
		WithDocumentRangeFormattingHandler(
			func(ctx *common.LSPContext, params *DocumentRangeFormattingParams) ([]TextEdit, error) {
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

	docRangeFormattingParams := &DocumentRangeFormattingParams{
		TextDocument: TextDocumentIdentifier{
			URI: "file:///test_doc_range_formatting.go",
		},
		Options: map[string]any{
			FormattingOptionTabSize: Integer(2),
		},
	}

	returnedTextEdits := []TextEdit{}
	err = clientLSPContext.Call(
		MethodDocumentRangeFormatting,
		docRangeFormattingParams,
		&returnedTextEdits,
	)
	s.Require().NoError(err)
	s.Require().Equal(textEdits, returnedTextEdits)
}

func (s *HandlerTestSuite) Test_calls_document_on_type_formatting_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	textEdits := []TextEdit{
		{
			Range: &Range{
				Start: Position{
					Line:      439,
					Character: 15,
				},
				End: Position{
					Line:      439,
					Character: 24,
				},
			},
			NewText: "type User struct{\n	ID string `json:\"id\"`\n}",
		},
	}
	serverHandler := NewHandler(
		WithDocumentOnTypeFormattingHandler(
			func(ctx *common.LSPContext, params *DocumentOnTypeFormattingParams) ([]TextEdit, error) {
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

	docOnTypeFormattingParams := &DocumentOnTypeFormattingParams{
		TextDocument: TextDocumentIdentifier{
			URI: "file:///test_doc_on_type_formatting.go",
		},
		Options: map[string]any{
			FormattingOptionTabSize: Integer(8),
		},
	}

	returnedTextEdits := []TextEdit{}
	err = clientLSPContext.Call(
		MethodDocumentOnTypeFormatting,
		docOnTypeFormattingParams,
		&returnedTextEdits,
	)
	s.Require().NoError(err)
	s.Require().Equal(textEdits, returnedTextEdits)
}

func (s *HandlerTestSuite) Test_calls_document_rename_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	workspaceEdit := WorkspaceEdit{}
	serverHandler := NewHandler(
		WithDocumentRenameHandler(
			func(ctx *common.LSPContext, params *RenameParams) (*WorkspaceEdit, error) {
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

	renameParams := &RenameParams{
		NewName: "userv2.go",
	}

	returnedWorkspaceEdit := WorkspaceEdit{}
	err = clientLSPContext.Call(
		MethodDocumentRename,
		renameParams,
		&returnedWorkspaceEdit,
	)
	s.Require().NoError(err)
	s.Require().Equal(workspaceEdit, returnedWorkspaceEdit)
}

func (s *HandlerTestSuite) Test_calls_document_prepare_rename_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	rangeWithPlaceholder := RangeWithPlaceholder{
		Range: Range{
			Start: Position{
				Line:      300,
				Character: 24,
			},
			End: Position{
				Line:      300,
				Character: 115,
			},
		},
		Placeholder: "user.go",
	}
	serverHandler := NewHandler(
		WithDocumentPrepareRenameHandler(
			func(ctx *common.LSPContext, params *PrepareRenameParams) (any, error) {
				return rangeWithPlaceholder, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	prepareRenameParams := &PrepareRenameParams{
		TextDocumentPositionParams: TextDocumentPositionParams{
			TextDocument: TextDocumentIdentifier{
				URI: "file:///test_doc_prepare_rename.go",
			},
			Position: Position{
				Line:      320,
				Character: 24,
			},
		},
	}

	returnedRangeWithPlaceholder := RangeWithPlaceholder{}
	err = clientLSPContext.Call(
		MethodDocumentPrepareRename,
		prepareRenameParams,
		&returnedRangeWithPlaceholder,
	)
	s.Require().NoError(err)
	s.Require().Equal(rangeWithPlaceholder, returnedRangeWithPlaceholder)
}

func (s *HandlerTestSuite) Test_calls_document_linked_editing_range_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	wordPattern := "\\d+"
	linkedEditingRanges := LinkedEditingRanges{
		Ranges: []Range{
			{
				Start: Position{
					Line:      300,
					Character: 24,
				},
				End: Position{
					Line:      300,
					Character: 115,
				},
			},
		},
		WordPattern: &wordPattern,
	}
	serverHandler := NewHandler(
		WithDocumentLinkedEditingRangeHandler(
			func(ctx *common.LSPContext, params *LinkedEditingRangeParams) (*LinkedEditingRanges, error) {
				return &linkedEditingRanges, nil
			},
		),
	)
	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	linkedEditingRangeParams := &LinkedEditingRangeParams{
		TextDocumentPositionParams: TextDocumentPositionParams{
			TextDocument: TextDocumentIdentifier{
				URI: "file:///test_doc_linked_editing_range.go",
			},
			Position: Position{
				Line:      311,
				Character: 24,
			},
		},
	}

	returnedLinkedEditingRanges := LinkedEditingRanges{}
	err = clientLSPContext.Call(
		MethodDocumentLinkedEditingRange,
		linkedEditingRangeParams,
		&returnedLinkedEditingRanges,
	)
	s.Require().NoError(err)
	s.Require().Equal(linkedEditingRanges, returnedLinkedEditingRanges)
}
