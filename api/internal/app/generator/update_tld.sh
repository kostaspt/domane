#!/bin/bash

TLDs=$(curl -s -L 'https://raw.githubusercontent.com/umpirsky/tld-list/master/data/en/tld.json' --header 'Content-Type: application/json')
sed -i '' 's/src := .*/src := `'$TLDs'`/' ./tlds.go