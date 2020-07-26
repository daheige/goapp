#!/bin/bash
root_dir=$(cd "$(dirname "$0")"; cd ..; pwd)

# protoc inject tag
protoc_inject=$(which "protoc-go-inject-tag")

if [ -z $protoc_inject ]; then
    echo 'Please install protoc-go-inject-tag'
    echo "Please run go get -u github.com/favadi/protoc-go-inject-tag"
    exit 0
fi

for file in $root_dir/pb/*.pb.go; do
  echo "protoc inject tag file: "$file
  $protoc_inject -input=$file
done
