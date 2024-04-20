provider "google" {
  project     = var.project_id
  credentials = file("../.credentials/service_account.json")
  region      = var.region
  zone        = var.zone
}

module "run" {
  source = "./modules/run"
}

#module "http_handler" {
#  source = "modules/functions/http_function"
#}
