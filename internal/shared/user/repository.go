package user

type Repository interface {
	Count(params Params) int64
	CountByUnderUserId(userId string, params Params) int64
	GetList(params Params) []User
	GetById(id string) User
	Add(data *User) error
	Update(data *User) error
	UpdatePassword(userId string, password string) error
	UpdateLastLogin(userId string) error
	Delete(id string) error
	IsUserUnder(userId1 string, userId2 string) bool
}

type repository struct {
	Ds DataSource
}

func (r *repository) Count(params Params) int64 {
	//TODO implement me
	panic("implement me")
}

func (r *repository) CountByUnderUserId(userId string, params Params) int64 {
	//TODO implement me
	panic("implement me")
}

func (r *repository) GetList(params Params) []User {
	//TODO implement me
	panic("implement me")
}

func (r *repository) GetById(id string) User {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Add(data *User) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Update(data *User) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) UpdatePassword(userId string, password string) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) UpdateLastLogin(userId string) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) IsUserUnder(userId1 string, userId2 string) bool {
	//TODO implement me
	panic("implement me")
}

func NewRepository(ds DataSource) Repository {
	return &repository{
		Ds: ds,
	}
}
