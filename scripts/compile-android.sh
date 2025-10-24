#!/bin/sh

cd "$(dirname "$0")/.."
if [ ! -d "dist/" ]; then
    mkdir dist/
fi

fyne package -os android -app-id me.islekcaganmert.experiment.projectile_to_height -icon icon.png -name "Projectile to Height Converter" -release
mv Projectile_to_Height_Converter.apk dist/
