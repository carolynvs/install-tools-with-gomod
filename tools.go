// +build tools

package tools

// Any tools listed below will be installed into ./bin when you run make tools
import (
	_ "github.com/golang/mock/mockgen"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
  )