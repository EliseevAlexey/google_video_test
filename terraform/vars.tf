variable "project_id" {
  default = "primeval-wind-420823"
}

variable "project_name" {
  default = "terraform"
}

variable "service_account_name" {
  default = "terraform-sa"
}

variable "region" {
  default = "us-west2"
}

variable "zone" {
  default = "us-west2-b"
}

variable "os_image" {
  default = "cos-cloud/cos-stable"
}

variable "machine_type" {
  default = "e2-micro"
}
