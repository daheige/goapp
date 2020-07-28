#!/bin/bash
root_dir=$(cd "$(dirname "$0")"; cd ..; pwd)

echo $root_dir;

grep "\w*@validator=*" $root_dir/protos/*.proto > $root_dir/validator.log
