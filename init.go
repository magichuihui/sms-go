package sdk

// init.go - just import the sub packages

// Package sdk imports all sub packages to build all of them when calling `go install', `go build'
// or `go get' commands.
import (
        _ "github.com/magichuihui/sms-go/baidu"
)
