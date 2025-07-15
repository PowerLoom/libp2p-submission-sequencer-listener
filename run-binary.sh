#!/bin/bash

# Build the application
go build -o ./cmd/main ./cmd

# Load environment variables from .env, excluding comments and empty lines
# and run the application
env $(grep -v -e '^#' -e '^$' .env | xargs) ./cmd/main "$@"
