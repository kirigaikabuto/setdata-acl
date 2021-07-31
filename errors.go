package setdata_acl

import (
	"errors"
	com "github.com/kirigaikabuto/setdata-common"
)

var (
	ErrCreateRoleUnknown       = com.NewMiddleError(errors.New("could not create role:unknown error"), 500, 50)
	ErrRoleNotFound            = com.NewMiddleError(errors.New("role not found"), 404, 51)
	ErrRoleIdNotProvided       = com.NewMiddleError(errors.New("role id is not provided"), 400, 52)
	ErrCreatePermissionUnknown = com.NewMiddleError(errors.New("could not create permission:unknown error"), 500, 53)
	ErrPermissionNotFound            = com.NewMiddleError(errors.New("permission not found"), 404, 54)
)
