package setdata_acl

type CreateRolePermissionCommand struct {
	RoleId       string `json:"role_id"`
	PermissionId string `json:"permission_id"`
}

func (cmd *CreateRolePermissionCommand) Exec(service interface{}) (interface{}, error) {
	return service.(RolePermissionService).CreateRolePermission(cmd)
}

type GetRolePermissionCommand struct {
	Id string `json:"id"`
}

func (cmd *GetRolePermissionCommand) Exec(service interface{}) (interface{}, error) {
	return service.(RolePermissionService).GetRolePermission(cmd)
}

type ListRolePermissionCommand struct {
	RoleId       string `json:"role_id"`
	PermissionId string `json:"permission_id"`
}

func (cmd *ListRolePermissionCommand) Exec(service interface{}) (interface{}, error) {
	return service.(RolePermissionService).ListRolePermission(cmd)
}

type DeleteRolePermissionCommand struct {
	Id string `json:"id"`
}

func (cmd *DeleteRolePermissionCommand) Exec(service interface{}) (interface{}, error) {
	return nil, service.(RolePermissionService).DeleteRolePermission(cmd)
}
