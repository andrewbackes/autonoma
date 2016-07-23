#!/bin/bash

cd "$(dirname "$0")"

ip="10.0.0.14"
user="pi"
remote_path="/usr/local/autonoma"

while [ true ]; do
    files=`find ./manual/ -type d -name .git -prune -o -mmin -1 -type f -print`

    if [ ! -z "$files" ]; then
        scp -r "${PWD}/manual" "${user}@${ip}:${remote_path}"
    else
        echo "No changes"
    fi

    sleep 60
done