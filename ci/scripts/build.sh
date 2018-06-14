#!/bin/bash
# go-blockchain build.sh

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

# Put the binary go-blockchain filename in /dist
go build -o dist/go-blockchain ./main.go
#cp ./main.go dist/go-blockchain
cp ./main.go dist

# cp the Dockerfile into /dist
cp Dockerfile dist/Dockerfile

# Check
echo "List whats in the /dist directory"
ls -lat dist

# Move what you need to $GOPATH/dist
# BECAUSE the resource type docker-image works in /dist.
cp -R ./dist $GOPATH/.

# Check whats here
echo "List whats in top directory"
ls -lat 

echo ""