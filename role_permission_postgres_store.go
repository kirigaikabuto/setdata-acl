package setdata_acl

import (
	"database/sql"
	"log"
	"strconv"
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
	rolePermission := &RolePermission{}
	query := "select id, role_id, permission_id from role_permissions where id = $1 limit 1"
	err := r.db.QueryRow(query, id).Scan(&rolePermission.Id, &rolePermission.RoleId, &rolePermission.PermissionId)
	if err == sql.ErrNoRows {
		return nil, ErrRolePermissionNotFound
	} else if err != nil {
		return nil, err
	}
	return rolePermission, nil
}

func (r *rolePermissionStore) List(roleId, permissionId string) ([]RolePermission, error) {
	rolePermissions := []RolePermission{}
	query := "select id, role_id, permission_id from role_permissions "
	var values []interface{}
	cnt := 1
	if roleId != "" {
		query = query + "where role_id = $" + strconv.Itoa(cnt)
		values = append(values, roleId)
		cnt += 1
		if permissionId != "" {
			query = query + "and permission_id = $" + strconv.Itoa(cnt)
			values = append(values, permissionId)
			cnt += 1
		}
	} else if permissionId != "" {
		query = query + "where permission_id = $" + strconv.Itoa(cnt)
		values = append(values, permissionId)
		cnt += 1
	}
	rows, err := r.db.Query(query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		item := RolePermission{}
		err = rows.Scan(&item.Id, &item.RoleId, &item.PermissionId)
		if err != nil {
			return nil, err
		}
		rolePermissions = append(rolePermissions, item)
	}
	return rolePermissions, nil
}

func (r *rolePermissionStore) Delete(id string) error {
	query := "delete from role_permissions where id = $1"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n <= 0 {
		return ErrRolePermissionNotFound
	}
	return nil
}
