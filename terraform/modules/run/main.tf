resource "google_cloud_run_service" "run-service" {
  name     = "${var.project_name}-run-service"
  location = var.region

  template {
    spec {
      service_account_name = var.service_account_email
      containers {
        image = "gcr.io/${var.project_id}/run:latest"
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}

data "google_iam_policy" "noauth" {
  binding {
    role    = "roles/run.invoker"
    members = [
      "allUsers",
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "noauth" {
  location = google_cloud_run_service.run-service.location
  project  = google_cloud_run_service.run-service.project
  service  = google_cloud_run_service.run-service.name

  policy_data = data.google_iam_policy.noauth.policy_data
}
