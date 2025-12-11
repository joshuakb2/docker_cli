//go:build windows && 386

//go:generate goversioninfo -o=./winresources/resource.syso -icon=internal/assets/docker.ico -manifest=internal/assets/docker.exe.manifest ./winresources/versioninfo.json

package main

import _ "github.com/joshuakb2/docker_cli/cmd/docker/winresources"
