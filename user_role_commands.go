package setdata_acl

type CreateUserRoleCommand struct {
	RoleId string `json:"role_id"`
	UserId string `json:"user_id"`
}

func (cmd *CreateUserRoleCommand) Exec(service interface{}) (interface{}, error) {
	return service.(UserRoleService).CreateUserRole(cmd)
}

type GetUserRoleCommand struct {
	Id string `json:"id"`
}

func (cmd *GetUserRoleCommand) Exec(service interface{}) (interface{}, error) {
	return service.(UserRoleService).GetUserRole(cmd)
}

type ListUserRoleCommand struct {
	RoleId string `json:"role_id"`
	UserId string `json:"user_id"`
}

func (cmd *ListUserRoleCommand) Exec(service interface{}) (interface{}, error) {
	return service.(UserRoleService).ListUserRole(cmd)
}

type DeleteUserRoleCommand struct {
	Id string `json:"id"`
}

func (cmd *DeleteUserRoleCommand) Exec(service interface{}) (interface{}, error) {
	return nil, service.(UserRoleService).DeleteUserRole(cmd)
}

type GetUserRolePermissionsCommand struct {
	UserId string `json:"user_id"`
}

func (cmd *GetUserRolePermissionsCommand) Exec(service interface{}) (interface{}, error) {
	return service.(UserRoleService).GetUserRolePermissions(cmd)
}

type RolePermissionsResponse struct {
	Role        string       `json:"role_id"`
	Permissions []Permission `json:"permissions"`
}

type GetUserRolePermissionsResponse struct {
	UserId           string                    `json:"user_id"`
	RolesPermissions []RolePermissionsResponse `json:"roles_permissions"`
}
