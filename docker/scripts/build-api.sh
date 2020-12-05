#!/bin/bash 

BASEDIR=$(dirname "$0")
cd $BASEDIR/../../api;
docker build -t docker.pkg.github.com/kostaspt/domane/domane-api:latest .