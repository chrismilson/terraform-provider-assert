// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &CheckDataSource{}

func NewCheckDataSource() datasource.DataSource {
	return &CheckDataSource{}
}

type CheckDataSource struct{}

type CheckDataSourceModel struct {
	Id         types.String          `tfsdk:"id"`
	Assertions []CheckAssertionModel `tfsdk:"assert"`
}

type CheckAssertionModel struct {
	Condition      types.Bool   `tfsdk:"condition"`
	ErrorMessage   types.String `tfsdk:"error_message"`
	WarningMessage types.String `tfsdk:"warning_message"`
}

func (d *CheckDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName
}

func AssertionBlock(additional map[string]schema.Attribute) schema.ListNestedBlock {
	attributes := map[string]schema.Attribute{
		"error_message": schema.StringAttribute{
			MarkdownDescription: "The message included in the error if the condition is not satisfied.",
			Optional:            true,
		},
		"warning_message": schema.StringAttribute{
			MarkdownDescription: "The message included in the warning if the condition is not satisfied.",
			Optional:            true,
		},
	}

	for k, v := range additional {
		attributes[k] = v
	}

	return schema.ListNestedBlock{NestedObject: schema.NestedBlockObject{Attributes: attributes}}
}

func (d *CheckDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example data source",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "This is set to a random value at read time.",
				Computed:            true,
			},
		},
		Blocks: map[string]schema.Block{
			"assert": AssertionBlock(map[string]schema.Attribute{
				"condition": schema.BoolAttribute{
					Required: true,
				},
			}),
			// TODO assert_equal
			// TODO assert_greater_than
			// TODO etc.
		},
	}
}

func (d *CheckDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data CheckDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	for i, assertion := range data.Assertions {
		condition_path := path.Path{}.AtName("assert").AtListIndex(i).AtName("condition")
		if !assertion.Condition.ValueBool() {
			if assertion.ErrorMessage != types.StringNull() {
				resp.Diagnostics.AddAttributeError(
					condition_path,
					"Unsatisfied Condition",
					assertion.ErrorMessage.ValueString(),
				)
			}
			if assertion.WarningMessage != types.StringNull() {
				resp.Diagnostics.AddAttributeWarning(
					condition_path,
					"Unsatisfied Condition",
					assertion.WarningMessage.ValueString(),
				)
			}
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = types.StringValue(fmt.Sprintf("%d", rand.Int()))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *CheckDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
}
