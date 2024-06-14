package lsp

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type LanguageFeaturesTestSuite struct {
	suite.Suite
}

func (s *LanguageFeaturesTestSuite) Test_unmarshal_hover_provider_capability() {
	workDoneProgress := true
	tests := []serverCapabilityFixture{
		{
			name:  "unmarshals hover options",
			input: "{\"hoverProvider\":{\"workDoneProgress\":true}}",
			expected: &ServerCapabilities{
				HoverProvider: HoverOptions{
					WorkDoneProgressOptions: WorkDoneProgressOptions{
						WorkDoneProgress: &workDoneProgress,
					},
				},
			},
		},
		{
			name:  "unmarshals boolean",
			input: "{\"hoverProvider\":true}",
			expected: &ServerCapabilities{
				HoverProvider: true,
			},
		},
		{
			name:     "unmarshals missing value to nil",
			input:    "{}",
			expected: &ServerCapabilities{},
		},
	}

	testServerCapabilities(&s.Suite, tests)
}

func (s *LanguageFeaturesTestSuite) Test_unmarshal_declaration_provider_capability() {
	workDoneProgress := true
	testRegisterID := "test-register-id"
	language := "typescript"
	fileScheme := "file"
	tests := []serverCapabilityFixture{
		{
			name: "unmarshals declaration registration options",
			input: `
			{
				"declarationProvider": {
					"workDoneProgress": true,
					"id": "test-register-id",
					"documentSelector": [
						{
							"language": "typescript",
							"scheme": "file"
						}
					]
				}
			}`,
			expected: &ServerCapabilities{
				DeclarationProvider: DeclarationRegistrationOptions{
					DeclarationOptions: DeclarationOptions{
						WorkDoneProgressOptions: WorkDoneProgressOptions{
							WorkDoneProgress: &workDoneProgress,
						},
					},
					StaticRegistrationOptions: StaticRegistrationOptions{
						ID: &testRegisterID,
					},
					TextDocumentRegistrationOptions: TextDocumentRegistrationOptions{
						DocumentSelector: &DocumentSelector{
							{
								Language: &language,
								Scheme:   &fileScheme,
							},
						},
					},
				},
			},
		},
		{
			name:  "unmarshals boolean",
			input: "{\"declarationProvider\":true}",
			expected: &ServerCapabilities{
				DeclarationProvider: true,
			},
		},
		{
			name:     "unmarshals missing value to nil",
			input:    "{}",
			expected: &ServerCapabilities{},
		},
	}

	testServerCapabilities(&s.Suite, tests)
}

func (s *LanguageFeaturesTestSuite) Test_unmarshal_definition_provider_capability() {
	workDoneProgress := true
	tests := []serverCapabilityFixture{
		{
			name: "unmarshals definition options",
			input: `
			{
				"definitionProvider": {
					"workDoneProgress": true
				}
			}`,
			expected: &ServerCapabilities{
				DefinitionProvider: DefinitionOptions{
					WorkDoneProgressOptions: WorkDoneProgressOptions{
						WorkDoneProgress: &workDoneProgress,
					},
				},
			},
		},
		{
			name:  "unmarshals boolean",
			input: "{\"definitionProvider\":true}",
			expected: &ServerCapabilities{
				DefinitionProvider: true,
			},
		},
		{
			name:     "unmarshals missing value to nil",
			input:    "{}",
			expected: &ServerCapabilities{},
		},
	}

	testServerCapabilities(&s.Suite, tests)
}

func (s *LanguageFeaturesTestSuite) Test_unmarshal_type_definition_provider_capability() {
	workDoneProgress := true
	testRegisterID := "test-register-id-type-def"
	language := "typescript"
	fileScheme := "file"
	tests := []serverCapabilityFixture{
		{
			name: "unmarshals type definition registration options",
			input: `
			{
				"typeDefinitionProvider": {
					"workDoneProgress": true,
					"id": "test-register-id-type-def",
					"documentSelector": [
						{
							"language": "typescript",
							"scheme": "file"
						}
					]
				}
			}`,
			expected: &ServerCapabilities{
				TypeDefinitionProvider: TypeDefinitionRegistrationOptions{
					TypeDefinitionOptions: TypeDefinitionOptions{
						WorkDoneProgressOptions: WorkDoneProgressOptions{
							WorkDoneProgress: &workDoneProgress,
						},
					},
					StaticRegistrationOptions: StaticRegistrationOptions{
						ID: &testRegisterID,
					},
					TextDocumentRegistrationOptions: TextDocumentRegistrationOptions{
						DocumentSelector: &DocumentSelector{
							{
								Language: &language,
								Scheme:   &fileScheme,
							},
						},
					},
				},
			},
		},
		{
			name:  "unmarshals boolean",
			input: "{\"typeDefinitionProvider\":true}",
			expected: &ServerCapabilities{
				TypeDefinitionProvider: true,
			},
		},
		{
			name:     "unmarshals missing value to nil",
			input:    "{}",
			expected: &ServerCapabilities{},
		},
	}

	testServerCapabilities(&s.Suite, tests)
}

func (s *LanguageFeaturesTestSuite) Test_unmarshal_implementation_provider_capability() {
	workDoneProgress := true
	testRegisterID := "test-register-id-impl"
	language := "typescript"
	fileScheme := "file"
	tests := []serverCapabilityFixture{
		{
			name: "unmarshals implementation registration options",
			input: `
			{
				"implementationProvider": {
					"workDoneProgress": true,
					"id": "test-register-id-impl",
					"documentSelector": [
						{
							"language": "typescript",
							"scheme": "file"
						}
					]
				}
			}`,
			expected: &ServerCapabilities{
				ImplementationProvider: ImplementationRegistrationOptions{
					ImplementationOptions: ImplementationOptions{
						WorkDoneProgressOptions: WorkDoneProgressOptions{
							WorkDoneProgress: &workDoneProgress,
						},
					},
					StaticRegistrationOptions: StaticRegistrationOptions{
						ID: &testRegisterID,
					},
					TextDocumentRegistrationOptions: TextDocumentRegistrationOptions{
						DocumentSelector: &DocumentSelector{
							{
								Language: &language,
								Scheme:   &fileScheme,
							},
						},
					},
				},
			},
		},
		{
			name:  "unmarshals boolean",
			input: "{\"implementationProvider\":true}",
			expected: &ServerCapabilities{
				ImplementationProvider: true,
			},
		},
		{
			name:     "unmarshals missing value to nil",
			input:    "{}",
			expected: &ServerCapabilities{},
		},
	}

	testServerCapabilities(&s.Suite, tests)
}

func (s *LanguageFeaturesTestSuite) Test_unmarshal_references_provider_capability() {
	workDoneProgress := true
	tests := []serverCapabilityFixture{
		{
			name: "unmarshals reference options",
			input: `
			{
				"referencesProvider": {
					"workDoneProgress": true
				}
			}`,
			expected: &ServerCapabilities{
				ReferencesProvider: ReferenceOptions{
					WorkDoneProgressOptions: WorkDoneProgressOptions{
						WorkDoneProgress: &workDoneProgress,
					},
				},
			},
		},
		{
			name:  "unmarshals boolean",
			input: "{\"referencesProvider\":true}",
			expected: &ServerCapabilities{
				ReferencesProvider: true,
			},
		},
		{
			name:     "unmarshals missing value to nil",
			input:    "{}",
			expected: &ServerCapabilities{},
		},
	}

	testServerCapabilities(&s.Suite, tests)
}

func (s *LanguageFeaturesTestSuite) Test_unmarshal_document_highlight_provider_capability() {
	workDoneProgress := true
	tests := []serverCapabilityFixture{
		{
			name: "unmarshals document highlight options",
			input: `
			{
				"documentHighlightProvider": {
					"workDoneProgress": true
				}
			}`,
			expected: &ServerCapabilities{
				DocumentHighlightProvider: DocumentHighlightOptions{
					WorkDoneProgressOptions: WorkDoneProgressOptions{
						WorkDoneProgress: &workDoneProgress,
					},
				},
			},
		},
		{
			name:  "unmarshals boolean",
			input: "{\"documentHighlightProvider\":true}",
			expected: &ServerCapabilities{
				DocumentHighlightProvider: true,
			},
		},
		{
			name:     "unmarshals missing value to nil",
			input:    "{}",
			expected: &ServerCapabilities{},
		},
	}

	testServerCapabilities(&s.Suite, tests)
}

func (s *LanguageFeaturesTestSuite) Test_unmarshal_document_symbol_provider_capability() {
	workDoneProgress := true
	label := "Test Label"
	tests := []serverCapabilityFixture{
		{
			name: "unmarshals document symbol options",
			input: `
			{
				"documentSymbolProvider": {
					"workDoneProgress": true,
					"label": "Test Label"
				}
			}`,
			expected: &ServerCapabilities{
				DocumentSymbolProvider: DocumentSymbolOptions{
					WorkDoneProgressOptions: WorkDoneProgressOptions{
						WorkDoneProgress: &workDoneProgress,
					},
					Label: &label,
				},
			},
		},
		{
			name:  "unmarshals boolean",
			input: "{\"documentSymbolProvider\":true}",
			expected: &ServerCapabilities{
				DocumentSymbolProvider: true,
			},
		},
		{
			name:     "unmarshals missing value to nil",
			input:    "{}",
			expected: &ServerCapabilities{},
		},
	}

	testServerCapabilities(&s.Suite, tests)
}

func (s *LanguageFeaturesTestSuite) Test_unmarshal_code_action_provider_capability() {
	workDoneProgress := true
	resolveProvider := false
	tests := []serverCapabilityFixture{
		{
			name: "unmarshals code action options",
			input: `
			{
				"codeActionProvider": {
					"workDoneProgress": true,
					"codeActionKinds": ["quickfix", "refactor", "refactor.extract"],
					"resolveProvider": false
				}
			}`,
			expected: &ServerCapabilities{
				CodeActionProvider: CodeActionOptions{
					WorkDoneProgressOptions: WorkDoneProgressOptions{
						WorkDoneProgress: &workDoneProgress,
					},
					CodeActionKinds: []CodeActionKind{
						CodeActionKindQuickFix,
						CodeActionKindRefactor,
						CodeActionKindRefactorExtract,
					},
					ResolveProvider: &resolveProvider,
				},
			},
		},
		{
			name:  "unmarshals boolean",
			input: "{\"codeActionProvider\":true}",
			expected: &ServerCapabilities{
				CodeActionProvider: true,
			},
		},
		{
			name:     "unmarshals missing value to nil",
			input:    "{}",
			expected: &ServerCapabilities{},
		},
	}

	testServerCapabilities(&s.Suite, tests)
}

func (s *LanguageFeaturesTestSuite) Test_unmarshal_color_provider_capability() {
	workDoneProgress := true
	testRegisterID := "test-register-id-color"
	language := "typescript"
	fileScheme := "file"
	tests := []serverCapabilityFixture{
		{
			name: "unmarshals document color options",
			input: `
			{
				"colorProvider": {
					"workDoneProgress": true,
					"id": "test-register-id-color",
					"documentSelector": [
						{
							"language": "typescript",
							"scheme": "file"
						}
					]
				}
			}`,
			expected: &ServerCapabilities{
				ColorProvider: DocumentColorRegistrationOptions{
					DocumentColorOptions: DocumentColorOptions{
						WorkDoneProgressOptions: WorkDoneProgressOptions{
							WorkDoneProgress: &workDoneProgress,
						},
					},
					StaticRegistrationOptions: StaticRegistrationOptions{
						ID: &testRegisterID,
					},
					TextDocumentRegistrationOptions: TextDocumentRegistrationOptions{
						DocumentSelector: &DocumentSelector{
							{
								Language: &language,
								Scheme:   &fileScheme,
							},
						},
					},
				},
			},
		},
		{
			name:  "unmarshals boolean",
			input: "{\"colorProvider\":true}",
			expected: &ServerCapabilities{
				ColorProvider: true,
			},
		},
		{
			name:     "unmarshals missing value to nil",
			input:    "{}",
			expected: &ServerCapabilities{},
		},
	}

	testServerCapabilities(&s.Suite, tests)
}

func (s *LanguageFeaturesTestSuite) Test_unmarshal_document_formatting_provider_capability() {
	workDoneProgress := true
	tests := []serverCapabilityFixture{
		{
			name: "unmarshals document formatting options",
			input: `
			{
				"documentFormattingProvider": {
					"workDoneProgress": true
				}
			}`,
			expected: &ServerCapabilities{
				DocumentFormattingProvider: DocumentFormattingOptions{
					WorkDoneProgressOptions: WorkDoneProgressOptions{
						WorkDoneProgress: &workDoneProgress,
					},
				},
			},
		},
		{
			name:  "unmarshals boolean",
			input: "{\"documentFormattingProvider\":true}",
			expected: &ServerCapabilities{
				DocumentFormattingProvider: true,
			},
		},
		{
			name:     "unmarshals missing value to nil",
			input:    "{}",
			expected: &ServerCapabilities{},
		},
	}

	testServerCapabilities(&s.Suite, tests)
}

func (s *LanguageFeaturesTestSuite) Test_unmarshal_document_range_formatting_provider_capability() {
	workDoneProgress := true
	tests := []serverCapabilityFixture{
		{
			name: "unmarshals document range formatting options",
			input: `
			{
				"documentRangeFormattingProvider": {
					"workDoneProgress": true
				}
			}`,
			expected: &ServerCapabilities{
				DocumentRangeFormattingProvider: DocumentRangeFormattingOptions{
					WorkDoneProgressOptions: WorkDoneProgressOptions{
						WorkDoneProgress: &workDoneProgress,
					},
				},
			},
		},
		{
			name:  "unmarshals boolean",
			input: "{\"documentRangeFormattingProvider\":true}",
			expected: &ServerCapabilities{
				DocumentRangeFormattingProvider: true,
			},
		},
		{
			name:     "unmarshals missing value to nil",
			input:    "{}",
			expected: &ServerCapabilities{},
		},
	}

	testServerCapabilities(&s.Suite, tests)
}

func (s *LanguageFeaturesTestSuite) Test_unmarshal_rename_provider_capability() {
	workDoneProgress := true
	prepareProvider := true
	tests := []serverCapabilityFixture{
		{
			name: "unmarshals rename options",
			input: `
			{
				"renameProvider": {
					"workDoneProgress": true,
					"prepareProvider": true
				}
			}`,
			expected: &ServerCapabilities{
				RenameProvider: RenameOptions{
					WorkDoneProgressOptions: WorkDoneProgressOptions{
						WorkDoneProgress: &workDoneProgress,
					},
					PrepareProvider: &prepareProvider,
				},
			},
		},
		{
			name:  "unmarshals boolean",
			input: "{\"renameProvider\":true}",
			expected: &ServerCapabilities{
				RenameProvider: true,
			},
		},
		{
			name:     "unmarshals missing value to nil",
			input:    "{}",
			expected: &ServerCapabilities{},
		},
	}

	testServerCapabilities(&s.Suite, tests)
}

func TestLanguageFeaturesTestSuite(t *testing.T) {
	suite.Run(t, new(LanguageFeaturesTestSuite))
}
