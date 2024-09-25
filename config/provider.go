/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/headyj/provider-neo4j/config/database"
	"github.com/headyj/provider-neo4j/config/grant"
	"github.com/headyj/provider-neo4j/config/role"
	"github.com/headyj/provider-neo4j/config/user"
)

const (
	resourcePrefix = "neo4j"
	modulePath     = "github.com/headyj/provider-neo4j"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("neo4j.headyj.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		database.Configure,
		user.Configure,
		role.Configure,
		grant.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
