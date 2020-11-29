#!/bin/sh

WORKDIR=~/deploy
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

# Run server
chmod +x ./server && \
export ENV_NAME=prod && \
sudo ./server