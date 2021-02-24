package saltstack

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestsaltstackGroupVar(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testsaltstackGroupVarConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"saltstack_group_var.groupvar_1", "id", "web/baz"),
					resource.TestCheckResourceAttr(
						"saltstack_group_var.groupvar_1", "inventory_group_name", "web"),
					resource.TestCheckResourceAttr(
						"saltstack_group_var.groupvar_1", "key", "baz"),
					resource.TestCheckResourceAttr(
						"saltstack_group_var.groupvar_1", "value", "bup"),
					resource.TestCheckResourceAttr(
						"saltstack_group_var.groupvar_1", "variable_priority", "60"),
				),
			},
		},
	})
}

const testsaltstackGroupVarConfig = `
  resource "saltstack_group_var" "groupvar_1" {
	  inventory_group_name = "web"
	  key                  = "baz"
	  value                = "bup"
  }
`
