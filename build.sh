#!/bin/bash
NAME="emq-auth-http"
MAIN_PATH="cmd/emq-auth-http/main.go"
ORG="lishimeng"

# shellcheck disable=SC2046
VERSION=$(git describe --tags $(git rev-list --tags --max-count=1))
# shellcheck disable=SC2154
COMMIT=$(git log --pretty=format:"%h" -1)
BUILD_TIME=$(date +%FT%T%z)

build_application(){
  git checkout "${VERSION}"
  docker build -t "${ORG}/${NAME}:${VERSION}" \
  --build-arg NAME="${NAME}" \
  --build-arg VERSION="${VERSION}" \
  --build-arg COMMIT="${COMMIT}" \
  --build-arg BUILD_TIME="${BUILD_TIME}" \
  --build-arg MAIN_PATH="${MAIN_PATH}" .
}

help_print(){
  echo "****************************************"
  echo "App:${NAME}"
  echo "Version:${VERSION}"
  echo "Commit:${COMMIT}"
  echo "BuildTime:${BUILD_TIME}"
  echo "Main_Path:${MAIN_PATH}"
  echo "****************************************"
  echo ""
}

help_print
build_application