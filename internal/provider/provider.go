// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ provider.Provider = &CheckProvider{}

type CheckProvider struct {
	version string
}

type CheckProviderModel struct {
}

func (p *CheckProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "check"
	resp.Version = p.version
}

func (p *CheckProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ``,
	}
}

func (p *CheckProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p *CheckProvider) Resources(ctx context.Context) []func() resource.Resource {
	return nil
}

func (p *CheckProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewCheckDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &CheckProvider{
			version: version,
		}
	}
}
