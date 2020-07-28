#!/usr/bin/env bash
root_dir=$(cd "$(dirname "$0")"; cd ..; pwd)

# grep all request validator
sh $root_dir/bin/validator_grep.sh

# request validator code
go build -o $root_dir/tools/validator_gen/validator_gen $root_dir/tools/validator_gen/validator.go

$root_dir/tools/validator_gen/validator_gen -pb_dir=$root_dir/pb -validator_log_dir=$root_dir

echo "generate validator request interceptor success"

exit 0
