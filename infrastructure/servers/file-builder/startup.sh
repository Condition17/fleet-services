#!/bin/sh

# TODO: ensure port forwarding is setup
WORKDIR=~/deploy
mkdir -p $WORKDIR && cd $WORKDIR || exit

# Download and extract service archive
gsutil cp gs://fleet-deploy-binaries/file-builder.tar.gz
tar -xzf file-builder.tar.gz .

# Run server
chmod +x ./server && ./server