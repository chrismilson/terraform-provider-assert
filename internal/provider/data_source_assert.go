// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource                     = &AssertDataSource{}
	_ datasource.DataSourceWithConfigValidators = &AssertDataSource{}
)

func NewAssertDataSource() datasource.DataSource {
	return &AssertDataSource{}
}

type AssertDataSource struct{}

type AssertDataSourceModel struct {
	Id             types.String `tfsdk:"id"`
	Condition      types.Bool   `tfsdk:"condition"`
	ErrorMessage   types.String `tfsdk:"error_message"`
	WarningMessage types.String `tfsdk:"warning_message"`
}

func (d *AssertDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName
}

func (d *AssertDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Asserts that a condition is true, otherwise logs a warning or error to terraform.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "This is set to a random value at read time.",
				Computed:            true,
			},
			"condition": schema.BoolAttribute{
				MarkdownDescription: "The assertion that is expected to be true.",
				Required:            true,
			},
			"error_message": schema.StringAttribute{
				MarkdownDescription: "The message included in the error if the condition is not satisfied.",
				Optional:            true,
			},
			"warning_message": schema.StringAttribute{
				MarkdownDescription: "The message included in the warning if the condition is not satisfied.",
				Optional:            true,
			},
			// TODO assert_equal
			// TODO etc.
		},
	}
}

func (d *AssertDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(
			path.MatchRoot("error_message"),
			path.MatchRoot("warning_message"),
		),
	}
}

func (d *AssertDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AssertDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	condition_path := path.Path{}.AtName("condition")
	if !data.Condition.ValueBool() {
		if data.ErrorMessage != types.StringNull() {
			resp.Diagnostics.AddAttributeError(
				condition_path,
				"Unsatisfied Condition",
				data.ErrorMessage.ValueString(),
			)
		}
		if data.WarningMessage != types.StringNull() {
			resp.Diagnostics.AddAttributeWarning(
				condition_path,
				"Unsatisfied Condition",
				data.WarningMessage.ValueString(),
			)
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = types.StringValue(fmt.Sprintf("%d", rand.Int()))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *AssertDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
}
