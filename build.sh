#!/usr/bin/bash
BUILD=$(git rev-parse --short HEAD)

echo Building $BUILD

go build -ldflags "-X 'roach/version.BuildNumber=${BUILD}'" .

mv roach ~/.local/bin
