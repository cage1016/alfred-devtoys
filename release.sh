#!/usr/bin/env bash

source env.sh
GOARCH=amd64 GOOS=darwin go build -ldflags "-s -w" -o ".workflow/exe" .
/usr/local/bin/envsubst >.workflow/info.plist <./info.plist.template
cd .workflow && zip -r -D ../$(echo "${alfred_workflow_name}-${alfred_workflow_version}.alfredworkflow" | sed "s/ //g") ./*

echo "release ./$(echo "${alfred_workflow_name}-${alfred_workflow_version}.alfredworkflow" | sed "s/ //g") success"