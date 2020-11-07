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
    "storage.objects.update"
  ]
}

# Create CircleCI service account
resource "google_service_account" "circleci_service_account" {
  account_id   = "circleci"
  display_name = "Service account used by CircleCI"
}

resource "google_project_iam_member" "project" {
  project = var.project_id
  role    = "projects/${var.project_id}/roles/${google_project_iam_custom_role.gcr_custom_admin.role_id}"
  member  = "serviceAccount:${google_service_account.circleci_service_account.email}"
}

# Setup Pub/Sub

# chunk-gcs-upload-subscription

resource "google_pubsub_topic" "chunk_upload_topic" {
  name = "chunk-gcs-upload"
}

# test-run-state

resource "google_pubsub_topic" "test-run-state" {
  name = "test-run-state"
}

# wss-events
resource "google_pubsub_topic" "wss-events" {
  name = "wss-events"
}

resource "google_pubsub_subscription" "wss-subscription" {
  name = "wss-subscription"
  topic = google_pubsub_topic.wss-events.name
  enable_message_ordering = true
}

# storage-uploaded-chunks
resource "google_pubsub_topic" "storage-uploaded-chunks" {
  name = "storage-uploaded-chunks"
}

# Setup GKE

resource "google_container_cluster" "primary_cluster" {
  name = "${var.project_id}-cluster"
  location = var.zone

  remove_default_node_pool = true
  initial_node_count = 1

  master_auth {
    username = ""
    password = ""

    client_certificate_config {
      issue_client_certificate = false
    }
  }
}

# Managed Node pool

resource "google_container_node_pool" "primary_cluster_nodes" {
  name = "${google_container_cluster.primary_cluster.name}-node-pool"
  location = var.zone
  cluster = google_container_cluster.primary_cluster.name
  node_count = 3

  node_config {
    oauth_scopes = [
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring",
      "https://www.googleapis.com/auth/pubsub",
      "https://www.googleapis.com/auth/cloud-platform"
    ]

    machine_type = "e2-medium"
  }
}

# Create chunks storage bucket

resource "google_storage_bucket" "chunks_bucket" {
  name = "fleet-files-chunks"
}