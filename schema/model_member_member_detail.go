package schema

type MemberDetails struct {
	ID string `json:"id,omitempty"`

	ObjectID string `json:"objectId,omitempty"`

	GivenName string `json:"givenName"`

	Surname string `json:"surname"`

	Email string `json:"email,omitempty"`

	UserPrincipalName string `json:"userPrincipalName,omitempty"`

	TenantID string `json:"tenantId"`

	UserRole string `json:"userRole,omitempty"`
}
