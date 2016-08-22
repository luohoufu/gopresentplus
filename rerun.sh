#!/bin/bash
current_dir=$(cd $(dirname $0); pwd)
go clean
rm -rf ./static/tmp
go build
./gopresentplus
