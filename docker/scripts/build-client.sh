#!/bin/bash 

BASEDIR=$(dirname "$0")
cd $BASEDIR/../../client;
docker build -t docker.pkg.github.com/kostaspt/domane/domane-client:latest .