variable "project_id" {
  default = "primeval-wind-420823"
}

variable "project_name" {
  default = "terraform"
}

variable "service_account_email" {
  default = "terraform-sa@primeval-wind-420823.iam.gserviceaccount.com"
}

variable "region" {
  default = "us-west2"
}

variable "image" {
  default = "us-west2-docker.pkg.dev/primeval-wind-420823/run:latest"
}
