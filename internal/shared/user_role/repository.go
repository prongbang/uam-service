package user_role

type Repository interface {
}

type repository struct {
	Ds DataSource
}

func NewRepository(ds DataSource) Repository {
	return &repository{
		Ds: ds,
	}
}
