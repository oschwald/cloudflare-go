//go:build tools
// +build tools

package tools

//go:generate go install golang.org/x/tools/gopls@latest
//go:generate go install github.com/orijtech/structslop/cmd/structslop@latest
//go:generate go install github.com/ramya-rao-a/go-outline@latest
//go:generate go install github.com/cweill/gotests/gotests@latest
//go:generate go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
//go:generate go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest
//go:generate go install github.com/go-delve/delve/cmd/dlv@latest
//go:generate go install golang.org/x/lint/golint@latest
//go:generate go install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
//go:generate go install github.com/bflad/tfproviderlint/cmd/tfproviderlintx
//go:generate go install github.com/client9/misspell/cmd/misspell
//go:generate go install github.com/golangci/golangci-lint/cmd/golangci-lint
//go:generate go install github.com/hashicorp/go-changelog/cmd/changelog-build
//go:generate go install github.com/google/go-github/github
//go:generate go install golang.org/x/oauth2

import (
	_ "github.com/bflad/tfproviderlint/cmd/tfproviderlintx"
	_ "github.com/client9/misspell/cmd/misspell"
	_ "github.com/cweill/gotests/gotests"
	_ "github.com/go-delve/delve/cmd/dlv"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/google/go-github/github"
	_ "github.com/hashicorp/go-changelog/cmd/changelog-build"
	_ "github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs"
	_ "github.com/orijtech/structslop/cmd/structslop"
	_ "github.com/ramya-rao-a/go-outline"
	_ "github.com/uudashr/gopkgs/v2/cmd/gopkgs"
	_ "golang.org/x/lint/golint"
	_ "golang.org/x/oauth2"
	_ "golang.org/x/tools/gopls"
)
