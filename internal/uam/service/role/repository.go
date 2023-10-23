package role

type Repository interface {
	Count(params Params) int64
	GetList(params Params) []Role
	GetListByUnderRoles(roles []string) []Role
	GetById(params ParamsGetById) Role
	GetByName(name string) Role
	GetListByUserId(userId string) []Role
	Add(data *CreateRole) error
	Update(data *UpdateRole) error
	Delete(id string) error
}

type repository struct {
	Ds DataSource
}

func (r *repository) Count(params Params) int64 {
	return r.Ds.Count(params)
}

func (r *repository) GetList(params Params) []Role {
	return r.Ds.GetList(params)
}

func (r *repository) GetListByUnderRoles(roles []string) []Role {
	return r.Ds.GetListByUnderRoles(roles)
}

func (r *repository) GetById(params ParamsGetById) Role {
	return r.Ds.GetById(params)
}

func (r *repository) GetByName(name string) Role {
	return r.Ds.GetByName(name)
}

func (r *repository) GetListByUserId(userId string) []Role {
	return r.Ds.GetListByUserId(userId)
}

func (r *repository) Add(data *CreateRole) error {
	return r.Ds.Add(data)
}

func (r *repository) Update(data *UpdateRole) error {
	return r.Ds.Update(data)
}

func (r *repository) Delete(id string) error {
	return r.Ds.Delete(id)
}

func NewRepository(ds DataSource) Repository {
	return &repository{
		Ds: ds,
	}
}
