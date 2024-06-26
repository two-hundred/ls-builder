package lsp

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/two-hundred/ls-builder/common"
	"github.com/two-hundred/ls-builder/server"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Example() {
	// This example demonstrates how to create an LSP server
	// using the ls-builder LSP 3.17 package.
	// This only initialises a connection as a part of the actual example test,
	// but when wired up to a real client, such as visual studio code,
	// the rest of the handlers will be called as expected.
	logger, err := setupLogger()
	if err != nil {
		log.Fatal(err)
	}

	state := NewState()
	traceService := NewTraceService(logger)
	app := NewApplication(state, traceService, logger)
	app.Setup()

	srv := server.NewServer(app.Handler(), true, logger, nil)

	// A testable example can not use transports over the network,
	// so we use an emulation of a stdio connection to test communication
	// between the client and server instead.
	// In a real-world scenario, you would use a transport like stdio, TCP
	// or WebSockets.
	//
	// It's also worth noting that the Language Server Protocol also expects the
	// server to provide command line arguments to allow the
	// client to select the transport to be used.
	// See: https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#implementationConsiderations
	connContainer := createTestConnectionsContainer(srv.NewHandler())
	go srv.Serve(connContainer.serverConn, logger)

	ctx, cancel := context.WithTimeout(context.Background(), server.DefaultTimeout)
	defer cancel()

	clientLSPContext := server.NewLSPContext(ctx, connContainer.clientConn, nil)

	initialiseParams := InitializeParams{}
	err = json.Unmarshal([]byte(initialiseParamsStr), &initialiseParams)
	if err != nil {
		log.Fatal(err)
	}

	initialiseResult := InitializeResult{}
	err = clientLSPContext.Call(MethodInitialize, initialiseParams, &initialiseResult)
	if err != nil {
		log.Fatal(err)
	}
	err = clientLSPContext.Notify(MethodInitialized, InitializedParams{})
	if err != nil {
		log.Fatal(err)
	}

	// Output:
	// info	new stream connection
	// debug	Initialising server...
}

const Name string = "Test Language Server"

var Version string = "0.1.0"

var ConfigSection string = "test-language-server"

type State struct {
	hasWorkspaceFolderCapability bool
	hasConfigurationCapability   bool
	documentSettings             map[string]*DocSettings
	lock                         sync.Mutex
}

func NewState() *State {
	return &State{
		documentSettings: make(map[string]*DocSettings),
	}
}

type DocSettings struct {
	Trace               DocTraceSettings `json:"trace"`
	MaxNumberOfProblems int              `json:"maxNumberOfProblems"`
}

type DocTraceSettings struct {
	Server string `json:"server"`
}

func (s *State) HasWorkspaceFolderCapability() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.hasWorkspaceFolderCapability
}

func (s *State) SetWorkspaceFolderCapability(value bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.hasWorkspaceFolderCapability = value
}

func (s *State) HasConfigurationCapability() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.hasConfigurationCapability
}

func (s *State) SetConfigurationCapability(value bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.hasConfigurationCapability = value
}

func (s *State) ClearDocSettings() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.documentSettings = make(map[string]*DocSettings)
}

type Application struct {
	handler      *Handler
	state        *State
	logger       *zap.Logger
	traceService *TraceService
}

func NewApplication(state *State, traceService *TraceService, logger *zap.Logger) *Application {
	return &Application{
		state:        state,
		logger:       logger,
		traceService: traceService,
	}
}

func (a *Application) handleInitialise(ctx *common.LSPContext, params *InitializeParams) (any, error) {
	a.logger.Debug("Initialising server...")
	clientCapabilities := params.Capabilities
	capabilities := a.handler.CreateServerCapabilities()
	capabilities.CompletionProvider = &CompletionOptions{}

	hasWorkspaceFolderCapability := clientCapabilities.Workspace != nil && clientCapabilities.Workspace.WorkspaceFolders != nil
	a.state.SetWorkspaceFolderCapability(hasWorkspaceFolderCapability)

	hasConfigurationCapability := clientCapabilities.Workspace != nil && clientCapabilities.Workspace.Configuration != nil
	a.state.SetConfigurationCapability(hasConfigurationCapability)

	result := InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &InitializeResultServerInfo{
			Name:    Name,
			Version: &Version,
		},
	}

	if hasWorkspaceFolderCapability {
		result.Capabilities.Workspace = &ServerWorkspaceCapabilities{
			WorkspaceFolders: &WorkspaceFoldersServerCapabilities{
				Supported: &hasWorkspaceFolderCapability,
			},
		}
	}

	return result, nil
}

func (a *Application) handleInitialised(ctx *common.LSPContext, params *InitializedParams) error {
	if a.state.HasConfigurationCapability() {
		a.handler.SetWorkspaceDidChangeConfigurationHandler(
			a.handleWorkspaceDidChangeConfiguration,
		)
	}
	return nil
}

func (a *Application) handleWorkspaceDidChangeConfiguration(ctx *common.LSPContext, params *DidChangeConfigurationParams) error {
	if a.state.HasConfigurationCapability() {
		// Reset all the cached document settings.
		a.state.ClearDocSettings()
	}

	return nil
}

func (a *Application) handleTextDocumentDidChange(ctx *common.LSPContext, params *DidChangeTextDocumentParams) error {
	dispatcher := NewDispatcher(ctx)
	err := dispatcher.LogMessage(LogMessageParams{
		Type:    MessageTypeInfo,
		Message: "Text document changed (server received)",
	})
	if err != nil {
		return err
	}

	ValidateTextDocument(dispatcher, a.state, params, a.logger)
	return nil
}

func GetDocumentSettings(dispatcher *Dispatcher, state *State, uri string) (*DocSettings, error) {
	state.lock.Lock()
	defer state.lock.Unlock()

	if settings, ok := state.documentSettings[uri]; ok {
		return settings, nil
	} else {
		configResponse := []DocSettings{}
		err := dispatcher.WorkspaceConfiguration(ConfigurationParams{
			Items: []ConfigurationItem{
				{
					ScopeURI: &uri,
					Section:  &ConfigSection,
				},
			},
		}, &configResponse)
		if err != nil {
			return nil, err
		}

		err = dispatcher.LogMessage(LogMessageParams{
			Type:    MessageTypeInfo,
			Message: "document workspace configuration (server received)",
		})
		if err != nil {
			return nil, err
		}

		if len(configResponse) > 0 {
			state.documentSettings[uri] = &configResponse[0]
			return &configResponse[0], nil
		}
	}

	return &DocSettings{
		Trace: DocTraceSettings{
			Server: "off",
		},
		MaxNumberOfProblems: 100,
	}, nil
}

func ValidateTextDocument(
	dispatcher *Dispatcher,
	state *State,
	changeParams *DidChangeTextDocumentParams,
	logger *zap.Logger,
) ([]Diagnostic, error) {
	var diagnostics []Diagnostic
	settings, err := GetDocumentSettings(dispatcher, state, changeParams.TextDocument.URI)
	if err != nil {
		return diagnostics, err

	}
	logger.Debug(fmt.Sprintf("Settings: %v", settings))
	return diagnostics, nil
}

func (a *Application) handleShutdown(ctx *common.LSPContext) error {
	a.logger.Info("Shutting down server...")
	return nil
}

func (a *Application) Setup() {
	a.handler = NewHandler(
		WithInitializeHandler(a.handleInitialise),
		WithInitializedHandler(a.handleInitialised),
		WithShutdownHandler(a.handleShutdown),
		WithTextDocumentDidChangeHandler(a.handleTextDocumentDidChange),
		WithSetTraceHandler(a.traceService.CreateSetTraceHandler()),
	)
}

func setupLogger() (*zap.Logger, error) {
	pe := zap.NewProductionEncoderConfig()
	pe.EncodeTime = nil
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(pe),
		zapcore.AddSync(os.Stdout),
		zap.DebugLevel,
	)
	logger := zap.New(core)
	return logger, nil
}

func (a *Application) Handler() *Handler {
	return a.handler
}

const (
	initialiseParamsStr = `
{
    "processId": 18301,
    "clientInfo": {
        "name": "Visual Studio Code",
        "version": "1.90.2"
    },
    "locale": "en",
    "rootPath": "/Users/testuser/sandbox-vscode-extension/server-example-test",
    "rootUri": "file:///Users/testuser/sandbox-vscode-extension/server-example-test",
    "capabilities": {
        "workspace": {
            "applyEdit": true,
            "workspaceEdit": {
                "documentChanges": true,
                "resourceOperations": [
                    "create",
                    "rename",
                    "delete"
                ],
                "failureHandling": "textOnlyTransactional",
                "normalizesLineEndings": true,
                "changeAnnotationSupport": {
                    "groupsOnLabel": true
                }
            },
            "configuration": true,
            "didChangeWatchedFiles": {
                "dynamicRegistration": true,
                "relativePatternSupport": true
            },
            "symbol": {
                "dynamicRegistration": true,
                "symbolKind": {
                    "valueSet": [
                        1,
                        2,
                        3,
                        4,
                        5,
                        6,
                        7,
                        8,
                        9,
                        10,
                        11,
                        12,
                        13,
                        14,
                        15,
                        16,
                        17,
                        18,
                        19,
                        20,
                        21,
                        22,
                        23,
                        24,
                        25,
                        26
                    ]
                },
                "tagSupport": {
                    "valueSet": [
                        1
                    ]
                },
                "resolveSupport": {
                    "properties": [
                        "location.range"
                    ]
                }
            },
            "codeLens": {
                "refreshSupport": true
            },
            "executeCommand": {
                "dynamicRegistration": true
            },
            "didChangeConfiguration": {
                "dynamicRegistration": true
            },
            "workspaceFolders": true,
            "foldingRange": {
                "refreshSupport": true
            },
            "semanticTokens": {
                "refreshSupport": true
            },
            "fileOperations": {
                "dynamicRegistration": true,
                "didCreate": true,
                "didRename": true,
                "didDelete": true,
                "willCreate": true,
                "willRename": true,
                "willDelete": true
            },
            "inlineValue": {
                "refreshSupport": true
            },
            "inlayHint": {
                "refreshSupport": true
            },
            "diagnostics": {
                "refreshSupport": true
            }
        },
        "textDocument": {
            "publishDiagnostics": {
                "relatedInformation": true,
                "versionSupport": false,
                "tagSupport": {
                    "valueSet": [
                        1,
                        2
                    ]
                },
                "codeDescriptionSupport": true,
                "dataSupport": true
            },
            "synchronization": {
                "dynamicRegistration": true,
                "willSave": true,
                "willSaveWaitUntil": true,
                "didSave": true
            },
            "completion": {
                "dynamicRegistration": true,
                "contextSupport": true,
                "completionItem": {
                    "snippetSupport": true,
                    "commitCharactersSupport": true,
                    "documentationFormat": [
                        "markdown",
                        "plaintext"
                    ],
                    "deprecatedSupport": true,
                    "preselectSupport": true,
                    "tagSupport": {
                        "valueSet": [
                            1
                        ]
                    },
                    "insertReplaceSupport": true,
                    "resolveSupport": {
                        "properties": [
                            "documentation",
                            "detail",
                            "additionalTextEdits"
                        ]
                    },
                    "insertTextModeSupport": {
                        "valueSet": [
                            1,
                            2
                        ]
                    },
                    "labelDetailsSupport": true
                },
                "insertTextMode": 2,
                "completionItemKind": {
                    "valueSet": [
                        1,
                        2,
                        3,
                        4,
                        5,
                        6,
                        7,
                        8,
                        9,
                        10,
                        11,
                        12,
                        13,
                        14,
                        15,
                        16,
                        17,
                        18,
                        19,
                        20,
                        21,
                        22,
                        23,
                        24,
                        25
                    ]
                },
                "completionList": {
                    "itemDefaults": [
                        "commitCharacters",
                        "editRange",
                        "insertTextFormat",
                        "insertTextMode",
                        "data"
                    ]
                }
            },
            "hover": {
                "dynamicRegistration": true,
                "contentFormat": [
                    "markdown",
                    "plaintext"
                ]
            },
            "signatureHelp": {
                "dynamicRegistration": true,
                "signatureInformation": {
                    "documentationFormat": [
                        "markdown",
                        "plaintext"
                    ],
                    "parameterInformation": {
                        "labelOffsetSupport": true
                    },
                    "activeParameterSupport": true
                },
                "contextSupport": true
            },
            "definition": {
                "dynamicRegistration": true,
                "linkSupport": true
            },
            "references": {
                "dynamicRegistration": true
            },
            "documentHighlight": {
                "dynamicRegistration": true
            },
            "documentSymbol": {
                "dynamicRegistration": true,
                "symbolKind": {
                    "valueSet": [
                        1,
                        2,
                        3,
                        4,
                        5,
                        6,
                        7,
                        8,
                        9,
                        10,
                        11,
                        12,
                        13,
                        14,
                        15,
                        16,
                        17,
                        18,
                        19,
                        20,
                        21,
                        22,
                        23,
                        24,
                        25,
                        26
                    ]
                },
                "hierarchicalDocumentSymbolSupport": true,
                "tagSupport": {
                    "valueSet": [
                        1
                    ]
                },
                "labelSupport": true
            },
            "codeAction": {
                "dynamicRegistration": true,
                "isPreferredSupport": true,
                "disabledSupport": true,
                "dataSupport": true,
                "resolveSupport": {
                    "properties": [
                        "edit"
                    ]
                },
                "codeActionLiteralSupport": {
                    "codeActionKind": {
                        "valueSet": [
                            "",
                            "quickfix",
                            "refactor",
                            "refactor.extract",
                            "refactor.inline",
                            "refactor.rewrite",
                            "source",
                            "source.organizeImports"
                        ]
                    }
                },
                "honorsChangeAnnotations": true
            },
            "codeLens": {
                "dynamicRegistration": true
            },
            "formatting": {
                "dynamicRegistration": true
            },
            "rangeFormatting": {
                "dynamicRegistration": true,
                "rangesSupport": true
            },
            "onTypeFormatting": {
                "dynamicRegistration": true
            },
            "rename": {
                "dynamicRegistration": true,
                "prepareSupport": true,
                "prepareSupportDefaultBehavior": 1,
                "honorsChangeAnnotations": true
            },
            "documentLink": {
                "dynamicRegistration": true,
                "tooltipSupport": true
            },
            "typeDefinition": {
                "dynamicRegistration": true,
                "linkSupport": true
            },
            "implementation": {
                "dynamicRegistration": true,
                "linkSupport": true
            },
            "colorProvider": {
                "dynamicRegistration": true
            },
            "foldingRange": {
                "dynamicRegistration": true,
                "rangeLimit": 5000,
                "lineFoldingOnly": true,
                "foldingRangeKind": {
                    "valueSet": [
                        "comment",
                        "imports",
                        "region"
                    ]
                },
                "foldingRange": {
                    "collapsedText": false
                }
            },
            "declaration": {
                "dynamicRegistration": true,
                "linkSupport": true
            },
            "selectionRange": {
                "dynamicRegistration": true
            },
            "callHierarchy": {
                "dynamicRegistration": true
            },
            "semanticTokens": {
                "dynamicRegistration": true,
                "tokenTypes": [
                    "namespace",
                    "type",
                    "class",
                    "enum",
                    "interface",
                    "struct",
                    "typeParameter",
                    "parameter",
                    "variable",
                    "property",
                    "enumMember",
                    "event",
                    "function",
                    "method",
                    "macro",
                    "keyword",
                    "modifier",
                    "comment",
                    "string",
                    "number",
                    "regexp",
                    "operator",
                    "decorator"
                ],
                "tokenModifiers": [
                    "declaration",
                    "definition",
                    "readonly",
                    "static",
                    "deprecated",
                    "abstract",
                    "async",
                    "modification",
                    "documentation",
                    "defaultLibrary"
                ],
                "formats": [
                    "relative"
                ],
                "requests": {
                    "range": true,
                    "full": {
                        "delta": true
                    }
                },
                "multilineTokenSupport": false,
                "overlappingTokenSupport": false,
                "serverCancelSupport": true,
                "augmentsSyntaxTokens": true
            },
            "linkedEditingRange": {
                "dynamicRegistration": true
            },
            "typeHierarchy": {
                "dynamicRegistration": true
            },
            "inlineValue": {
                "dynamicRegistration": true
            },
            "inlayHint": {
                "dynamicRegistration": true,
                "resolveSupport": {
                    "properties": [
                        "tooltip",
                        "textEdits",
                        "label.tooltip",
                        "label.location",
                        "label.command"
                    ]
                }
            },
            "diagnostic": {
                "dynamicRegistration": true,
                "relatedDocumentSupport": false
            }
        },
        "window": {
            "showMessage": {
                "messageActionItem": {
                    "additionalPropertiesSupport": true
                }
            },
            "showDocument": {
                "support": true
            },
            "workDoneProgress": true
        },
        "general": {
            "staleRequestSupport": {
                "cancel": true,
                "retryOnContentModified": [
                    "textDocument/semanticTokens/full",
                    "textDocument/semanticTokens/range",
                    "textDocument/semanticTokens/full/delta"
                ]
            },
            "regularExpressions": {
                "engine": "ECMAScript",
                "version": "ES2020"
            },
            "markdown": {
                "parser": "marked",
                "version": "1.1.0"
            },
            "positionEncodings": [
                "utf-16"
            ]
        },
        "notebookDocument": {
            "synchronization": {
                "dynamicRegistration": true,
                "executionSummarySupport": true
            }
        }
    },
    "trace": "verbose",
    "workspaceFolders": [
        {
            "uri": "file:///Users/testuser/sandbox-vscode-extension/server-example-test",
            "name": "server-example-test"
        }
    ]
}
	`
)
