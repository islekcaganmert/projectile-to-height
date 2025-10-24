#!/bin/sh

cd "$(dirname "$0")/.."
if [ ! -d "dist/" ]; then
    mkdir dist/
fi

fyne-cross windows -app-id me.islekcaganmert.experiment.projectile_to_height -arch amd64
fyne-cross windows -app-id me.islekcaganmert.experiment.projectile_to_height -arch arm64
mv fyne-cross/bin/windows-amd64/* dist/projectile_to_height-amd64.exe
mv fyne-cross/bin/windows-arm64/* dist/projectile_to_height-arm64.exe
rm -rf fyne-cross