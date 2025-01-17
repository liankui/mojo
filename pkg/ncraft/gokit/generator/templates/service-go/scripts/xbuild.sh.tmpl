#!/bin/bash
set -e

case "$1" in
amd64 | arm64) ARCH="$1" ;;
"") ARCH="amd64" ;;
*) echo "illegal option: $1" && exit 1 ;;
esac

CURRENT_DIR=$(cd "$(dirname "$0")" && pwd)
SERVICE_DIR=$(cd "$CURRENT_DIR"/.. && pwd)
CMD_DIR="$SERVICE_DIR"/cmd
REMOTE_PREFIX=harbor/demo
TAG=latest

(cd "$CMD_DIR"/demo-server && ./xbuild.sh linux "$ARCH") &
(cd "$CMD_DIR"/agent-server && ./xbuild.sh linux "$ARCH")

gitHash=$(git rev-list -1 HEAD)
gitBranch=$(git branch --show-current)

SERVERS=(demo-server agent-server)
cd "$SERVICE_DIR"/build/demo-server-docker
for server in "${SERVERS[@]}"; do
  echo "docker build $server"
  cd ../"$server"-docker
  (
    set -x
    docker build -f Dockerfile --platform linux/"$ARCH" --build-arg GIT_BRANCH="$gitBranch" --build-arg GIT_HASH="$gitHash" -t "$REMOTE_PREFIX"/demo-"$server":"$TAG" "$SERVICE_DIR"
    docker push "$REMOTE_PREFIX"/demo-"$server":"$TAG"
  )
done
