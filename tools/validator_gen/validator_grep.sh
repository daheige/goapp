#!/bin/bash
root_dir=$(cd "$(dirname "$0")"; cd ../../; pwd)

echo $root_dir;

grep "@validator=*" $root_dir/protos/*.proto > $root_dir/tools/validator_gen/validator.log

cd $root_dir/tools/validator_gen

#sed -i "" 's/\/\/ \@validator\=//g' `grep validator -rl ./`
