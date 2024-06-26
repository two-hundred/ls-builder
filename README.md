# ls-builder

[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=two-hundred_ls-builder&metric=coverage)](https://sonarcloud.io/summary/new_code?id=two-hundred_ls-builder)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=two-hundred_ls-builder&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=two-hundred_ls-builder)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=two-hundred_ls-builder&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=two-hundred_ls-builder)
[![Go Reference](https://pkg.go.dev/badge/github.com/two-hundred/ls-builder.svg)](https://pkg.go.dev/github.com/two-hundred/ls-builder)


Language Server Builder (SDK) for Go. This is a Go implementation of the [Language Server Protocol](https://microsoft.github.io/language-server-protocol/)'s server component.

The supported LSP implementations are:

- [Language Server Protocol 3.17.0](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/)

## Example Application

```go
package lsp

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/two-hundred/ls-builder/common"
	"github.com/two-hundred/ls-builder/server"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

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

func setupLogger() (*zap.Logger, *os.File, error) {
	logFileHandle, err := os.OpenFile("test-ls.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, nil, err
	}
	writerSync := zapcore.NewMultiWriteSyncer(
		// stdout and stdin are used for communication with the client
		// and should not be logged to.
		zapcore.AddSync(os.Stderr),
		zapcore.AddSync(logFileHandle),
	)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		writerSync,
		zap.DebugLevel,
	)
	logger := zap.New(core)
	return logger, logFileHandle, nil
}

func (a *Application) Handler() *Handler {
	return a.handler
}

func main() {
	logger, logFile, err := setupLogger()
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	state := NewState()
	traceService := NewTraceService(logger)
	app := NewApplication(state, traceService, logger)
	app.Setup()

	srv := server.NewServer(app.Handler(), true, logger, nil)

	stdio := server.Stdio{}
	conn := server.NewStreamConnection(srv.NewHandler(), stdio)
	srv.Serve(conn, logger)
}
```

## Additional documentation

- [Contributing](CONTRIBUTING.md)
