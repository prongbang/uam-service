package user_role

type Validate interface {
}

type validate struct {
}

func NewValidate() Validate {
	return &validate{}
}
