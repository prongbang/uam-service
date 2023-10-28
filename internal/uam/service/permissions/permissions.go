package permissions

const (
	All                = "all"
	Read               = "read"
	ReadMe             = "readme"
	Create             = "create"
	Update             = "update"
	UpdateMe           = "updateme"
	Delete             = "delete"
	Grpc               = "gRPC"
	UamPermissionUsers = "uam.permission.users"
	UamPermissionRoles = "uam.permission.roles"
)

func IsAllowed(results ...bool) bool {
	for _, v := range results {
		if v {
			return true
		}
	}
	return false
}
