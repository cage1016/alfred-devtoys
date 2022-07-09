#!/usr/bin/env bash

source env.sh
go mod init ${alfred_workflow_package}

if [[ "$enabledCobra" == "YES" ]]; then
	cobra init
	go get github.com/spf13/cobra
else
cat > main.go <<EOF
package main

import (
	"errors"
	"log"
	"os"
	"os/exec"

	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
)

const updateJobName = "checkForUpdate"

var (
	repo = "${alfred_workflow_package#*/}"
	wf   *aw.Workflow
)

func init() {
	wf = aw.New(update.GitHub(repo), aw.HelpURL(repo+"/issues"))
}

func run() {
	args := wf.Args()
	if len(args) == 0 {
		wf.FatalError(errors.New("please provide some input 👀"))
	}

	handlers := map[string]func(*aw.Workflow, []string) error{
    "update": func(wf *aw.Workflow, _ []string) error {
			wf.Configure(aw.TextErrors(true))
			log.Println("Checking for updates...")
			if err := wf.CheckForUpdate(); err != nil {
				wf.FatalError(err)
			}
			return nil
		},
  }

	if wf.UpdateCheckDue() && !wf.IsRunning(updateJobName) {
		log.Println("Running update check in background...")

		cmd := exec.Command(os.Args[0], "update")
		if err := wf.RunInBackground(updateJobName, cmd); err != nil {
			log.Printf("Error starting update check: %s", err)
		}
	}

	if wf.UpdateAvailable() {
		wf.Configure(aw.SuppressUIDs(true))
		log.Println("Update available!")
		wf.NewItem("An update is available!").
			Subtitle("⇥ or ↩ to install update").
			Valid(false).
			Autocomplete("workflow:update").
			Icon(&aw.Icon{Value: "update-available.png"})
	}


	h, found := handlers[args[0]]
	if !found {
		wf.FatalError(errors.New("command not recognized 👀"))
	}

	err := h(wf, args[1:])
	if err != nil {
		wf.FatalError(err)
	}

	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
EOF
fi

## prepare go mod package
go get github.com/deanishe/awgo

virtualenv -p ~/.pyenv/versions/$(<.python-version)/bin/python venv
source venv/bin/activate
pip install -r requirements.txt

# create info.plist.template
cat > info.plist.template << EOF
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>bundleid</key>
	<string>${alfred_workflow_bundleid}</string>
	<key>connections</key>
	<dict/>
	<key>createdby</key>
	<string>${alfred_workflow_created_by}</string>
	<key>description</key>
	<string>${alfred_workflow_description}</string>
	<key>disabled</key>
	<false/>
	<key>name</key>
	<string>${alfred_workflow_name}</string>
	<key>objects</key>
	<array/>
	<key>readme</key>
	<string></string>
	<key>uidata</key>
	<dict/>
	<key>version</key>
	<string></string>
	<key>webaddress</key>
	<string>${alfred_workflow_website}</string>
</dict>
</plist>
EOF

## create .github/workflows/release.yml
mkdir -p .github/workflows
cat > .github/workflows/release.yml <<EOF
name: Release
on:
  release:
    types:
      - published

jobs:
  build:
    strategy:
      matrix:
        go-version: [1.17.x]
        platform: [macOS-latest]
    runs-on: \${{ matrix.platform }}
    steps:
    - uses: actions/checkout@v2
    - name: Install Go
      if: success()
      uses: actions/setup-go@v2
      with:
        go-version: \${{ matrix.go-version }}
    - name: Run unit tests
      run: go test -v ./...
    - name: Parse Event
      id: event
      run: |
        echo "::set-output name=tag::\$(jq -r '.release.tag_name' "\${GITHUB_EVENT_PATH}" | sed s/^v//)"
    - name: Build golang build
      id: build
      run: |
        sed -i -e "s|export alfred_workflow_version=\"0.1.0\"|export alfred_workflow_version=\"\${{ steps.event.outputs.tag }}\"|g" env.sh
        ./release.sh
        echo "::set-output name=artifact::\$(echo "${alfred_workflow_name}-\${{ steps.event.outputs.tag }}.alfredworkflow" | sed "s/ //g")"
    - uses: shogo82148/actions-upload-release-asset@v1
      with:
        upload_url: \${{ github.event.release.upload_url }}
        asset_path: "\${{ steps.build.outputs.artifact }}"
EOF