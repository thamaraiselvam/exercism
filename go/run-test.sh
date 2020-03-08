#!/bin/bash

set -e

dirs=$(ls -d */)

for dir in ${dirs[@]}
do
    cd $(pwd)/${dir}
    go test . --bench --benchmem
    cd ..
done


