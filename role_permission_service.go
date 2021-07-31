package setdata_acl

import "github.com/google/uuid"

type RolePermissionService interface {
	CreateRolePermission(cmd *CreateRolePermissionCommand) (*RolePermission, error)
	GetRolePermission(cmd *GetRolePermissionCommand) (*RolePermission, error)
	ListRolePermission(cmd *ListRolePermissionCommand) ([]RolePermission, error)
	DeleteRolePermission(cmd *DeleteRolePermissionCommand) error
}

type rolePermissionService struct {
	store RolePermissionStore
}

func NewRolePermissionService(s RolePermissionStore) RolePermissionService {
	return &rolePermissionService{store: s}
}

func (r *rolePermissionService) CreateRolePermission(cmd *CreateRolePermissionCommand) (*RolePermission, error) {
	rolePermission := &RolePermission{Id: uuid.New().String()}
	rolePermission.PermissionId = cmd.PermissionId
	return r.store.Create(rolePermission)
}

func (r *rolePermissionService) GetRolePermission(cmd *GetRolePermissionCommand) (*RolePermission, error) {
	return r.store.Get(cmd.Id)
}

func (r *rolePermissionService) ListRolePermission(cmd *ListRolePermissionCommand) ([]RolePermission, error) {
	return r.store.List(cmd.RoleId, cmd.PermissionId)
}

func (r *rolePermissionService) DeleteRolePermission(cmd *DeleteRolePermissionCommand) error {
	return r.store.Delete(cmd.Id)
}
