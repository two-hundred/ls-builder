# Contributing to the language server builder

## Getting set up

### Prerequisites

- [Go](https://golang.org/dl/) >=1.22.2
- [Node.js](https://nodejs.org/en/download/) >=18.17.0 (For pre-commit hooks)
- [Yarn](https://yarnpkg.com/getting-started/install) >=1.22.21 (For pre-commit hooks)

Dependencies are managed with Go modules (go.mod) and will be installed automatically when you first
run tests.

If you want to install dependencies manually you can run:

```bash
go mod download
```

#### Installing pre-commit hook dependencies

[Commitlint](https://commitlint.js.org/) is used to enforce commit message conventions. To set up the pre-commit hooks, you need to install the dependencies:

```bash
yarn
```

#### Installing pre-commit hooks

Ensure to make git use the custom directory for git hooks

```bash
git config core.hooksPath .githooks
```

## Running tests

```bash
bash ./scripts/run-tests.sh
```

## Releasing

To release a new version of the library, you need to create a new tag and push it to the repository.

The format must be `vX.Y.Z` where `X.Y.Z` is the semantic version number.


See [here](https://go.dev/wiki/Modules#publishing-a-release).

1. add a change log entry to the `CHANGELOG.md` file following the template below:

```markdown
## [0.2.0] - 2024-06-05

### Fixed:

- Corrects bug with sending notifications from the server.

### Added

- Add convenience functions to send specific requests and notifications to clients.
```

2. Create and push the new tag:

```bash
git tag -a v0.2.0 -m "chore: Release v0.2.0"
git push --tags
```

Be sure to add a release for the tag with notes following this template:

Title: `v0.2.0`

```markdown
## Fixed:

- Corrects bug with sending notifications from the server.

## Added

- Add convenience functions to send specific requests and notifications to clients.
```

3. Prompt Go to update its index of modules with the new release:

```bash
GOPROXY=proxy.golang.org go list -m github.com/two-hundred/ls-builder@v0.2.0
```

## Commit scope

**blueprint**

Example commit:

```bash
git commit -m 'fix: correct bug with sending notifications from the server'
```
