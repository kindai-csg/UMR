#!/bin/ash

###############################
# generate interface`s mock code
# ./gen_mock.sh  interface_source_code_dir mock_source_code_file_name
###############################

SCRIPT_DIR="$(cd $(dirname $0); pwd)/"

export GOFLAGS=-mod=vendor
export GO111MODULE=on

/.go/bin/mockgen -source $1 -destination $SCRIPT_DIR$2
