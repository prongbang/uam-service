package schema

type Schema interface {
	Initial()
}

type schema struct {
	UserSchema     UserSchema
	RoleSchema     RoleSchema
	UserRoleSchema UserRoleSchema
}

func (s *schema) Initial() {
	s.UserSchema.Initial()
	s.RoleSchema.Initial()
	s.UserRoleSchema.Initial()
}

func NewSchema(
	userSchema UserSchema,
	roleSchema RoleSchema,
	userRoleSchema UserRoleSchema,
) Schema {
	return &schema{
		UserSchema:     userSchema,
		RoleSchema:     roleSchema,
		UserRoleSchema: userRoleSchema,
	}
}
