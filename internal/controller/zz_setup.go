// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	database "github.com/headyj/provider-neo4j/internal/controller/neo4j/database"
	grant "github.com/headyj/provider-neo4j/internal/controller/neo4j/grant"
	role "github.com/headyj/provider-neo4j/internal/controller/neo4j/role"
	user "github.com/headyj/provider-neo4j/internal/controller/neo4j/user"
	providerconfig "github.com/headyj/provider-neo4j/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		database.Setup,
		grant.Setup,
		role.Setup,
		user.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
