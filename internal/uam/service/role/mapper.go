package role

func ToRoleListMapper(data []Role) []*RoleResponse {
	list := []*RoleResponse{}
	for _, u := range data {
		list = append(list, &RoleResponse{Id: u.ID, Name: u.Name})
	}
	return list
}
