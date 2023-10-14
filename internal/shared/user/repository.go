package user

type Repository interface {
	Count(params Params) int64
	CountByUnderUserId(userId string, params Params) int64
	GetList(params Params) []User
	GetListByUnderUserId(userId string, params Params) []User
	GetById(id string) User
	GetByEmail(email string) User
	GetByUsername(username string) User
	Add(data *CreateUser) error
	Update(data *UpdateUser) error
	UpdatePassword(userId string, password string) error
	Delete(id string) error
}

type repository struct {
	Ds DataSource
}

func (r *repository) Count(params Params) int64 {
	return r.Ds.Count(params)
}

func (r *repository) CountByUnderUserId(userId string, params Params) int64 {
	return r.Ds.CountByUnderUserId(userId, params)
}

func (r *repository) GetList(params Params) []User {
	return r.Ds.GetList(params)
}

func (r *repository) GetListByUnderUserId(userId string, params Params) []User {
	return r.Ds.GetListByUnderUserId(userId, params)
}

func (r *repository) GetById(id string) User {
	return r.Ds.GetById(id)
}

func (r *repository) GetByEmail(email string) User {
	return r.Ds.GetByEmail(email)
}

func (r *repository) GetByUsername(username string) User {
	return r.Ds.GetByUsername(username)
}

func (r *repository) Add(data *CreateUser) error {
	return r.Ds.Add(data)
}

func (r *repository) Update(data *UpdateUser) error {
	return r.Ds.Update(data)
}

func (r *repository) UpdatePassword(userId string, password string) error {
	return r.Ds.UpdatePassword(userId, password)
}

func (r *repository) Delete(id string) error {
	return r.Ds.Delete(id)
}

func NewRepository(ds DataSource) Repository {
	return &repository{
		Ds: ds,
	}
}
