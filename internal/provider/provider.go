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

var _ provider.Provider = &AssertProvider{}

type AssertProvider struct {
	version string
}

type AssertProviderModel struct {
}

func (p *AssertProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "assert"
	resp.Version = p.version
}

func (p *AssertProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ``,
	}
}

func (p *AssertProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p *AssertProvider) Resources(ctx context.Context) []func() resource.Resource {
	return nil
}

func (p *AssertProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewAssertDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &AssertProvider{
			version: version,
		}
	}
}
