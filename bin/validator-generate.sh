#!/usr/bin/env bash
root_dir=$(cd "$(dirname "$0")"; cd ..; pwd)

# grep all request validator
sh $root_dir/bin/validator_grep.sh

validatorGenExec=$(which "validator_gen")
if [ -z $validatorGenExec ]; then
  # request validator code
  go build -o $root_dir/tools/validator_gen/validator_gen $root_dir/tools/validator_gen/validator.go
  chmod +x $root_dir/tools/validator_gen/validator_gen

  # copy this binary file to $GOPATH/bin
  cp $root_dir/tools/validator_gen/validator_gen $GOPATH/bin/validator_gen

  validatorGenExec=$root_dir/tools/validator_gen/validator_gen
fi

$validatorGenExec -pb_dir=$root_dir/pb -validator_log_dir=$root_dir

echo "generate validator request interceptor success"

exit 0
