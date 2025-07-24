#!/bin/bash
docker rm -f algorithm
docker pull hunterhug/algorithm:docsify
docker run --name algorithm -d -p 12346:3000 hunterhug/algorithm:docsify