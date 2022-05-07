#!/usr/bin/env bash

source env.sh
GOARCH=amd64 GOOS=darwin go build -ldflags "-s -w" -o ".workflow/exe" .
/usr/local/bin/envsubst >.workflow/info.plist <./info.plist.template
cd .workflow && zip -r -D ../${alfred_workflow_name}-${alfred_workflow_version}.alfredworkflow ./*

echo "release ./${alfred_workflow_name}-${alfred_workflow_version}.alfredworkflow success"