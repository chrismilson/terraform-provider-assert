// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAssertDataSource_Condition_Valid_Warning_NotFired(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				data "assert" "test" {
					condition       = true
					warning_message = "test warning"
				}
				`,
				// There is no way to actually check if a warning fired.
			},
		},
	})
}

func TestAccAssertDataSource_Condition_Invalid_Warning_Fired(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				data "assert" "test" {
					condition       = false
					warning_message = "test warning"
				}
				`,
				// There is no way to actually check if a warning fired.
			},
		},
	})
}

func TestAccAssertDataSource_Condition_Valid_Error_NotFired(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				data "assert" "test" {
					condition     = true
					error_message = "test error"
				}
				`,
			},
		},
	})
}

func TestAccAssertDataSource_Condition_Invalid_Error_Fired(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				data "assert" "test" {
					condition     = false
					error_message = "test error"
				}
				`,
				ExpectError: regexp.MustCompile("test error"),
			},
		},
	})
}

func TestAccAssertDataSource_ErrorAndWarningMessageConfig_Invalid_Error_Fired(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				data "assert" "test" {
					condition       = true
					error_message   = "test error"
					warning_message = "test warning"
				}
				`,
				ExpectError: regexp.MustCompile("Invalid Attribute Combination"),
			},
		},
	})
}

func TestAccAssertDataSource_NoMessageConfig_Invalid_Error_Fired(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				data "assert" "test" {
					condition       = true
				}
				`,
				ExpectError: regexp.MustCompile("Missing Attribute Configuration"),
			},
		},
	})
}
