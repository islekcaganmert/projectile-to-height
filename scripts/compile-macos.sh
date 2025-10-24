#!/bin/sh

cd "$(dirname "$0")/.."

if [ -d "build/" ]; then
    rm -rf build/
fi
mkdir build/

mkdir build/Projectile\ to\ Height\ Converter.app/
mkdir build/Projectile\ to\ Height\ Converter.app/Contents/
mkdir build/Projectile\ to\ Height\ Converter.app/Contents/MacOS/

GOOS=darwin GOARCH=arm64 go build -o build/Projectile\ to\ Height\ Converter.app/Contents/MacOS/Projectile\ to\ Height\ Converter

if [ ! -z build/Projectile\ to\ Height\ Converter.app/Contents/MacOS/Projectile\ to\ Height\ Converter ]; then
    echo "Build succeeded."
else
    echo "Build failed."
    exit 1
fi

cp configs/Info.plist build/Projectile\ to\ Height\ Converter.app/Contents/Info.plist

if [ ! -d "dist/" ]; then
    mkdir dist/
fi

tar -cvzf dist/Projectile_to_Height_Converter.tar.gz -C build Projectile\ to\ Height\ Converter.app

rm -rf build/