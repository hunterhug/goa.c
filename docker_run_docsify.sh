#!/bin/bash
docker rm -f algorithm_docsify
docker run --name algorithm_docsify -d -p 12346:3000 hunterhug/algorithm:docsify