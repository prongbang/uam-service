package schema

type Schema interface {
	Initial()
}

type schema struct {
	UserSchema     UserSchema
	RoleSchema     RoleSchema
	UserRoleSchema UserRoleSchema
	RBACSchema     RBACSchema
}

func (s *schema) Initial() {
	s.UserSchema.Initial()
	s.RoleSchema.Initial()
	s.UserRoleSchema.Initial()
	s.RBACSchema.Initial()
}

func NewSchema(
	userSchema UserSchema,
	roleSchema RoleSchema,
	userRoleSchema UserRoleSchema,
	rbacSchema RBACSchema,
) Schema {
	return &schema{
		UserSchema:     userSchema,
		RoleSchema:     roleSchema,
		UserRoleSchema: userRoleSchema,
		RBACSchema:     rbacSchema,
	}
}
