package grant

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("neo4j_grant", func(r *config.Resource) {
		r.References["role"] = config.Reference{
			TerraformName: "neo4j_role",
		}
	})
}
