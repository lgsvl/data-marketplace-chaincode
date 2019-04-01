#!/usr/bin/env bash

set -e -u

cd $(dirname $0)/..

echo "Setting up ginkgo and gomega"
go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega

echo "Setting up glide"
go get -v github.com/Masterminds/glide

echo "Running glide up with strip vendor flag"
glide up --strip-vendor

ginkgo -r \
  -skipPackage utils \
  -cover \
  -coverpkg $(go list ./... | grep -v /cmd/ | xargs echo | tr ' ' ',') \
  -covermode atomic \
  "$@"