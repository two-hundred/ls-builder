package lsp

import (
	"context"

	"github.com/two-hundred/ls-builder/common"
	"github.com/two-hundred/ls-builder/server"
	"go.uber.org/zap"
)

func (s *HandlerTestSuite) Test_calls_document_diagnostics_request_handler_receives_full_results() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	resultID := "test-result-1"
	related1ResultID := "related-test-result-1"
	related2ResultID := "related-test-result-2"
	diagnosticSeverity := DiagnosticSeverityError
	testDiagnosticsErrorCode := "TestDiagnosticsErrorCode"
	diagnostics := RelatedFullDocumentDiagnosticReport{
		FullDocumentDiagnosticReport: FullDocumentDiagnosticReport{
			Kind:     DocumentDiagnosticReportKindFull,
			ResultID: &resultID,
			Items: []Diagnostic{
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
					Severity: &diagnosticSeverity,
					Code: &IntOrString{
						StrVal: &testDiagnosticsErrorCode,
					},
					Message: "Test diagnostic message",
					Tags: []DiagnosticTag{
						DiagnosticTagUnnecessary,
					},
				},
			},
		},
		RelatedDocuments: map[string]any{
			"file:///test_document.go": FullDocumentDiagnosticReport{
				Kind:     DocumentDiagnosticReportKindFull,
				ResultID: &related1ResultID,
				Items: []Diagnostic{
					{
						Range: Range{
							Start: Position{
								Line:      140,
								Character: 15,
							},
							End: Position{
								Line:      140,
								Character: 24,
							},
						},
						Severity: &diagnosticSeverity,
						Code: &IntOrString{
							StrVal: &testDiagnosticsErrorCode,
						},
						Message: "Test related diagnostic message",
						Tags: []DiagnosticTag{
							DiagnosticTagUnnecessary,
						},
					},
				},
			},
			"file:///test_document2.go": UnchangedDocumentDiagnosticReport{
				Kind:     DocumentDiagnosticReportKindUnchanged,
				ResultID: related2ResultID,
			},
		},
	}
	serverHandler := NewHandler(
		WithDocumentDiagnosticsHandler(
			func(ctx *common.LSPContext, params *DocumentDiagnosticParams) (any, error) {
				return diagnostics, nil
			},
		),
	)

	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	identifier := "test-identifier"
	prevResultID := "prev-result-id"
	diagnosticParams := DocumentDiagnosticParams{
		TextDocument: TextDocumentIdentifier{
			URI: "file:///test_document.go",
		},
		Identifier:       &identifier,
		PreviousResultID: &prevResultID,
	}
	returnedDiagnostics := RelatedFullDocumentDiagnosticReport{}
	err = clientLSPContext.Call(
		MethodDocumentDiagnostic,
		diagnosticParams,
		&returnedDiagnostics,
	)
	s.Require().NoError(err)
	s.Require().Equal(diagnostics, returnedDiagnostics)
}

func (s *HandlerTestSuite) Test_calls_document_diagnostics_request_handler_receives_unchanged_results() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	resultID := "test-unchanged-result-1"
	related1ResultID := "related-test-result-3"
	related2ResultID := "related-test-result-4"
	diagnosticSeverity := DiagnosticSeverityError
	testDiagnosticsErrorCode := "TestDiagnosticsErrorCode"
	diagnostics := RelatedUnchangedDocumentDiagnosticReport{
		UnchangedDocumentDiagnosticReport: UnchangedDocumentDiagnosticReport{
			Kind:     DocumentDiagnosticReportKindUnchanged,
			ResultID: resultID,
		},
		RelatedDocuments: map[string]any{
			"file:///test_document.go": FullDocumentDiagnosticReport{
				Kind:     DocumentDiagnosticReportKindFull,
				ResultID: &related1ResultID,
				Items: []Diagnostic{
					{
						Range: Range{
							Start: Position{
								Line:      140,
								Character: 15,
							},
							End: Position{
								Line:      140,
								Character: 24,
							},
						},
						Severity: &diagnosticSeverity,
						Code: &IntOrString{
							StrVal: &testDiagnosticsErrorCode,
						},
						Message: "Test related diagnostic message",
						Tags: []DiagnosticTag{
							DiagnosticTagUnnecessary,
						},
					},
				},
			},
			"file:///test_document2.go": UnchangedDocumentDiagnosticReport{
				Kind:     DocumentDiagnosticReportKindUnchanged,
				ResultID: related2ResultID,
			},
		},
	}
	serverHandler := NewHandler(
		WithDocumentDiagnosticsHandler(
			func(ctx *common.LSPContext, params *DocumentDiagnosticParams) (any, error) {
				return diagnostics, nil
			},
		),
	)

	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	identifier := "test-identifier-unchanged"
	prevResultID := "prev-result-id-unchanged"
	diagnosticParams := DocumentDiagnosticParams{
		TextDocument: TextDocumentIdentifier{
			URI: "file:///test_document.go",
		},
		Identifier:       &identifier,
		PreviousResultID: &prevResultID,
	}
	returnedDiagnostics := RelatedUnchangedDocumentDiagnosticReport{}
	err = clientLSPContext.Call(
		MethodDocumentDiagnostic,
		diagnosticParams,
		&returnedDiagnostics,
	)
	s.Require().NoError(err)
	s.Require().Equal(diagnostics, returnedDiagnostics)
}

func (s *HandlerTestSuite) Test_calls_document_diagnostics_request_handler_receives_partial_results() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	related1ResultID := "related-test-result-1"
	related2ResultID := "related-test-result-2"
	diagnosticSeverity := DiagnosticSeverityError
	testDiagnosticsErrorCode := "TestDiagnosticsErrorCode"
	diagnostics := DocumentDiagnosticReportPartialResult{
		RelatedDocuments: map[string]any{
			// For partial results, the first send should be the full document diagnostics
			// (DocumentDiagnosticReport) and the subsequent sends should be
			// DocumentDiagnosticPartialReport.
			"file:///test_document.go": FullDocumentDiagnosticReport{
				Kind:     DocumentDiagnosticReportKindFull,
				ResultID: &related1ResultID,
				Items: []Diagnostic{
					{
						Range: Range{
							Start: Position{
								Line:      140,
								Character: 15,
							},
							End: Position{
								Line:      140,
								Character: 24,
							},
						},
						Severity: &diagnosticSeverity,
						Code: &IntOrString{
							StrVal: &testDiagnosticsErrorCode,
						},
						Message: "Test related diagnostic message",
						Tags: []DiagnosticTag{
							DiagnosticTagUnnecessary,
						},
					},
				},
			},
			"file:///test_document2.go": UnchangedDocumentDiagnosticReport{
				Kind:     DocumentDiagnosticReportKindUnchanged,
				ResultID: related2ResultID,
			},
		},
	}
	serverHandler := NewHandler(
		WithDocumentDiagnosticsHandler(
			func(ctx *common.LSPContext, params *DocumentDiagnosticParams) (any, error) {
				return diagnostics, nil
			},
		),
	)

	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	identifier := "test-identifier-partial"
	prevResultID := "prev-result-id-partial"
	diagnosticParams := DocumentDiagnosticParams{
		TextDocument: TextDocumentIdentifier{
			URI: "file:///test_document_3.go",
		},
		Identifier:       &identifier,
		PreviousResultID: &prevResultID,
	}
	returnedDiagnostics := DocumentDiagnosticReportPartialResult{}
	err = clientLSPContext.Call(
		MethodDocumentDiagnostic,
		diagnosticParams,
		&returnedDiagnostics,
	)
	s.Require().NoError(err)
	s.Require().Equal(diagnostics, returnedDiagnostics)
}

func (s *HandlerTestSuite) Test_calls_workspace_diagnostics_request_handler() {
	logger, err := zap.NewDevelopment()
	s.Require().NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	result1ID := "test-result-1"
	report1Version := Integer(1)
	result2ID := "test-result-2"
	report2Version := Integer(2)
	diagnosticSeverity := DiagnosticSeverityError
	testDiagnosticsErrorCode := "TestDiagnosticsErrorCode"
	diagnostics := WorkspaceDiagnosticReport{
		Items: []any{
			WorkspaceFullDocumentDiagnosticReport{
				FullDocumentDiagnosticReport: FullDocumentDiagnosticReport{
					Kind:     DocumentDiagnosticReportKindFull,
					ResultID: &result1ID,
					Items: []Diagnostic{
						{
							Range: Range{
								Start: Position{
									Line:      230,
									Character: 15,
								},
								End: Position{
									Line:      230,
									Character: 24,
								},
							},
							Severity: &diagnosticSeverity,
							Code: &IntOrString{
								StrVal: &testDiagnosticsErrorCode,
							},
							Message: "Test diagnostic message",
							Tags: []DiagnosticTag{
								DiagnosticTagUnnecessary,
							},
						},
					},
				},
				URI:     "file:///test_document.go",
				Version: &report1Version,
			},
			WorkspaceUnchangedDocumentDiagnosticReport{
				UnchangedDocumentDiagnosticReport: UnchangedDocumentDiagnosticReport{
					Kind:     DocumentDiagnosticReportKindUnchanged,
					ResultID: result2ID,
				},
				URI:     "file:///test_document2.go",
				Version: &report2Version,
			},
		},
	}
	serverHandler := NewHandler(
		WithWorkspaceDiagnosticHandler(
			func(ctx *common.LSPContext, params *WorkspaceDiagnosticParams) (*WorkspaceDiagnosticReport, error) {
				return &diagnostics, nil
			},
		),
	)

	// Emulate the LSP initialisation process.
	serverHandler.SetInitialized(true)
	srv := server.NewServer(serverHandler, true, nil, nil)

	container := createTestConnectionsContainer(srv.NewHandler())

	go srv.Serve(container.serverConn, logger)

	clientLSPContext := server.NewLSPContext(ctx, container.clientConn, nil)

	identifier := "test-identifier-full"
	diagnosticParams := WorkspaceDiagnosticParams{
		Identifier: &identifier,
	}
	returnedDiagnostics := WorkspaceDiagnosticReport{}
	err = clientLSPContext.Call(
		MethodWorkspaceDiagnostic,
		diagnosticParams,
		&returnedDiagnostics,
	)
	s.Require().NoError(err)
	s.Require().Equal(diagnostics, returnedDiagnostics)
}
