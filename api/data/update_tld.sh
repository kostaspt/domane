#!/bin/bash

curl -L 'https://raw.githubusercontent.com/umpirsky/tld-list/master/data/en/tld.json' \
    --header 'Content-Type: application/json' > tld.json