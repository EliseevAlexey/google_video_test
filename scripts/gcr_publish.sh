#!/bin/zsh
set -e

IMAGE_NAME="run"
TAG="latest"


# vars
PROJECT_ID=$(yq e '.google.project_id' ../config/base.yaml)
IMAGE_FULL_NAME="gcr.io/${PROJECT_ID}/${IMAGE_NAME}:${TAG}"


cd ../run
gcloud auth configure-docker
docker build --platform linux/amd64 -t "${IMAGE_NAME}" .
docker tag "${IMAGE_NAME}" "${IMAGE_FULL_NAME}"
docker push "${IMAGE_FULL_NAME}"
