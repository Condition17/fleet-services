terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
    }
  }
}

provider "google" {
  credentials = file("terraform_service_account.json")

  project = var.project_id
  region  = var.region
  zone    = var.zone
}

#  Ensure GCR registry is set up
resource "google_container_registry" "registry" {}

# Create CircleCI role

resource "google_project_iam_custom_role" "gcr_custom_admin" {
  project     = var.project_id
  role_id     = "gcrCustomAdmin"
  title       = "GCR custom admin"
  description = "Custom admin role for GCR. It should permit all the operations needed on containers."
  permissions = [
    "storage.buckets.create",
    "storage.buckets.delete",
    "storage.buckets.get",
    "storage.buckets.list",
    "storage.buckets.update",
    "storage.objects.create",
    "storage.objects.delete",
    "storage.objects.get",
    "storage.objects.list",
    s"storage.objects.update"
  ]
}

# Create CircleCI service account

resource "google_service_account" "circleci_service_account" {
  account_id   = "circleci"
  display_name = "Service account used by CircleCI"
}


resource "google_project_iam_policy" "project" {
  project     = var.project_id
  policy_data = data.google_iam_policy.admin.policy_data
}

data "google_iam_policy" "admin" {
  binding {
    role = "projects/fleet-271114/roles/gcrCustomAdmin"

    members = [
      "serviceAccount:${google_service_account.circleci_service_account.email}",
    ]
  }
}
