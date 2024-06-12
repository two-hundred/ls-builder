package lsp

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/two-hundred/ls-builder/common"
)

type Method = string

// Integer defins an integer number in the range
// -2^31 to 2^31-1.
type Integer = int32

// UInteger defines an unsigned integer number in the range
// 0 to 2^32-1.
type UInteger = uint32

// Decimal defins a decimal number. Since decimal numbers are very
// rare in the language server specification we denote the
// exact range with every decimal using the mathematics
// interval notation (e.g. [0, 1] denotes all decimals d with
// 0 <= d <= 1.
type Decimal = float64

// LSPAny is the LSP "any" type
// @since 3.17.0
//
// At runtime, type assertions are used to determine the actual
// set of allowed types.
// (LSPObject | LSPArray | string | integer | uinteger | decimal | boolean | null)
type LSPAny = interface{}

// LSPObject is the LSP "object" type
// @since 3.17.0
type LSPObject = map[string]LSPAny

// IntOrString represents a value that can be either an
// integer or string in LSP messages.
// integer | string (used in Diagnostic for example)
type IntOrString struct {
	IntVal *Integer
	StrVal *string
}

// Fulfils the json.Marshaler interface.
func (i *IntOrString) MarshalJSON() ([]byte, error) {
	if i.IntVal != nil {
		return json.Marshal(i.IntVal)
	}

	if i.StrVal != nil {
		return json.Marshal(i.StrVal)
	}

	return nil, errors.New("IntOrString must have either IntVal or StrVal")
}

// Fulfils the json.Unmarshaler interface.
func (i *IntOrString) UnmarshalJSON(data []byte) error {
	var intVal Integer
	if err := json.Unmarshal(data, &intVal); err == nil {
		i.IntVal = &intVal
		return nil
	}

	var strVal string
	err := json.Unmarshal(data, &strVal)
	if err != nil {
		return fmt.Errorf("expected an integer or string value, got error: %s", err.Error())

	}

	i.StrVal = &strVal
	return nil
}

// BoolOrString represents a value that can be either a boolean or a string
// in LSP messages.
// boolean | string
// (used in WorkspaceFoldersServerCapabilities for example)
type BoolOrString struct {
	BoolVal *bool
	StrVal  *string
}

// Fulfils the json.Marshaler interface.
func (b *BoolOrString) MarshalJSON() ([]byte, error) {
	if b.BoolVal != nil {
		return json.Marshal(b.BoolVal)
	}

	if b.StrVal != nil {
		return json.Marshal(b.StrVal)
	}

	return nil, errors.New("BoolOrString must have either BoolVal or StrVal")
}

// Fulfils the json.Unmarshaler interface
func (b *BoolOrString) UnmarshalJSON(data []byte) error {
	var boolVal bool
	if err := json.Unmarshal(data, &boolVal); err == nil {
		b.BoolVal = &boolVal
		return nil
	}

	var strVal string
	err := json.Unmarshal(data, &strVal)
	if err != nil {
		return fmt.Errorf("expected a boolean or string value, got error: %s", err.Error())
	}

	b.StrVal = &strVal
	return nil
}

// Fulfils the fmt.Stringer interface.
func (b *BoolOrString) String() string {
	if b.BoolVal != nil {
		return strconv.FormatBool(*b.BoolVal)
	} else {
		return *b.StrVal
	}
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#cancelRequest

const MethodCancelRequest = Method("$/cancelRequest")

// CancelParams contains the cancel request parameters.
type CancelParams struct {
	// ID of the request to cancel.
	ID IntOrString `json:"id"`
}

// CancelRequestFunc is the function signature for the cancelRequest request
// that can be registered for a language server.
type CancelRequestFunc func(ctx *common.LSPContext, params *CancelParams) error

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#progress

const MethodProgress = Method("$/progress")

// ProgressParams contains the progress notification parameters.
type ProgressParams struct {
	// The progress token provided by the client or server.
	Token ProgressToken `json:"token"`
	// The progress data.
	Value LSPAny `json:"value"`
}

// ProgressToken is either an integer or a string
// which is used to report progress out of band and notifications.
// The token is not the same as the request ID.
type ProgressToken = IntOrString
