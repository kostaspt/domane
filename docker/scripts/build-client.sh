#!/bin/bash 

BASEDIR=$(dirname "$0")
cd $BASEDIR/../../client;
docker build -t ghcr.io/kostaspt/domane-client:latest .