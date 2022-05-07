#!/usr/bin/env bash

source env.sh
GOARCH=amd64 GOOS=darwin go build -ldflags "-s -w" -o ".workflow/exe" .
/usr/local/bin/envsubst >.workflow/info.plist <./info.plist.template
zip -r -D ./${alfred_workflow_name}-${alfred_workflow_version}.alfredworkflow .workflow/*

echo "release ./${alfred_workflow_name}-${alfred_workflow_version}.alfredworkflow success"