package setdata_acl

type RolePermissionStore interface {
	Create(rolePerm *RolePermission) (*RolePermission, error)
	Get(id string) (*RolePermission, error)
	List(roleId, permissionId string) ([]RolePermission, error)
	Delete(id string) error
}
