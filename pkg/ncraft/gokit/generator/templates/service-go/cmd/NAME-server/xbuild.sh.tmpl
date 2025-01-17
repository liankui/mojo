#!/bin/bash
set -e

case "$1" in
linux | windows | darwin) OS="$1" ;;
mac) OS="darwin" ;;
"") OS="" ;;
*) echo "illegal option: $1" && exit 1 ;;
esac

case "$2" in
amd64 | arm64) ARCH="$2" ;;
"") ARCH="" ;;
*) echo "illegal option: $2" && exit 1 ;;
esac

gitHash=$(git rev-list -1 HEAD)
gitBranch=$(git branch --show-current)
buildTime=$(date "+%Y-%m-%d %H:%M:%S %Z")

set -x
env CGO_ENABLED=0 GOOS="$OS" GOARCH="$ARCH" go build -ldflags "-s -w -X main.gitHash=$gitHash -X main.gitBranch=$gitBranch -X 'main.buildTime=$buildTime'"
