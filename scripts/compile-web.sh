#!/bin/sh

cd "$(dirname "$0")/.."
if [ ! -d "dist/" ]; then
    mkdir dist/
fi

fyne package -os web
mv wasm dist/web