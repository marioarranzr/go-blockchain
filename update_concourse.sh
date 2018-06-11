#!/bin/bash
# go-blockchain update_concourse.sh

fly -t ci set-pipeline -p go-blockchain -c ci/pipeline.yml --load-vars-from ci/.credentials.yml
