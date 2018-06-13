#!/bin/bash
# go-blockchain unit-test.sh

set -e -x

# The code is located in /go-blockchain
echo "pwd is: " $PWD
echo "List whats in the current directory"
ls -lat 

# Setup the gopath based on current directory.
export GOPATH=$PWD

# Now we must move our code from the current directory ./go-blockchain to $GOPATH/src/github.com/mario800ml/go-blockchain
mkdir -p src/github.com/mario800ml/
cp -R ./go-blockchain src/github.com/mario800ml/.

# All set and everything is in the right place for go
echo "Gopath is: " $GOPATH
echo "pwd is: " $PWD
cd src/github.com/mario800ml/go-blockchain
ls -lat

go get ./...

# RUN unit_tests and it shows the percentage coverage
# print to stdout and file using tee
go test -cover ./... | tee test_coverage.txt
# add some whitespace to the begining of each line
sed -i -e 's/^/     /' test_coverage.txt

# Move to coverage-results directory.
mv test_coverage.txt $GOPATH/coverage-results/.
