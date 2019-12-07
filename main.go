package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/victorsalaun/terraform-provider-nexus-repository-manager/nexus-repository-manager"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: nexus.Provider})
}
