#!/bin/bash

GOARCH=arm GOOS=linux GOARM=6 go build
scp ohnamil $1:~/
# scp ohnamil.service $1:~/
