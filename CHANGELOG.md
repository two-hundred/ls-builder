# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.2.3] - 2024-09-14

### Fixed

- Adds previously missing child context utilising the timeout configured with the server for every request received from the client to the server.

## [0.2.2] - 2024-09-14

### Fixed

- Removes confusing timeouts that cancel early and create bugs for a language server when making calls to the client. `streamTimeout` and `websocketTimeout` are no longer supported and have been removed as they never worked as intended.

## [0.2.1] - 2024-06-26

### Fixed

- Corrects `Dispatcher.WorkspaceConfiguration` to accept a user-defined target to unmarshal the configuration result from the client into.

## [0.2.0] - 2024-06-26

### Added

- Convenience trace service for storing shared trace configuration and sending log messages to clients.

## [0.1.0] - 2024-06-25

### Added

- Initial implementation of the toolkit for building language servers compatible with [3.17.0](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/) of the Language Server Protocol.
