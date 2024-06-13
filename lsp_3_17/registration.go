package lsp

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#client_registerCapability

// ClientRegisterCapability is the method that can be used by a server
// to register a new capability with the client.
const ClientRegisterCapability = Method("client/registerCapability")

// Registration provides general parameters to register
// for a capability.
type Registration struct {
	// The id used to register the request. The ID can be used to deregister
	// the request in the future.
	ID string `json:"id"`

	// The method / capability to register for.
	Method string `json:"method"`

	// Options necessary for the registration.
	RegisterOptions LSPAny `json:"registerOptions,omitempty"`
}

// RegistrationParams contains a list of registrations.
type RegistrationParams struct {
	Registrations []Registration `json:"registrations"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#client_unregisterCapability

// ClientUnregisterCapability is the method that can be used by a server
// to de-register a capability with the client.
const ClientUnregisterCapability = Method("client/unregisterCapability")

// Unregistration provides general parameters to unregister
// a capability.
type Unregistration struct {
	// The id used to unregister the request or notification. Usually an id
	// provided during the register request.
	ID string `json:"id"`

	// The method / capability to unregister for.
	Method string `json:"method"`
}

// UnregistrationParams contains a list of unregistrations.
type UnregistrationParams struct {
	// This should correctly be named `unregistrations`.
	// However, changing this is a breaking change and needs to wait
	// util 4.x version of the specification is delivered.
	Unregistrations []Unregistration `json:"unregisterations"`
}

type StaticRegistrationOptions struct {
	// The id used to register the request. The ID can be used to deregister
	// the request in the future. Tou can also see Registration#id.
	ID *string `json:"id,omitempty"`
}
