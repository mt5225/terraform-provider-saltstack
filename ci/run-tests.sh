#!/usr/bin/env bash

set -e -x

cat terraform-provider-ansible/.git/ref

mkdir -p /go/src/github.com/nbering

ln -s  "$(pwd)/terraform-provider-ansible" "/go/src/github.com/mt5225/terraform-provider-saltstack"

pushd /go/src/github.com/mt5225/terraform-provider-saltstack
    make test
    make testacc
    make vet
popd
