#!/usr/bin/env bash

if [ -f bin/linux-terminal-notifier ]; then
    rm bin/linux-terminal-notifier
fi

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o bin/linux-terminal-notifier cmd/linux-terminal-notifier/main.go
