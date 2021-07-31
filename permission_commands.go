package setdata_acl

type CreatePermissionCommand struct {
	Resource string `json:"resource"`
	Action   string `json:"action"`
}

func (cmd *CreatePermissionCommand) Exec(service interface{}) (interface{}, error) {
	return service.(PermissionService).CreatePermission(cmd)
}

type GetPermissionCommand struct {
	Id string `json:"id"`
}

func (cmd *GetPermissionCommand) Exec(service interface{}) (interface{}, error) {
	return service.(PermissionService).GetPermission(cmd)
}

type DeletePermissionCommand struct {
	Id string `json:"id"`
}

func (cmd *DeletePermissionCommand) Exec(service interface{}) (interface{}, error) {
	return nil, service.(PermissionService).DeletePermission(cmd)
}

type ListPermissionCommand struct {
	Resource string `json:"resource"`
	Action   string `json:"action"`
}

func (cmd *ListPermissionCommand) Exec(service interface{}) (interface{}, error) {
	return service.(PermissionService).ListPermission(cmd)
}
