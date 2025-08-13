@echo off
echo Generating Go code from protobuf...

REM Create proto directory if it doesn't exist
if not exist "proto" mkdir proto

REM Generate Go code
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/user.proto

if %ERRORLEVEL% EQU 0 (
    echo Protobuf code generated successfully!
    echo Now you can run: go mod tidy
) else (
    echo Failed to generate protobuf code
    echo Make sure protoc and plugins are installed
)

pause

