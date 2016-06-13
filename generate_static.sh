#!/bin/sh

cd htdocs/json && curl -k https://localhost:8080/json/tree.json | tee tree.json | jsonformat | egrep -o 'json/.*.json' | while read line
do
    curl -k -O https://localhost:8080/$line 
done
