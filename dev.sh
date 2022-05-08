#!/usr/bin/env bash

if [ -s .workflow/info.plist ]; then
    cp .workflow/info.plist info.plist.template
fi
plutil -replace version -string '${VERSION}' info.plist.template
GOARCH=amd64 GOOS=darwin go build -ldflags "-s -w" -o ".workflow/exe" .
/usr/local/bin/envsubst >.workflow/info.plist <./info.plist.template
./venv/bin/python ./workflow-install.py -s .workflow

echo "run dev.sh success"