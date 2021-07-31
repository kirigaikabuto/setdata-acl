package setdata_acl

import (
	"fmt"
	"github.com/google/uuid"
)

type UserRoleService interface {
	CreateUserRole(cmd *CreateUserRoleCommand) (*UserRole, error)
	GetUserRole(cmd *GetUserRoleCommand) (*UserRole, error)
	ListUserRole(cmd *ListUserRoleCommand) ([]UserRole, error)
	DeleteUserRole(cmd *DeleteUserRoleCommand) error
	GetUserRolePermissions(cmd *GetUserRolePermissionsCommand) (*GetUserRolePermissionsResponse, error)
}

type userRoleService struct {
	store         UserRoleStore
	rolePermStore RolePermissionStore
	permStore     PermissionStore
}

func NewUserRoleService(s UserRoleStore) UserRoleService {
	return &userRoleService{store: s}
}

func (u *userRoleService) CreateUserRole(cmd *CreateUserRoleCommand) (*UserRole, error) {
	userRole := &UserRole{Id: uuid.New().String()}
	userRole.RoleId = cmd.RoleId
	userRole.UserId = cmd.UserId
	return u.store.Create(userRole)
}

func (u *userRoleService) GetUserRole(cmd *GetUserRoleCommand) (*UserRole, error) {
	return u.store.Get(cmd.Id)
}

func (u *userRoleService) ListUserRole(cmd *ListUserRoleCommand) ([]UserRole, error) {
	return u.store.List(cmd.RoleId, cmd.UserId)
}

func (u *userRoleService) DeleteUserRole(cmd *DeleteUserRoleCommand) error {
	return u.store.Delete(cmd.Id)
}

func (u *userRoleService) GetUserRolePermissions(cmd *GetUserRolePermissionsCommand) (*GetUserRolePermissionsResponse, error) {
	response := &GetUserRolePermissionsResponse{UserId: cmd.UserId}
	userRoles, err := u.store.List("", cmd.UserId)
	if err != nil {
		return nil, err
	}
	roles := []RolePermissionsResponse{}
	fmt.Println("userRoles ->", userRoles)
	for _, v := range userRoles {
		role := RolePermissionsResponse{Role: v.RoleId}
		rolePerms, err := u.rolePermStore.List(v.RoleId, "")
		if err != nil {
			return nil, err
		}
		fmt.Println("role", role, "userRoles ->", rolePerms)
		perms := []Permission{}
		for _, r := range rolePerms {
			perm, err := u.permStore.Get(r.PermissionId)
			if err != nil {
				return nil, err
			}
			perms = append(perms, *perm)
		}
		role.Permissions = perms
		roles = append(roles, role)
	}
	response.RolesPermissions = roles
	return response, nil
}
