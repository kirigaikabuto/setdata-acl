package setdata_acl

import setdata_common "github.com/kirigaikabuto/setdata-common"

type Permission struct {
	Id       string                     `json:"id"`
	Resource setdata_common.AclResource `json:"resource"`
	Action   setdata_common.AclAction   `json:"action"`
}
