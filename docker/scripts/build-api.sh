#!/bin/bash 

BASEDIR=$(dirname "$0")
cd $BASEDIR/../../api;
docker build -t ghcr.io/kostaspt/domane-api:latest .