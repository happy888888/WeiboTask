#!/bin/sh
name="WeiboTask"

export CGO_ENABLED=0
export GOPATH=$(pwd)

export GOOS="windows"
export GOARCH="386"
go install -ldflags "-s -w" $name

GOOS="linux"
GOARCH="amd64"
go install -ldflags "-s -w" $name

GOARCH="amd64"
go install -ldflags "-s -w" $name

GOARCH="arm"
go install -ldflags "-s -w" $name

GOARCH="arm64"
go install -ldflags "-s -w" $name

export GOMIPS="softfloat"
GOARCH="mips"
go install -ldflags "-s -w" $name

GOARCH="mipsle"
go install -ldflags "-s -w" $name

mkdir ./release
zip -q -j ./release/$name-windows_86-v$version.zip ./bin/windows_386/$name.exe ./config.json
tar -cvf ./release/$name-linux_64-v$version.tar --transform s=./bin/== ./bin/$name config.json
tar -cvf ./release/$name-linux_arm-v$version.tar --transform s=./bin/linux_arm/== ./bin/linux_arm/$name config.json
tar -cvf ./release/$name-linux_arm64-v$version.tar --transform s=./bin/linux_arm64/== ./bin/linux_arm64/$name config.json
tar -cvf ./release/$name-linux_mips-v$version.tar --transform s=./bin/linux_mips/== ./bin/linux_mips/$name config.json
tar -cvf ./release/$name-linux_mipsle-v$version.tar --transform s=./bin/linux_mipsle/== ./bin/linux_mipsle/$name config.json
