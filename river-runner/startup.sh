#! /bin/bash
sudo su
#intall dependencies
apt-get -y update 
apt-get install -y nfs-common
# install Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh