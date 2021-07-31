package setdata_acl

import (
	"github.com/google/uuid"
)

type UserRoleService interface {
	CreateUserRole(cmd *CreateUserRoleCommand) (*UserRole, error)
	GetUserRole(cmd *GetUserRoleCommand) (*UserRole, error)
	ListUserRole(cmd *ListUserRoleCommand) ([]UserRole, error)
	DeleteUserRole(cmd *DeleteUserRoleCommand) error
}

type userRoleService struct {
	store UserRoleStore
}

func NewUserRoleService(s UserRoleStore) UserRoleService{
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
