// This file was auto-generated from our API Definition.

package permissions

import (
	polytomicgo "github.com/polytomic/polytomic-go"
)

type V2CreatePolicyRequest struct {
	Name           string                        `json:"name" url:"name"`
	OrganizationId *string                       `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	PolicyActions  []*polytomicgo.V2PolicyAction `json:"policy_actions,omitempty" url:"policy_actions,omitempty"`
}

type V2UpdatePolicyRequest struct {
	Name           string                        `json:"name" url:"name"`
	OrganizationId *string                       `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	PolicyActions  []*polytomicgo.V2PolicyAction `json:"policy_actions,omitempty" url:"policy_actions,omitempty"`
}
