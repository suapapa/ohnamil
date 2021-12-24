#!/bin/bash

GOARCH=arm GOOS=linux go build
scp ohnamil pi@$1:~/
scp ohnamil.service pi@$1:~/
