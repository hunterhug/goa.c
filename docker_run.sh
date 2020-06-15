#!/bin/bash
docker rm -f algorithm
docker run --name algorithm -d -p 12346:80 hunterhug/algorithm:latest