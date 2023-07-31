#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd "$DIR"
go build goshooter.go

sudo cp goshooter /usr/local/bin/

if command -v /usr/local/bin/goshooter &> /dev/null
then
    PATH=$PATH:/usr/local/bin 
    echo "Installation completed successfully."
else
    echo "Installation failed. goshooter could not be copied to /usr/local/bin."
    exit 1
fi
