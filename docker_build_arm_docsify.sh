#!/bin/bash
docker buildx create --use --name multi-arch-builder
docker buildx inspect --bootstrap
docker buildx build --platform linux/amd64,linux/arm64 --push -t hunterhug/algorithm:docsify -f Dockerfile_docsify .