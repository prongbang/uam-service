package role

func ToRoleListMapper(data []Role) []*RoleData {
	list := []*RoleData{}
	for _, u := range data {
		list = append(list, &RoleData{Id: u.ID, Name: u.Name})
	}
	return list
}
