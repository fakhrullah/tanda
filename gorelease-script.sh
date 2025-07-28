#!/usr/bin/env bash
# Define the path to your .env file
ENV_FILE=".env"

# Check if the .env file exists
if [ -f "$ENV_FILE" ]; then
    echo "Loading environment variables from $ENV_FILE..."
    # set -a: Automatically exports all subsequent variables
    set -a
    # source: Reads and executes commands from the file in the current shell context
    . "$ENV_FILE"
    # set +a: Turns off auto-exporting
    set +a
    echo "Environment variables loaded."
else
    echo "Warning: .env file not found at $ENV_FILE. Proceeding without it." >&2
fi

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