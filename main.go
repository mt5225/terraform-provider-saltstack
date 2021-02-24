package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/mt5225/terraform-provider-saltstack/saltstack"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: saltstack.Provider,
	})
}
