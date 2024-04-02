// package repo contains data values that are common to all CUE configurations
// in this repo. The list of configurations includes GitHub workflows, but also
// things like gerrit configuration etc.
package repo

import (
	"cuelang.org/c/cue/internal/ci/base"
)

base

githubRepositoryPath: "cue-lang/libcue"

botGitHubUser:      "cueckoo"
botGitHubUserEmail: "cueckoo@gmail.com"

defaultBranch: "main"

linuxMachine: "ubuntu-22.04"

latestGo: "1.21.x"
