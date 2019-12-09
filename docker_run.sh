#!/bin/bash
docker rm -f algorithm
docker run --name algorithm -d -p 9999:80 algorithm:latest