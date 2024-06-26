# ls-builder - Language Server Protocol 3.17.0

```go
package main

import (
    lsp "github.com/two-hundred/ls-builder/lsp_3_17"
)
```

This package provides a toolkit for building language servers compatible with [3.17.0](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/) of the Language Server Protocol.

The API primarily consists of `Handler` which can be used to configure capabilities and handlers for client to server requests and notifications and `Dispatcher` for server to client requests and notifications.
