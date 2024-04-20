#!/bin/zsh
set -e


# vars
PROJECT_ID=$(yq e '.google.project_id' ../config/base.yaml)
SERVICE_ACCOUNT_NAME=$(yq e '.google.service_account_name' ../config/base.yaml)
SERVICE_ACCOUNT_EMAIL="${SERVICE_ACCOUNT_NAME}@${PROJECT_ID}.iam.gserviceaccount.com"


# Login
gcloud auth login


# Enable services
gcloud services enable \
         cloudbuild.googleapis.com \
            compute.googleapis.com \
  containerregistry.googleapis.com \
                run.googleapis.com \
      secretmanager.googleapis.com


# Create a Service Account if needed
if gcloud iam service-accounts list --format="value(email)" | grep -q "^${SERVICE_ACCOUNT_EMAIL}$"; then
  echo "Service account with email '${SERVICE_ACCOUNT_EMAIL}' exists."
else
  gcloud iam \
    service-accounts create "${SERVICE_ACCOUNT_NAME}" \
    --display-name "${SERVICE_ACCOUNT_NAME}"
fi


# Grand Permissions to Service Account
gcloud projects \
  add-iam-policy-binding "${PROJECT_ID}" \
  --member serviceAccount:"${SERVICE_ACCOUNT_EMAIL}" \
  --role roles/owner


# Create and download Service Account Key
gcloud iam \
  service-accounts keys create ../.credentials/service_account.json \
  --iam-account="${SERVICE_ACCOUNT_EMAIL}"


# Terraform create environment
cd ../terraform
terraform init
terraform apply -auto-approve
