resource "google_cloudfunctions_function" "http-function" {
  name                  = "http-function"
  service_account_name  = var.service_account_email
  trigger_http          = true
  source_archive_bucket = ""
  runtime               = "go122"
  entry_point           = "HelloWorld"
}
