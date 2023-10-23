package schema

type Schema interface {
	Initial()
}

type schema struct {
	UserSchema        UserSchema
	RoleSchema        RoleSchema
	UserRoleSchema    UserRoleSchema
	UserCreatorSchema UserCreatorSchema
	RBACSchema        RBACSchema
}

func (s *schema) Initial() {
	s.UserSchema.Initial()
	s.RoleSchema.Initial()
	s.UserRoleSchema.Initial()
	s.UserCreatorSchema.Initial()
	s.RBACSchema.Initial()
}

func NewSchema(
	userSchema UserSchema,
	roleSchema RoleSchema,
	userRoleSchema UserRoleSchema,
	userCreatorSchema UserCreatorSchema,
	rbacSchema RBACSchema,
) Schema {
	return &schema{
		UserSchema:        userSchema,
		RoleSchema:        roleSchema,
		UserRoleSchema:    userRoleSchema,
		UserCreatorSchema: userCreatorSchema,
		RBACSchema:        rbacSchema,
	}
}
