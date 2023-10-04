package uam

type APIValidate interface {
}

type apiValidate struct {
}

func NewValidate() APIValidate {
	return &apiValidate{}
}
