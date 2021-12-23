#!/bin/bash

GOARCH=arm GOOS=linux go build
scp kwcal-frame pi@$1:~/
scp kwcal-frame.service pi@$1:~/
