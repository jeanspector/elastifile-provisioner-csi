#!/usr/bin/env bash

POD_MANIFEST=$1
: ${POD_MANIFEST:="pod-with-io.yaml"}

../deploy/deploy-plugin.sh
./create-pod.sh $1

