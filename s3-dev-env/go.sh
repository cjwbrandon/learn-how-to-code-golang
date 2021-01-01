# Go version
go version

# Go Environment Variables
go env

# Commands
go help
go help [topic]

# Format code
go fmt
go fmt ./... # Formatting entire repo

# Run file
go run {file_path}.go

# Builds file, reports errors, puts executable into current folder
go build

# Compiles program (build it), names the executable the folder name holding the code, puts in workspace/bin
go install {file_path}.go
go install *.go

# Get repo
go get -d {url}/{user}/{repo_name}

# Go Modules
# Create a new module
go mod init {namespace} # e.g. go mod init {domain}/{username}/{repo}
# Creates go.mod & go.sum
# go.mod stores the namespace
# go.sum ensures future downloads of these modules retrieve the same bits as the first download

# package-building
go build
go test

# Print current module's dependencies
go list -m all
go list -m -version {module} # list available module versions

# Change required version of a dependency
go get {module}
go get {module}@{version}

# Remove unused dependencies
go mod tidy
