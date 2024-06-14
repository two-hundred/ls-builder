package lsp

import (
	"encoding/json"
)

type CompletionClientCapabilities struct{}

type HoverClientCapabilities struct{}

type SignatureHelpClientCapabilities struct{}

type DeclarationClientCapabilities struct{}

type DefinitionClientCapabilities struct{}

type TypeDefinitionClientCapabilities struct{}

type ImplementationClientCapabilities struct{}

type ReferenceClientCapabilities struct{}

type DocumentHighlightClientCapabilities struct{}

type DocumentSymbolClientCapabilities struct{}

type CodeActionClientCapabilities struct{}

type CodeLensClientCapabilities struct{}

type DocumentLinkClientCapabilities struct{}

type DocumentColorClientCapabilities struct{}

type DocumentFormattingClientCapabilities struct{}

type DocumentRangeFormattingClientCapabilities struct{}

type DocumentOnTypeFormattingClientCapabilities struct{}

type RenameClientCapabilities struct{}

type FoldingRangeClientCapabilities struct{}

type SelectionRangeClientCapabilities struct{}

type LinkedEditingRangeClientCapabilities struct{}

type CallHierarchyClientCapabilities struct{}

type SemanticTokensClientCapabilities struct{}

type MonikerClientCapabilities struct{}

type InlayHintClientCapabilities struct{}

type InlineValueClientCapabilities struct{}

type TypeHierarchyClientCapabilities struct{}

type CompletionOptions struct{}

type SignatureHelpOptions struct{}

// CodeLensOptions provides server capability options for code lens requests.
type CodeLensOptions struct {
	WorkDoneProgressOptions

	// Code lens has a resolve provider as well.
	ResolveProvider *bool `json:"resolveProvider,omitempty"`
}

// DocumentLinkOptions provides server capability options for document link requests.
type DocumentLinkOptions struct {
	WorkDoneProgressOptions

	// Code lens has a resolve provider as well.
	ResolveProvider *bool `json:"resolveProvider,omitempty"`
}

// DocumentOnTypeFormattingOptions provides server capability options for
// document on type formatting requests.
type DocumentOnTypeFormattingOptions struct {
	// A character on which formatting should be triggered, like `}`.
	FirstTriggerCharacter string `json:"firstTriggerCharacter"`

	// More trigger characters.
	MoreTriggerCharacter []string `json:"moreTriggerCharacter,omitempty"`
}

func unmarshalLanguageFeatureServerCapabilities(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if err := unmarshalServerCapabilityHoverProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilityDeclarationProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilityDefinitionProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilityTypeDefinitionProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilityImplementationProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilityReferencesProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilityDocumentHighlightProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilityDocumentSymbolProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilityCodeActionProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilityColorProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilityDocumentFormattingProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilityDocumentRangeFormattingProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	if err := unmarshalServerCapabilityRenameProvider(serverCapabilities, intermediate); err != nil {
		return err
	}

	return nil
}

// HoverOptions provides server capability options for hover requests.
type HoverOptions struct {
	WorkDoneProgressOptions
}

// unmarshals the HoverProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
func unmarshalServerCapabilityHoverProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.HoverProvider == nil {
		return nil
	}

	var optVal HoverOptions
	if err := json.Unmarshal(intermediate.HoverProvider, &optVal); err == nil {
		serverCapabilities.HoverProvider = optVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.HoverProvider, &boolVal); err == nil {
			serverCapabilities.HoverProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// DeclarationOptions provides server capability options for goto declaration requests.
type DeclarationOptions struct {
	WorkDoneProgressOptions
}

// DeclarationRegistrationOptions provides server capability registration
// options for goto declaration requests.
type DeclarationRegistrationOptions struct {
	DeclarationOptions
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
}

// unmarshals the DeclarationProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
// DeclarationRegistrationOptions is a superset of DeclarationOptions
// so we only need to unmarshal to the former with nil values for the
// empty fields.
func unmarshalServerCapabilityDeclarationProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.DeclarationProvider == nil {
		return nil
	}

	var optRegVal DeclarationRegistrationOptions
	if err := json.Unmarshal(intermediate.DeclarationProvider, &optRegVal); err == nil {
		serverCapabilities.DeclarationProvider = optRegVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.DeclarationProvider, &boolVal); err == nil {
			serverCapabilities.DeclarationProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// DefinitionOptions provides server capability options for goto definition requests.
type DefinitionOptions struct {
	WorkDoneProgressOptions
}

// unmarshals the DefinitionProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
func unmarshalServerCapabilityDefinitionProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.DefinitionProvider == nil {
		return nil
	}

	var optVal DefinitionOptions
	if err := json.Unmarshal(intermediate.DefinitionProvider, &optVal); err == nil {
		serverCapabilities.DefinitionProvider = optVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.DefinitionProvider, &boolVal); err == nil {
			serverCapabilities.DefinitionProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// TypeDefinitionOptions provides server capability options for goto type definition requests.
type TypeDefinitionOptions struct {
	WorkDoneProgressOptions
}

// TypeDefinitionRegistrationOptions provides server capability registration
// options for goto type definition requests.
type TypeDefinitionRegistrationOptions struct {
	TypeDefinitionOptions
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
}

// unmarshals the TypeDefinitionProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
// TypeDefinitionRegistrationOptions is a superset of TypeDefinitionOptions
// so we only need to unmarshal to the former with nil values for the
// empty fields.
func unmarshalServerCapabilityTypeDefinitionProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.TypeDefinitionProvider == nil {
		return nil
	}

	var optRegVal TypeDefinitionRegistrationOptions
	if err := json.Unmarshal(intermediate.TypeDefinitionProvider, &optRegVal); err == nil {
		serverCapabilities.TypeDefinitionProvider = optRegVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.TypeDefinitionProvider, &boolVal); err == nil {
			serverCapabilities.TypeDefinitionProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// ImplementationOptions provides server capability options for goto implementation requests.
type ImplementationOptions struct {
	WorkDoneProgressOptions
}

// ImplementationRegistrationOptions provides server capability registration
// options for goto implementation requests.
type ImplementationRegistrationOptions struct {
	ImplementationOptions
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
}

// unmarshals the ImplementationProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
// ImplementationRegistrationOptions is a superset of ImplementationOptions
// so we only need to unmarshal to the former with nil values for the
// empty fields.
func unmarshalServerCapabilityImplementationProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.ImplementationProvider == nil {
		return nil
	}

	var optRegVal ImplementationRegistrationOptions
	if err := json.Unmarshal(intermediate.ImplementationProvider, &optRegVal); err == nil {
		serverCapabilities.ImplementationProvider = optRegVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.ImplementationProvider, &boolVal); err == nil {
			serverCapabilities.ImplementationProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// ReferenceOptions provides server capability options for find references requests.
type ReferenceOptions struct {
	WorkDoneProgressOptions
}

// unmarshals the ReferencesProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
func unmarshalServerCapabilityReferencesProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.ReferencesProvider == nil {
		return nil
	}

	var optVal ReferenceOptions
	if err := json.Unmarshal(intermediate.ReferencesProvider, &optVal); err == nil {
		serverCapabilities.ReferencesProvider = optVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.ReferencesProvider, &boolVal); err == nil {
			serverCapabilities.ReferencesProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// DocumentHighlightOptions provides server capability options for document highlight requests.
type DocumentHighlightOptions struct {
	WorkDoneProgressOptions
}

// unmarshals the DocumentHighlightProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
func unmarshalServerCapabilityDocumentHighlightProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.DocumentHighlightProvider == nil {
		return nil
	}

	var optVal DocumentHighlightOptions
	if err := json.Unmarshal(intermediate.DocumentHighlightProvider, &optVal); err == nil {
		serverCapabilities.DocumentHighlightProvider = optVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.DocumentHighlightProvider, &boolVal); err == nil {
			serverCapabilities.DocumentHighlightProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// DocumentSymbolOptions provides server capability options for document symbol requests.
type DocumentSymbolOptions struct {
	WorkDoneProgressOptions

	// A human-readable string that is shown when multiple outline trees
	// are shown for the same document.
	//
	// @since 3.16.0
	Label *string `json:"label,omitempty"`
}

// unmarshals the DocumentSymbolProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
func unmarshalServerCapabilityDocumentSymbolProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.DocumentSymbolProvider == nil {
		return nil
	}

	var optVal DocumentSymbolOptions
	if err := json.Unmarshal(intermediate.DocumentSymbolProvider, &optVal); err == nil {
		serverCapabilities.DocumentSymbolProvider = optVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.DocumentSymbolProvider, &boolVal); err == nil {
			serverCapabilities.DocumentSymbolProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// CodeActionOptions provides server capability options for code action requests.
type CodeActionOptions struct {
	WorkDoneProgressOptions

	// CodeActionKinds that this server may return.
	//
	// The list of kinds may be generic, such as `CodeActionKind.Refactor`,
	// or the server may list out every specific kind they provide.
	CodeActionKinds []CodeActionKind `json:"codeActionKinds,omitempty"`

	// The server provides support to resolve additional
	// information for a code action.
	//
	// @since 3.16.0
	ResolveProvider *bool `json:"resolveProvider,omitempty"`
}

// The kind of a code action.
//
// Kinds are a hierarchical list of identifiers separated by `.`,
// e.g. `"refactor.extract.function"`.
//
// The set of kinds is open and client needs to announce the kinds it
// supports to the server during initialization.
type CodeActionKind = string

const (
	// CodeActionKindEmpty for the empty kind.
	CodeActionKindEmpty CodeActionKind = ""

	// CodeActionKindQuickFix is the base kind
	// for quickfix actions: 'quickfix'.
	CodeActionKindQuickFix CodeActionKind = "quickfix"

	// CodeActionKindRefactor is the base kind
	// for refactoring actions: 'refactor'.
	CodeActionKindRefactor CodeActionKind = "refactor"

	// CodeActionKindRefactorExtract is the kind for
	// refactoring extraction actions: 'refactor.extract'.
	//
	// Example extract actions:
	//
	// - Extract method
	// - Extract function
	// - Extract variable
	// - Extract interface from class
	// - ...
	CodeActionKindRefactorExtract CodeActionKind = "refactor.extract"

	// CodeActionKindRefactorInline is the base kind for
	// refactoring inline actions: 'refactor.inline'.
	//
	// Example inline actions:
	//
	// - Inline function
	// - Inline variable
	// - Inline constant
	// - ...
	CodeActionKindRefactorInline CodeActionKind = "refactor.inline"

	// CodeActionKindRefactorRewrite is the base kind for
	// refactoring rewrite actions: 'refactor.rewrite'.
	//
	// Example rewrite actions:
	//
	// - Convert JavaScript function to class
	// - Add or remove parameter
	// - Encapsulate field
	// - Make method static
	// - Move method to base class
	// - ...
	CodeActionKindRefactorRewrite CodeActionKind = "refactor.rewrite"

	// CodeActionKindSource is the base kind for
	// source actions: `source`.
	//
	// Source code actions apply to the entire file.
	CodeActionKindSource CodeActionKind = "source"

	// CodeActionKindSourceOrganizeImports is the base kind for
	// an organize imports source action: `source.organizeImports`.
	CodeActionKindSourceOrganizeImports CodeActionKind = "source.organizeImports"

	// CodeActionKindSourceFixAll is the base kind for
	// a 'fix all' source action: `source.fixAll`.
	//
	// 'Fix all' actions automatically fix errors that hae a clear fix that
	// do not require user input. They should not suppress errors or perform
	// unsafe fixes such as generating new types or classes.
	//
	// @since 3.17.0
	CodeActionKindSourceFixAll CodeActionKind = "source.fixAll"
)

// unmarshals the CodeActionProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
func unmarshalServerCapabilityCodeActionProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.CodeActionProvider == nil {
		return nil
	}

	var optVal CodeActionOptions
	if err := json.Unmarshal(intermediate.CodeActionProvider, &optVal); err == nil {
		serverCapabilities.CodeActionProvider = optVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.CodeActionProvider, &boolVal); err == nil {
			serverCapabilities.CodeActionProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// DocumentColorOptions provides server capability options for document color requests.
type DocumentColorOptions struct {
	WorkDoneProgressOptions
}

// DocumentRegistrationOptions provides server capability registration
// options for document color requests.
type DocumentColorRegistrationOptions struct {
	DocumentColorOptions
	TextDocumentRegistrationOptions
	StaticRegistrationOptions
}

// unmarshals the ColorProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
// DocumentColorRegistrationOptions is a superset of DocumentColorOptions
// so we only need to unmarshal to the former with nil values for the
// empty fields.
func unmarshalServerCapabilityColorProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.ColorProvider == nil {
		return nil
	}

	var optRegVal DocumentColorRegistrationOptions
	if err := json.Unmarshal(intermediate.ColorProvider, &optRegVal); err == nil {
		serverCapabilities.ColorProvider = optRegVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.ColorProvider, &boolVal); err == nil {
			serverCapabilities.ColorProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// DocumentFormattingOptions provides server capability options for document formatting requests.
type DocumentFormattingOptions struct {
	WorkDoneProgressOptions
}

// unmarshals the DocumentFormattingProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
func unmarshalServerCapabilityDocumentFormattingProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.DocumentFormattingProvider == nil {
		return nil
	}

	var optVal DocumentFormattingOptions
	if err := json.Unmarshal(intermediate.DocumentFormattingProvider, &optVal); err == nil {
		serverCapabilities.DocumentFormattingProvider = optVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.DocumentFormattingProvider, &boolVal); err == nil {
			serverCapabilities.DocumentFormattingProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// DocumentRangeFormattingOptions provides server capability options for
// document range formatting requests.
type DocumentRangeFormattingOptions struct {
	WorkDoneProgressOptions
}

// unmarshals the DocumentRangeFormattingProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
func unmarshalServerCapabilityDocumentRangeFormattingProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.DocumentRangeFormattingProvider == nil {
		return nil
	}

	var optVal DocumentRangeFormattingOptions
	if err := json.Unmarshal(intermediate.DocumentRangeFormattingProvider, &optVal); err == nil {
		serverCapabilities.DocumentRangeFormattingProvider = optVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.DocumentRangeFormattingProvider, &boolVal); err == nil {
			serverCapabilities.DocumentRangeFormattingProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}

// RenameOptions provides server capability options for rename requests.
type RenameOptions struct {
	WorkDoneProgressOptions

	// Renames should be checked and tested before being executed.
	PrepareProvider *bool `json:"prepareProvider,omitempty"`
}

// unmarshals the RenameProvider
// field of a server capabilities object.
// This modifies the serverCapabilities object.
func unmarshalServerCapabilityRenameProvider(
	serverCapabilities *ServerCapabilities,
	intermediate *serverCapabilitiesIntermediate,
) error {
	if intermediate.RenameProvider == nil {
		return nil
	}

	var optVal RenameOptions
	if err := json.Unmarshal(intermediate.RenameProvider, &optVal); err == nil {
		serverCapabilities.RenameProvider = optVal
	} else {
		var boolVal bool
		if err := json.Unmarshal(intermediate.RenameProvider, &boolVal); err == nil {
			serverCapabilities.RenameProvider = boolVal
		} else {
			return err
		}
	}

	return nil
}
