#!/bin/bash

dir="/usr/local/autonoma/"

if [ -d "$dir" ]; then
    cd "$dir"
    sudo git checkout master
    sudo git pull
else
    cd "/usr/local/"
    sudo git clone https://github.com/andrewbackes/autonoma
fi
