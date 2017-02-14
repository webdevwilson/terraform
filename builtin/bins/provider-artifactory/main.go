package main

import (
	"github.com/hashicorp/terraform/builtin/providers/artifactory"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: artifactory.Provider,
	})
}
