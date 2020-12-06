#!/bin/bash

# RUN THIS SCRIPT AS ROOT USER

# ---- Install docker ----
    apt-get -y install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common
curl -fsSL https://download.docker.com/linux/debian/gpg | sudo apt-key add -
apt-key fingerprint 0EBFCD88
add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/debian \
   $(lsb_release -cs) \
   stable"
apt-get update
apt-get install -y docker-ce
docker run hello-world

# Docker post install steps post-install
groupadd docker
usermod -aG docker $USER
systemctl enable docker


# ---- Install prerequisites for NFS mounting ----
apt-get -y update && \
apt-get -y install nfs-common


# ---- Create deploy target directories ----
mkdir -p /root/file-builder-deploy && \
mkdir -p /root/river-runner-deploy


# ---- Download google application credentials ----
gsutil cp gs://fleet-deploy/app_user.key.json ~/application_credentials.key.json


# --- Setup daemons ---
SETUP_DIR=~/setup
mkdir -p $SETUP_DIR

gsutil rsync -r gs://fleet-deploy/server $SETUP_DIR

# Start file-builder daemon
mv $SETUP_DIR/file-builder.service /etc/systemd/system && \
systemctl start file-builder

# Start river-runner daemon
mv $SETUP_DIR/river-runner.service /etc/systemd/system && \
systemctl start river-runner

# Remove setup dir
rm -r $SETUP_DIR
