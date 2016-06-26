#!/bin/bash

dir="/usr/local/autonoma/"

if [ -d "$dir" ]; then
    cd "$dir"
    git checkout master
    git pull
else
    cd "/usr/local/"
    git clone https://github.com/andrewbackes/autonoma
fi
