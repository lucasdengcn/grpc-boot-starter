#!/bin/sh

go test -v -coverpkg=./... -coverprofile=profile.cov ./...

while read p || [ -n "$p" ] 
do  
sed -i '' "/${p//\//\\/}/d" ./profile.cov 
done < ./.coverage-filter.txt

go tool cover -html profile.cov