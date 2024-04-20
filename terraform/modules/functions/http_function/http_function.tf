resource "google_cloudfunctions_function" "function" {
  name                  = "http-function"
  service_account_name  = "${var.service_account_name}@${var.project_id}.iam.gserviceaccount.com"
  trigger_http          = true
  source_archive_bucket = ""
  runtime               = "go122"
  entry_point           = "HttpFunction"
}
