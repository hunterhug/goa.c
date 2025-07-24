#!/bin/bash
docker run -it --rm -p 12346:3000 --name=docsify -v $(pwd):/docs hunterhug/algorithm:docsify