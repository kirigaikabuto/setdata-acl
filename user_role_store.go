package setdata_acl

type UserRoleStore interface {
	Create(userRole *UserRole) (*UserRole, error)
	Get(id string) (*UserRole, error)
	List(roleId, userId string) ([]UserRole, error)
	Delete(id string) error
}
