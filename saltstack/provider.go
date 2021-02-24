package saltstack

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"saltstack_host":      resourceHost(),
			"saltstack_host_var":  resourceHostVar(),
			"saltstack_group":     resourceGroup(),
			"saltstack_group_var": resourceGroupVar(),
		},
	}
}
