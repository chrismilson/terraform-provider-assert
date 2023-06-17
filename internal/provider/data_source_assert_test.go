// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAssertDataSource_WarnSeverity_ConditionValid_NoWarning(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				data "assert" "test" {
					severity  = "warn"
					condition = true
				}
				`,
				// There is no way to actually check if a warning fired or not.
			},
		},
	})
}

func TestAccAssertDataSource_WarnSeverity_ConditionInvalid_Warning(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				data "assert" "test" {
					severity  = "warn"
					condition = false
				}
				`,
				// There is no way to actually check if a warning fired or not.
			},
		},
	})
}

func TestAccAssertDataSource_ErrorSeverity_ConditionValid_NoError(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				data "assert" "test" {
					severity  = "error"
					condition = true
				}
				`,
			},
		},
	})
}

func TestAccAssertDataSource_ErrorSeverity_ConditionInvalid_ErrorIncludesSummary(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				data "assert" "test" {
					severity  = "error"
					condition = false
					summary   = "Test Summary"
				}
				`,
				ExpectError: regexp.MustCompile("Test Summary"),
			},
		},
	})
}

func TestAccAssertDataSource_ErrorSeverity_ConditionInvalid_ErrorIncludesDetail(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				data "assert" "test" {
					severity  = "error"
					condition = false
					detail    = "Test details."
				}
				`,
				ExpectError: regexp.MustCompile("Test details\\."),
			},
		},
	})
}
