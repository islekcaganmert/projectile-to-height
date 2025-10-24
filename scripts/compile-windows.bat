@echo off
cd /d "%~dp0\.."

if not exist "dist\" (
    mkdir dist\
)

go build -o dist\projectile-to-height-x64.exe
go build -o dist\projectile-to-height-arm64.exe