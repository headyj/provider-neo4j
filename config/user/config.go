package user

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("neo4j_user", func(r *config.Resource) {
		r.References["roles"] = config.Reference{
			TerraformName:     "neo4j_role",
			RefFieldName:      "RoleRefs",
			SelectorFieldName: "RoleRefsSelector",
		}
	})
}
