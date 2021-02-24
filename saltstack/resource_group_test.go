package saltstack

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestsaltstackGroup(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testsaltstackGroupConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"saltstack_group.group_1", "id", "group_1"),
					resource.TestCheckResourceAttr(
						"saltstack_group.group_1", "children.0", "foo"),
					resource.TestCheckResourceAttr(
						"saltstack_group.group_1", "children.2", "baz"),
					resource.TestCheckResourceAttr(
						"saltstack_group.group_1", "vars.foo", "bar"),
					resource.TestCheckResourceAttr(
						"saltstack_group.group_1", "vars.bar", "2"),
					resource.TestCheckResourceAttr(
						"saltstack_group.group_1", "variable_priority", "50"),
				),
			},
		},
	})
}

const testsaltstackGroupConfig = `
  resource "saltstack_group" "group_1" {
    inventory_group_name = "group_1"
    children = ["foo", "bar", "baz"]
    vars = {
      foo = "bar"
      bar = 2
    }
  }
`
