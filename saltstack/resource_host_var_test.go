package saltstack

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestsaltstackHostVar(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testsaltstackHostVarConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"saltstack_host_var.hostvar_1", "id", "www.example.com/foo"),
					resource.TestCheckResourceAttr(
						"saltstack_host_var.hostvar_1", "inventory_hostname", "www.example.com"),
					resource.TestCheckResourceAttr(
						"saltstack_host_var.hostvar_1", "key", "foo"),
					resource.TestCheckResourceAttr(
						"saltstack_host_var.hostvar_1", "value", "bar"),
					resource.TestCheckResourceAttr(
						"saltstack_host_var.hostvar_1", "variable_priority", "60"),
				),
			},
		},
	})
}

const testsaltstackHostVarConfig = `
  resource "saltstack_host_var" "hostvar_1" {
	  inventory_hostname = "www.example.com"
	  key                = "foo"
	  value              = "bar"
  }
`
