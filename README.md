# git-cheat

A Go-based command-line utility for manipulating Git commit metadata, allowing you to create backdated commits and modify the author or timestamp of existing commits.

**Warning:** modifying the author or timestamp of an existing commit uses `git-filter-branch` command, which creates a new commit object with different hashes. This will cause conflicts if history has been pushed to shared repository, and therefore should be used with caution in collaborative environments.

**Usage:** there are no dependencies to install. Simply build the CLI using `go build -o git-cheat *.go`.