#!/usr/bin/env bash
rm ./bin/main
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o bin/main src/app/entry.go