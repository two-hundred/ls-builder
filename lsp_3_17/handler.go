package lsp

type Handler struct {
	// Base Protocol
	CancelRequest CancelRequestFunc
}
