#!/bin/bash

BINARY_NAME="./bin/todo"

DEST_DIR="/usr/local/bin"

sudo cp "$BINARY_NAME" "$DEST_DIR"

echo "The todo binary has been installed to $DEST_DIR."
