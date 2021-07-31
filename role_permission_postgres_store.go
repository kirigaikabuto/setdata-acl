package setdata_acl

import (
	"database/sql"
	"log"
)

var rolePermissionPostgresQueries = []string{
	`create table if not exists role_permissions(
		id text,
		role_id text,
		permission_id text,
		primary key(id),
		constraint fk_role_id foreign key(role_id) references roles(id),
		constraint fk_permission_id foreign key(permission_id) references permissions(id)
	);`,
}

type rolePermissionStore struct {
	db *sql.DB
}

func NewPostgresRolePermissionStore(cfg PostgresConfig) (RolePermissionStore, error) {
	db, err := getDbConn(getConnString(cfg))
	if err != nil {
		return nil, err
	}
	for _, q := range rolePermissionPostgresQueries {
		_, err := db.Exec(q)
		if err != nil {
			log.Println(err)
		}
	}
	db.SetMaxOpenConns(10)
	store := &rolePermissionStore{db: db}
	return store, nil
}

func (r *rolePermissionStore) Create(rolePerm *RolePermission) (*RolePermission, error) {
	query := "insert into role_permissions (id, role_id, permission_id) values ($1, $2, $3)"
	result, err := r.db.Exec(query, rolePerm.Id, rolePerm.RoleId, rolePerm.PermissionId)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, ErrCreateRolePermissionUnknown
	}
	return rolePerm, nil
}

func (r *rolePermissionStore) Get(id string) (*RolePermission, error) {
	return nil, nil
}

func (r *rolePermissionStore) List(roleId, permissionId string) ([]RolePermission, error) {
	return nil, nil
}

func (r *rolePermissionStore) Delete(id string) error {
	return nil
}
