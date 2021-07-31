package setdata_acl

import (
	"database/sql"
	"log"
	"strconv"
)

var permissionPostgresQueries = []string{
	`create table if not exists permissions(
		id text,
		resource text,
		action text,
		primary key(id)
	);`,
}

type permissionStore struct {
	db *sql.DB
}

func NewPostgresPermissionStore(cfg PostgresConfig) (PermissionStore, error) {
	db, err := getDbConn(getConnString(cfg))
	if err != nil {
		return nil, err
	}
	for _, q := range permissionPostgresQueries {
		_, err := db.Exec(q)
		if err != nil {
			log.Println(err)
		}
	}
	db.SetMaxOpenConns(10)
	store := &permissionStore{db: db}
	return store, nil
}

func (p *permissionStore) Create(perm *Permission) (*Permission, error) {
	query := "insert into permissions (id, resource, action) values ($1, $2, $3)"
	result, err := p.db.Exec(query, perm.Id, perm.Resource, perm.Action)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, ErrCreatePermissionUnknown
	}
	return perm, nil
}

func (p *permissionStore) Delete(id string) error {
	query := "delete from permissions where id = $1"
	result, err := p.db.Exec(query, id)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n <= 0 {
		return ErrPermissionNotFound
	}
	return nil
}

func (p *permissionStore) Get(id string) (*Permission, error) {
	permission := &Permission{}
	query := "select id, resource, action from permissions where id = $1 limit 1"
	err := p.db.QueryRow(query, id).Scan(&permission.Id, &permission.Resource, &permission.Action)
	if err == sql.ErrNoRows {
		return nil, ErrPermissionNotFound
	} else if err != nil {
		return nil, err
	}
	return permission, nil
}

func (p *permissionStore) List(resource, action string) ([]Permission, error) {
	perms := []Permission{}
	query := "select id, resource, action from permissions "
	var values []interface{}
	cnt := 1
	if resource != "" {
		query = query + "where resource = $" + strconv.Itoa(cnt)
		values = append(values, resource)
		cnt += 1
		if action != "" {
			query = query + "and action = $" + strconv.Itoa(cnt)
			values = append(values, action)
			cnt += 1
		}
	} else if action != "" {
		query = query + "where action = $" + strconv.Itoa(cnt)
		values = append(values, action)
		cnt += 1
	}
	rows, err := p.db.Query(query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		item := Permission{}
		err = rows.Scan(&item.Id, &item.Resource, &item.Action)
		if err != nil {
			return nil, err
		}
		perms = append(perms, item)
	}
	return perms, nil
}
