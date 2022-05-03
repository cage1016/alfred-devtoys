#!/usr/bin/env bash

source env.sh
go mod init ${alfred_workflow_package}
cobra init
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

# create .vscode/tasks.json
cat > .vscode/tasks.json << EOF
{
    // See https://go.microsoft.com/fwlink/?LinkId=733558
    // for the documentation about the tasks.json format
    "version": "2.0.0",
    "tasks": [
        {
            "label": "go-build",
            "type": "shell",
            "command": "go",
            "args": [
                "build",
                "-o",
                ".workflow/exe",
            ],
            "group": {
                "kind": "build",
                "isDefault": true
            },
            "problemMatcher": [],
        },
        {
            "label": "prepare-info.plist",
            "type": "shell",
            "command": "\${config:envsubst.path}",
            "args": [
                ">.workflow/info.plist",
                "<./info.plist.template",
            ],
            "group": {
                "kind": "build",
                "isDefault": true
            },
            "problemMatcher": [],
        },
        {
            "label": "install",
            "type": "process",
            "command": "\${config:python.pythonPath}",
            "args": [
                "\${workspaceFolder}/workflow-install.py",
                "-s",
                ".workflow"
            ],
            "group": {
                "kind": "build",
                "isDefault": true
            },
            "problemMatcher": [],
            "dependsOn": [
                "go-build",
                "prepare-info.plist",
            ],
        },
    ]
}
EOF

# create build.sh
cat > build.sh << EOF
#!/bin/sh

set -o errexit
set -o nounset

archive="${alfred_workflow_name}-\${VERSION}.alfredworkflow"

echo "Building go binary:"
GOARCH=amd64 GOOS=darwin go build -ldflags "-s -w" -o ".workflow/exe"\ .

echo ""
echo "Crearing archive:"
(
    envsubst >.workflow/info.plist <./info.plist.template
    cd ./.workflow || exit
    zip -r "../${archive}" ./*
)

echo ""
echo "Build completed: \"${archive}\""
EOF