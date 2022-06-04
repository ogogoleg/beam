resource "google_project_service" "firestore" {
  project = var.project_id
  service = "firestore.googleapis.com"

  disable_dependent_services = true
}

resource "google_app_engine_application" "app_playground" {
  project     = var.project_id
  location_id = var.location
  database_type = "CLOUD_FIRESTORE"
}