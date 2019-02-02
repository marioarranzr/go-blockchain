#!/bin/bash
# go-blockchain update_concourse.sh

fly -t main sp -p go-blockchain -c ci/pipeline.yml -l /.credentials.yml
