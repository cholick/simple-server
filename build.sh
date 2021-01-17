#!/usr/bin/env bash

make

docker build docker -t cholick/simple-server
docker push cholick/simple-server
