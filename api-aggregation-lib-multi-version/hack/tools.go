//go:build tools
// +build tools

// This package imports things required by build scripts, to force `go mod` to see them as dependencies
package tools

import _ "k8s.io/code-generator"
import _ "github.com/eonvon/cloud-native-tour/api/hello.eonvon.github.io/v1"
import _ "github.com/eonvon/cloud-native-tour/api/hello.eonvon.github.io/v2"
import _ "github.com/eonvon/cloud-native-tour/api/transformation/v1beta1"
