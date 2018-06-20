#!/bin/bash

docker run --rm -v "$(pwd)":/ssm-edit -w /ssm-edit tcnksm/gox:latest sh -c "go get -d ./... && make build-all"

PWD=$(pwd)
TAG=$(git describe --tags)

cd assets

for i in `ls -1`
do
  cd ${i} && tar czf ssm-edit_${TAG}_${i}.tar.gz ssm-edit && cd -
done

cd $PWD

shasum -a 256 assets/darwin_amd64/ssm-edit_${TAG}_darwin_amd64.tar.gz
