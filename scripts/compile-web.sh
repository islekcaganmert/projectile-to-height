#!/bin/sh

cd "$(dirname "$0")/.."
if [ ! -d "dist/" ]; then
    mkdir dist/
fi

fyne package -os web

if [ -d "dist/web" ]; then
    rm -rf dist/web
fi

mkdir dist/web
cp -r wasm/*.wasm dist/web/
cp wasm/wasm_exec.js dist/web/
cp wasm/icon.png dist/web/
cp web/index.html dist/web/
cp web/manifest.json dist/web/
rm -rf wasm/