#!/usr/bin/env bash


# Create the container

docker build --no-cache --tag scratch-filesystem .
docker rm scratch-filesystem
docker run --name scratch-filesystem scratch-filesystem


# Export the filesystem

rm -rf export
mkdir export
cd export
docker export scratch-filesystem > filesystem.tar
tar -xvf filesystem.tar
rm filesystem.tar
