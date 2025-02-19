#!/bin/bash
set -e

GOLANGCI_LINT_VERSION="v1.59.0"

setup_linters() {
    echo "Installing linters..."
	go install github.com/mgechev/revive@latest
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
		| sh -s -- -b "$(go env GOPATH)/bin" "$GOLANGCI_LINT_VERSION"
}

setup_utils() {
    echo "Installing utils..."
    go install mvdan.cc/gofumpt@latest
	  go install github.com/evilmartians/lefthook@latest
	  lefthook install
}

setup_linters
setup_utils
echo "Done!"
