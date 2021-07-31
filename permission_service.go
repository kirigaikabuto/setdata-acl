package setdata_acl

import (
	"github.com/google/uuid"
	setdata_common "github.com/kirigaikabuto/setdata-common"
)

type PermissionService interface {
	CreatePermission(cmd *CreatePermissionCommand) (*Permission, error)
	GetPermission(cmd *GetPermissionCommand) (*Permission, error)
	DeletePermission(cmd *DeletePermissionCommand) error
	ListPermission(cmd *ListPermissionCommand) ([]Permission, error)
}

type permissionService struct {
	store PermissionStore
}

func NewPermissionService(s PermissionStore) PermissionService {
	return &permissionService{s}
}

func (p *permissionService) CreatePermission(cmd *CreatePermissionCommand) (*Permission, error) {
	perm := &Permission{}
	perm.Id = uuid.New().String()
	action, err := setdata_common.GetAclAction(cmd.Action)
	if err != nil {
		return nil, err
	}
	perm.Action = action
	resource, err := setdata_common.GetAclResource(cmd.Resource)
	if err != nil {
		return nil, err
	}
	perm.Resource = resource
	return p.store.Create(perm)
}

func (p *permissionService) GetPermission(cmd *GetPermissionCommand) (*Permission, error) {
	return p.store.Get(cmd.Id)
}

func (p *permissionService) DeletePermission(cmd *DeletePermissionCommand) error {
	return p.store.Delete(cmd.Id)
}

func (p *permissionService) ListPermission(cmd *ListPermissionCommand) ([]Permission, error) {
	_, err := setdata_common.GetAclResource(cmd.Resource)
	if err != nil {
		return nil, err
	}
	_, err = setdata_common.GetAclAction(cmd.Action)
	if err != nil {
		return nil, err
	}
	return p.store.List(cmd.Resource, cmd.Action)
}
