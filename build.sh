#!/usr/bin/env bash

docker build -f Dockerfile -t ffv-backend:0.1.0 .
docker save ffv-backend:0.1.0 > ~/ffv-backend.tar
microk8s ctr image import ~/ffv-backend.tar