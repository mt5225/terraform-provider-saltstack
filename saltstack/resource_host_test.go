package saltstack

import (
	"fmt"
	"testing"

	r "github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestsaltstackHost(t *testing.T) {
	r.Test(t, r.TestCase{
		Providers: testAccProviders,
		Steps: []r.TestStep{
			r.TestStep{
				Config: testsaltstackHostConfig,
				Check: func(s *terraform.State) error {
					rID := s.RootModule().Outputs["host_id"].Value
					if "example.medstack.net" != rID {
						return fmt.Errorf("Unexpected value for resource ID: %s", rID)
					}

					return nil
				},
			},
		},
	})
}

var testsaltstackHostConfig = `
resource "saltstack_host" "test" {
	inventory_hostname = "example.medstack.net"
	groups = ["dbservers"]
	vars = {
		"first" = "foo"
		"second" = "bar"
	}
}

output "host_id" {
	value = "${saltstack_host.test.id}"
}
`
