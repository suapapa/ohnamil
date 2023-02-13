#!/bin/bash
set -e
go build
GOOGLE_APPLICATION_CREDENTIALS=$(pwd)/_secret/homin-gadget-7e5ac18defca.json ./ohnamil -kep 45 -gcal "ff4500@gmail.com" -d Cassian -n