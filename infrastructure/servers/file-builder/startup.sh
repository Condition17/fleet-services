#!/bin/sh

# Run this script as a root user


WORKDIR=~/file-builder-deploy
mkdir -p $WORKDIR && cd $WORKDIR || exit

# Download and extract service archive
gsutil cp gs://fleet-deploy/binaries/file-builder.tar.gz file-builder.tar.gz
tar -xzf file-builder.tar.gz

# Download environment metadata
mkdir env
gsutil rsync -r gs://fleet-deploy/env/file-builder env

# Install prerequisites for nfs mounting
sudo apt-get -y update && \
sudo apt-get -y install nfs-common

# Download and set Google Cloud application credentials key
gsutil cp gs://fleet-deploy/app_user.key.json application_credentials.key.json
export GOOGLE_APPLICATION_CREDENTIALS=$PWD/application_credentials.key.json

# Run server
chmod +x ./server && \
export ENV_NAME=prod && \
./server