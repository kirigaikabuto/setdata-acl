package setdata_acl

type PermissionStore interface {
	Create(perm *Permission) (*Permission, error)
	Delete(id string) error
	Get(id string) (*Permission, error)
	List() ([]Permission, error)
}
