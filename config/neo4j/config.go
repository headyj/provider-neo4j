package neo4j

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("neo4j_database", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
	})
	p.AddResourceConfigurator("neo4j_role", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
	})
	p.AddResourceConfigurator("neo4j_user", func(r *config.Resource) {
		r.References["roles"] = config.Reference{
			TerraformName: "neo4j_role",
			RefFieldName:      "RoleRefs",
			SelectorFieldName: "RoleRefsSelector",
		}
	})
	p.AddResourceConfigurator("neo4j_grant", func(r *config.Resource) {
		r.References["role"] = config.Reference{
			TerraformName: "neo4j_role",
		}
	})
}
