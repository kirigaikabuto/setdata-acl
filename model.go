package setdata_acl



type RolePermissions struct {
	Id           string `json:"id"`
	RoleId       string `json:"role_id"`
	PermissionId string `json:"permission_id"`
}


type UserRoles struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
	RoleId string `json:"role_id"`
}