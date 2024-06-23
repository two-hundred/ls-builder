package lsp

import "errors"

// ErrorWithData is an error that contains additional data such as that for
// server cancellations.
type ErrorWithData struct {
	// The error code (e.g. diagnostic error code)
	Code *IntOrString
	// The data to send with the error.
	Data any
}

var (
	ErrInvalidDocumentDiagnosticReportKind = errors.New("invalid document diagnostic report kind")
)
