#!/usr/bin/env bash

ARG="$1"

# If no argument is supplied, show the list of options
if [ -z "$ARG" ]; then
    echo "Usage: $0 <command>"
    echo ""
    echo "Available commands:"
    echo "  ./gorelease-script healthcheck - Perform a dry run to check if it's safe to release."
    echo "  ./gorelease-script build      - Build without publishing to check for code errors."
    echo "  ./gorelease-script release    - Release the project (includes a healthcheck)."
    exit 1
fi

# Process the supplied argument
case "$ARG" in
    "healthcheck")
        echo "Running goreleaser healthcheck..."
        goreleaser healthcheck
        ;;
    "build")
        echo "Running goreleaser build..."
        goreleaser build --snapshot --clean
        ;;
    "release")
        echo "Running goreleaser healthcheck before release..."
        goreleaser healthcheck && \
        echo "Running goreleaser release..." && \
        goreleaser release
        ;;
    *)
        echo "Error: Unknown command '$ARG'."
        echo "Use '$0' without arguments to see available commands."
        exit 1
        ;;
esac