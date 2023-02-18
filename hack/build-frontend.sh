#!/usr/bin/env bash

echo hello!

docker buildx build ./frontend \
    -f ../hack/Dockerfile \
    --output ./dist/build-f \
    --target builder
