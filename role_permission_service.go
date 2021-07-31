package setdata_acl

import "github.com/google/uuid"

type RolePermissionService interface {
	CreateRolePermission(cmd *CreateRolePermissionCommand) (*RolePermission, error)
	GetRolePermission(cmd *GetRolePermissionCommand) (*RolePermission, error)
	ListRolePermission(cmd *ListRolePermissionCommand) ([]RolePermission, error)
	DeleteRolePermission(cmd *DeleteRolePermissionCommand) error
}

type rolePermissionService struct {
	store           RolePermissionStore
	roleStore       RoleStore
	permissionStore PermissionStore
}

func NewRolePermissionService(s RolePermissionStore, r RoleStore, p PermissionStore) RolePermissionService {
	return &rolePermissionService{store: s, roleStore: r, permissionStore: p}
}

func (r *rolePermissionService) CreateRolePermission(cmd *CreateRolePermissionCommand) (*RolePermission, error) {
	rolePermission := &RolePermission{Id: uuid.New().String()}
	_, err := r.roleStore.Get(cmd.RoleId)
	if err != nil {
		return nil, err
	}
	rolePermission.RoleId = cmd.RoleId
	_, err = r.permissionStore.Get(cmd.PermissionId)
	if err != nil {
		return nil, err
	}
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
