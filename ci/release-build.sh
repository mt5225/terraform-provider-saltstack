#!/usr/bin/env bash

set -e -x

mkdir -p /go/src/github.com/nbering

ln -s  "$(pwd)/terraform-provider-ansible" "/go/src/github.com/mt5225/terraform-provider-saltstack"

pushd /go/src/github.com/mt5225/terraform-provider-saltstack
    ./scripts/build-release.sh
popd

pushd terraform-provider-ansible
    echo "$(git describe --abbrev=0)" > ../release-binaries/tag
    echo "Release $(git describe --abbrev=0)" > ../release-binaries/release-name
    mv pkg/*.zip ../release-binaries/
popd
