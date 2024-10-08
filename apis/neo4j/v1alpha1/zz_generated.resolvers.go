// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Grant.
func (mg *Grant) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Role),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RoleRef,
		Selector:     mg.Spec.ForProvider.RoleSelector,
		To: reference.To{
			List:    &RoleList{},
			Managed: &Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Role")
	}
	mg.Spec.ForProvider.Role = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Role),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.RoleRef,
		Selector:     mg.Spec.InitProvider.RoleSelector,
		To: reference.To{
			List:    &RoleList{},
			Managed: &Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Role")
	}
	mg.Spec.InitProvider.Role = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RoleRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this User.
func (mg *User) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Roles),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.RoleRefs,
		Selector:      mg.Spec.ForProvider.RoleRefsSelector,
		To: reference.To{
			List:    &RoleList{},
			Managed: &Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Roles")
	}
	mg.Spec.ForProvider.Roles = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.RoleRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.Roles),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.InitProvider.RoleRefs,
		Selector:      mg.Spec.InitProvider.RoleRefsSelector,
		To: reference.To{
			List:    &RoleList{},
			Managed: &Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Roles")
	}
	mg.Spec.InitProvider.Roles = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.RoleRefs = mrsp.ResolvedReferences

	return nil
}
