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

func (s *LanguageFeaturesTestSuite) Test_unmarshal_folding_range_provider_capability() {
	workDoneProgress := true
	testRegisterID := "test-register-id-folding-range"
	language := "typescript"
	fileScheme := "file"
	tests := []serverCapabilityFixture{
		{
			name: "unmarshals folding range registration options",
			input: `
			{
				"foldingRangeProvider": {
					"workDoneProgress": true,
					"id": "test-register-id-folding-range",
					"documentSelector": [
						{
							"language": "typescript",
							"scheme": "file"
						}
					]
				}
			}`,
			expected: &ServerCapabilities{
				FoldingRangeProvider: FoldingRangeRegistrationOptions{
					FoldingRangeOptions: FoldingRangeOptions{
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
			input: "{\"foldingRangeProvider\":true}",
			expected: &ServerCapabilities{
				FoldingRangeProvider: true,
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

func (s *LanguageFeaturesTestSuite) Test_unmarshal_selection_range_provider_capability() {
	workDoneProgress := true
	testRegisterID := "test-register-id-selection-range"
	language := "typescript"
	fileScheme := "file"
	tests := []serverCapabilityFixture{
		{
			name: "unmarshals selection range registration options",
			input: `
			{
				"selectionRangeProvider": {
					"workDoneProgress": true,
					"id": "test-register-id-selection-range",
					"documentSelector": [
						{
							"language": "typescript",
							"scheme": "file"
						}
					]
				}
			}`,
			expected: &ServerCapabilities{
				SelectionRangeProvider: SelectionRangeRegistrationOptions{
					SelectionRangeOptions: SelectionRangeOptions{
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
			input: "{\"selectionRangeProvider\":true}",
			expected: &ServerCapabilities{
				SelectionRangeProvider: true,
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

func (s *LanguageFeaturesTestSuite) Test_unmarshal_linked_editing_range_provider_capability() {
	workDoneProgress := true
	testRegisterID := "test-register-id-linked-editing-range"
	language := "typescript"
	fileScheme := "file"
	tests := []serverCapabilityFixture{
		{
			name: "unmarshals linked editing range registration options",
			input: `
			{
				"linkedEditingRangeProvider": {
					"workDoneProgress": true,
					"id": "test-register-id-linked-editing-range",
					"documentSelector": [
						{
							"language": "typescript",
							"scheme": "file"
						}
					]
				}
			}`,
			expected: &ServerCapabilities{
				LinkedEditingRangeProvider: LinkedEditingRangeRegistrationOptions{
					LinkedEditingRangeOptions: LinkedEditingRangeOptions{
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
			input: "{\"linkedEditingRangeProvider\":true}",
			expected: &ServerCapabilities{
				LinkedEditingRangeProvider: true,
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

func (s *LanguageFeaturesTestSuite) Test_unmarshal_call_hierarchy_provider_capability() {
	workDoneProgress := true
	testRegisterID := "test-register-id-call-hierarchy"
	language := "typescript"
	fileScheme := "file"
	tests := []serverCapabilityFixture{
		{
			name: "unmarshals call hierarchy registration options",
			input: `
			{
				"callHierarchyProvider": {
					"workDoneProgress": true,
					"id": "test-register-id-call-hierarchy",
					"documentSelector": [
						{
							"language": "typescript",
							"scheme": "file"
						}
					]
				}
			}`,
			expected: &ServerCapabilities{
				CallHierarchyProvider: CallHierarchyRegistrationOptions{
					CallHierarchyOptions: CallHierarchyOptions{
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
			input: "{\"callHierarchyProvider\":true}",
			expected: &ServerCapabilities{
				CallHierarchyProvider: true,
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

func (s *LanguageFeaturesTestSuite) Test_unmarshal_semantic_tokens_provider_capability() {
	workDoneProgress := true
	testRegisterID := "test-register-id-semantic-tokens"
	language := "typescript"
	fileScheme := "file"
	delta := true
	tests := []serverCapabilityFixture{
		{
			name: "unmarshals semantic token options with delta",
			input: `
			{
				"semanticTokensProvider": {
					"workDoneProgress": true,
					"id": "test-register-id-semantic-tokens",
					"documentSelector": [
						{
							"language": "typescript",
							"scheme": "file"
						}
					],
					"legend": {
						"tokenTypes": ["namespace", "type", "class", "enum", "interface", "struct", "typeParameter"],
						"tokenModifiers": ["declaration", "definition", "readonly", "static", "deprecated", "abstract", "async"]
					},
					"range": true,
					"full": {
						"delta": true
					}
				}
			}`,
			expected: &ServerCapabilities{
				SemanticTokensProvider: SemanticTokensRegistrationOptions{
					SemanticTokensOptions: SemanticTokensOptions{
						WorkDoneProgressOptions: WorkDoneProgressOptions{
							WorkDoneProgress: &workDoneProgress,
						},
						Legend: SemanticTokensLegend{
							TokenTypes:     []string{"namespace", "type", "class", "enum", "interface", "struct", "typeParameter"},
							TokenModifiers: []string{"declaration", "definition", "readonly", "static", "deprecated", "abstract", "async"},
						},
						Range: true,
						Full: SemanticDelta{
							Delta: &delta,
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
			name: "unmarshals semantic token options without delta",
			input: `
			{
				"semanticTokensProvider": {
					"workDoneProgress": true,
					"id": "test-register-id-semantic-tokens",
					"documentSelector": [
						{
							"language": "typescript",
							"scheme": "file"
						}
					],
					"legend": {
						"tokenTypes": ["namespace", "type", "class", "enum", "interface", "struct", "typeParameter"],
						"tokenModifiers": ["declaration", "definition", "readonly", "static", "deprecated", "abstract", "async"]
					},
					"range": {},
					"full": true
				}
			}`,
			expected: &ServerCapabilities{
				SemanticTokensProvider: SemanticTokensRegistrationOptions{
					SemanticTokensOptions: SemanticTokensOptions{
						WorkDoneProgressOptions: WorkDoneProgressOptions{
							WorkDoneProgress: &workDoneProgress,
						},
						Legend: SemanticTokensLegend{
							TokenTypes:     []string{"namespace", "type", "class", "enum", "interface", "struct", "typeParameter"},
							TokenModifiers: []string{"declaration", "definition", "readonly", "static", "deprecated", "abstract", "async"},
						},
						Range: struct{}{},
						Full:  true,
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
			name:     "unmarshals missing value to nil",
			input:    "{}",
			expected: &ServerCapabilities{},
		},
	}

	testServerCapabilities(&s.Suite, tests)
}

func (s *LanguageFeaturesTestSuite) Test_unmarshal_moniker_provider_capability() {
	workDoneProgress := true
	testRegisterID := "test-register-id-moniker"
	language := "typescript"
	fileScheme := "file"
	tests := []serverCapabilityFixture{
		{
			name: "unmarshals moniker registration options",
			input: `
			{
				"monikerProvider": {
					"workDoneProgress": true,
					"id": "test-register-id-moniker",
					"documentSelector": [
						{
							"language": "typescript",
							"scheme": "file"
						}
					]
				}
			}`,
			expected: &ServerCapabilities{
				MonikerProvider: MonikerRegistrationOptions{
					MonikerOptions: MonikerOptions{
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
			input: "{\"monikerProvider\":true}",
			expected: &ServerCapabilities{
				MonikerProvider: true,
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

func (s *LanguageFeaturesTestSuite) Test_unmarshal_type_hierarchy_provider_capability() {
	workDoneProgress := true
	testRegisterID := "test-register-id-type-hierarchy"
	language := "typescript"
	fileScheme := "file"
	tests := []serverCapabilityFixture{
		{
			name: "unmarshals type hierarchy registration options",
			input: `
			{
				"typeHierarchyProvider": {
					"workDoneProgress": true,
					"id": "test-register-id-type-hierarchy",
					"documentSelector": [
						{
							"language": "typescript",
							"scheme": "file"
						}
					]
				}
			}`,
			expected: &ServerCapabilities{
				TypeHierarchyProvider: TypeHierarchyRegistrationOptions{
					TypeHierarchyOptions: TypeHierarchyOptions{
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
			input: "{\"typeHierarchyProvider\":true}",
			expected: &ServerCapabilities{
				TypeHierarchyProvider: true,
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

func (s *LanguageFeaturesTestSuite) Test_unmarshal_inline_value_provider_capability() {
	workDoneProgress := true
	testRegisterID := "test-register-id-inline-value"
	language := "typescript"
	fileScheme := "file"
	tests := []serverCapabilityFixture{
		{
			name: "unmarshals inline value registration options",
			input: `
			{
				"inlineValueProvider": {
					"workDoneProgress": true,
					"id": "test-register-id-inline-value",
					"documentSelector": [
						{
							"language": "typescript",
							"scheme": "file"
						}
					]
				}
			}`,
			expected: &ServerCapabilities{
				InlineValueProvider: InlineValueRegistrationOptions{
					InlineValueOptions: InlineValueOptions{
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
			input: "{\"inlineValueProvider\":true}",
			expected: &ServerCapabilities{
				InlineValueProvider: true,
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

func (s *LanguageFeaturesTestSuite) Test_unmarshal_inlay_hint_provider_capability() {
	workDoneProgress := true
	testRegisterID := "test-register-id-inlay-hint"
	language := "typescript"
	fileScheme := "file"
	resolveProvider := true
	tests := []serverCapabilityFixture{
		{
			name: "unmarshals inlay hint registration options",
			input: `
			{
				"inlayHintProvider": {
					"workDoneProgress": true,
					"id": "test-register-id-inlay-hint",
					"documentSelector": [
						{
							"language": "typescript",
							"scheme": "file"
						}
					],
					"resolveProvider": true
				}
			}`,
			expected: &ServerCapabilities{
				InlayHintProvider: InlayHintRegistrationOptions{
					InlayHintOptions: InlayHintOptions{
						WorkDoneProgressOptions: WorkDoneProgressOptions{
							WorkDoneProgress: &workDoneProgress,
						},
						ResolveProvider: &resolveProvider,
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
			input: "{\"inlayHintProvider\":true}",
			expected: &ServerCapabilities{
				InlayHintProvider: true,
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

func (s *LanguageFeaturesTestSuite) Test_unmarshal_diagnostic_provider_capability() {
	workDoneProgress := true
	testRegisterID := "test-register-id-diagnostic"
	language := "typescript"
	fileScheme := "file"
	identifier := "test-identifier"
	interFileDependencies := true
	workspaceDiagnostics := false
	tests := []serverCapabilityFixture{
		{
			name: "unmarshals diagnostic registration options",
			input: `
			{
				"diagnosticProvider": {
					"workDoneProgress": true,
					"id": "test-register-id-diagnostic",
					"documentSelector": [
						{
							"language": "typescript",
							"scheme": "file"
						}
					],
					"identifier": "test-identifier",
					"interFileDependencies": true,
					"workspaceDiagnostics": false
				}
			}`,
			expected: &ServerCapabilities{
				DiagnosticProvider: DiagnosticRegistrationOptions{
					DiagnosticOptions: DiagnosticOptions{
						WorkDoneProgressOptions: WorkDoneProgressOptions{
							WorkDoneProgress: &workDoneProgress,
						},
						Identifier:            &identifier,
						InterFileDependencies: interFileDependencies,
						WorkspaceDiagnostics:  workspaceDiagnostics,
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
			name:     "unmarshals missing value to nil",
			input:    "{}",
			expected: &ServerCapabilities{},
		},
	}

	testServerCapabilities(&s.Suite, tests)
}

func (s *LanguageFeaturesTestSuite) Test_unmarshal_workspace_symbol_provider_capability() {
	workDoneProgress := true
	resolveProvider := true
	tests := []serverCapabilityFixture{
		{
			name: "unmarshals workspace symbol options",
			input: `
			{
				"workspaceSymbolProvider": {
					"workDoneProgress": true,
					"resolveProvider": true
				}
			}`,
			expected: &ServerCapabilities{
				WorkspaceSymbolProvider: WorkspaceSymbolOptions{
					WorkDoneProgressOptions: WorkDoneProgressOptions{
						WorkDoneProgress: &workDoneProgress,
					},
					ResolveProvider: &resolveProvider,
				},
			},
		},
		{
			name:  "unmarshals boolean",
			input: "{\"workspaceSymbolProvider\":true}",
			expected: &ServerCapabilities{
				WorkspaceSymbolProvider: true,
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
