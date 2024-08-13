package main

import (
	"github.com/Hostersi/terraform-provider-sops/sops"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: sops.Provider,
	})
}
