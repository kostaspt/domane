#!/bin/bash 

BASEDIR=$(dirname "$0")
$BASEDIR/build-api.sh && $BASEDIR/build-client.sh