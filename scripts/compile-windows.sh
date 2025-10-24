#!/bin/sh

cd "$(dirname "$0")/.."
if [ ! -d "dist/" ]; then
    mkdir dist/
fi

GOOS=windows GOARCH=amd64 go build -o "dist/projectile-to-height-x64.exe"
GOOS=windows GOARCH=arm64 go build -o "dist/projectile-to-height-arm64.exe"