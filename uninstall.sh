#!/bin/bash

BINARY_NAME="todo"

DEST_DIR="/usr/local/bin"

sudo rm "$DEST_DIR/$BINARY_NAME"

echo "The $BINARY_NAME binary has been removed from $DEST_DIR."
