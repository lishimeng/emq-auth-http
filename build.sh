#!/bin/bash
APP_NAME="emq-auth-http"
ORG="lishimeng"

# shellcheck disable=SC2046
TAG=$(git describe --tags $(git rev-list --tags --max-count=1))
# shellcheck disable=SC2154
COMMIT=${git log --pretty=format:"%h" -1}
BUILD_TIME=$(date +%FT%T%z)

build_application(){
  git checkout "${TAG}"
  docker build -t ${ORG}/${APP_NAME}:"$TAG" \
  --build-arg APP_VERSION="${TAG}" COMMIT="${COMMIT}" \
  BUILD_TIME="${BUILD_TIME}" .
}

help_print(){
  echo "build ${APP_NAME}"
  echo "Version:$TAG"
}

help_print
build_application