#!/bin/bash

set -e

pushd "$(dirname "$0")/../.." > /dev/null

SYSTEM="$(uname)"
USER_ID=$(id -u $USER)

if [ $USER_ID == 0 ]; then
    docker build -f docker/app/Dockerfile-root -t likechain/likechain .
else
    if [ $SYSTEM == "Darwin" ] || [ $USER_ID == 0 ]; then
        BUILD_UID=1000
        BUILD_GID=1000
    else
        BUILD_UID=$(id -u $USER)
        BUILD_GID=$(id -g $USER)
    fi
    docker build -f docker/app/Dockerfile -t likechain/likechain \
        --build-arg UID=$BUILD_UID --build-arg GID=$BUILD_GID .
fi
popd > /dev/null
