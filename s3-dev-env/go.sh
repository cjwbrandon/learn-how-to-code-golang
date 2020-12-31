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