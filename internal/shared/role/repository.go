package role

type Repository interface {
	Count() int64
	GetList(filter Filter) []Role
	GetListByUnderLevel(level int) []Role
	GetListByUnderRoles(roles []string) []Role
	GetById(id string) Role
	GetByName(name string) Role
	Add(data *Role) error
	Update(data *Role) error
	Delete(id string) error
}

type repository struct {
	Ds DataSource
}

func (r *repository) Count() int64 {
	return r.Ds.Count()
}

func (r *repository) GetList(filter Filter) []Role {
	return r.Ds.GetList(filter)
}

func (r *repository) GetListByUnderLevel(level int) []Role {
	return r.Ds.GetListByUnderLevel(level)
}

func (r *repository) GetListByUnderRoles(roles []string) []Role {
	return r.Ds.GetListByUnderRoles(roles)
}

func (r *repository) GetById(id string) Role {
	return r.Ds.GetById(id)
}

func (r *repository) GetByName(name string) Role {
	return r.Ds.GetByName(name)
}

func (r *repository) Add(data *Role) error {
	return r.Ds.Add(data)
}

func (r *repository) Update(data *Role) error {
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
