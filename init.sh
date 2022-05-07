#!/usr/bin/env bash

source env.sh
go mod init ${alfred_workflow_package}

if [[ "$enabledCobra" == "YES" ]]; then
	cobra init
else
cat > main.go <<EOF
package main

import (
	"errors"

	aw "github.com/deanishe/awgo"
)

var wf *aw.Workflow

func init() {
	wf = aw.New()
}

func run() {
	args := wf.Args()
	if len(args) == 0 {
		wf.FatalError(errors.New("please provide some input 👀"))
	}

	handlers := map[string]func(*aw.Workflow, []string) error{}

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
go get github.com/deanishe/awgo
fi

virtualenv -p ~/.pyenv/versions/3.10.0/bin/python venv
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
	<string>\${VERSION}</string>
	<key>webaddress</key>
	<string>${alfred_workflow_website}</string>
</dict>
</plist>
EOF