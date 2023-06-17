// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource = &AssertDataSource{}
)

func NewAssertDataSource() datasource.DataSource {
	return &AssertDataSource{}
}

type AssertDataSource struct{}

type AssertDataSourceModel struct {
	Id        types.String `tfsdk:"id"`
	Condition types.Bool   `tfsdk:"condition"`
	Detail    types.String `tfsdk:"detail"`
	Severity  types.String `tfsdk:"severity"`
	Summary   types.String `tfsdk:"summary"`
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
			"detail": schema.StringAttribute{
				MarkdownDescription: "A more in-depth description of the issue.",
				Required:            true,
			},
			"severity": schema.StringAttribute{
				MarkdownDescription: "The severity of the issue. Can be 'error' or 'warn'. Setting to 'error' will halt execution of terraform, while setting to 'warn' will not.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("error", "warn"),
				},
			},
			"summary": schema.StringAttribute{
				MarkdownDescription: "A short summary to identify the issue.",
				Required:            false,
			},
			// TODO assert_equal
			// TODO etc.
		},
	}
}

func (d *AssertDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AssertDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if !data.Condition.ValueBool() {
		path := path.Path{}.AtName("condition")
		summary := data.Summary.ValueString()
		detail := data.Detail.ValueString()

		if data.Severity.ValueString() == "error" {
			resp.Diagnostics.AddAttributeError(path, summary, detail)
		}
		if data.Severity.ValueString() == "warn" {
			resp.Diagnostics.AddAttributeWarning(path, summary, detail)
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
