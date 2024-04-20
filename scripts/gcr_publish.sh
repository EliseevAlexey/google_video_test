#!/bin/zsh
set -e

IMAGE_NAME="run"
TAG="latest"
PATH_TO_DOCKERFILE="../services/run"

# vars
PROJECT_ID=$(yq e '.google.project_id' ../config/base.yaml)
IMAGE_FULL_NAME="gcr.io/${PROJECT_ID}/${IMAGE_NAME}:${TAG}"


gcloud auth configure-docker && \
docker build --platform linux/amd64 -t "${IMAGE_NAME}" "${PATH_TO_DOCKERFILE}" && \
docker tag "${IMAGE_NAME}" "${IMAGE_FULL_NAME}" && \
docker push "${IMAGE_FULL_NAME}"
