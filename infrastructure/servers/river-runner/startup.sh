#!/bin/sh

# Run this script as a root user


WORKDIR=~/river-runner-deploy
mkdir -p $WORKDIR && cd $WORKDIR || exit

# Download and extract service archive
gsutil cp gs://fleet-deploy/binaries/river-runner.tar.gz river-runner.tar.gz
tar -xzf river-runner.tar.gz

# Download environment metadata
mkdir env
gsutil rsync -r gs://fleet-deploy/env/river-runner env

# Install prerequisites for nfs mounting
sudo apt-get -y update && \
sudo apt-get -y install nfs-common

# Download and set Google Cloud application credentials key
gsutil cp gs://fleet-deploy/app_user.key.json application_credentials.key.json
export GOOGLE_APPLICATION_CREDENTIALS=$PWD/application_credentials.key.json

# Run server
chmod +x ./server && \
export ENV_NAME=prod && \
./server &