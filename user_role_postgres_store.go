package setdata_acl

import (
	"database/sql"
	"log"
	"strconv"
)

var userRolePostgresQueries = []string{
	`create table if not exists user_roles(
		id text,
		role_id text,
		user_id text,
		primary key(id),
		constraint fk_role_id foreign key(role_id) references roles(id),
		constraint fk_user_id foreign key(user_id) references users(id)
	);`,
}

type userRoleStore struct {
	db *sql.DB
}

func NewPostgresUserRoleStore(cfg PostgresConfig) (UserRoleStore, error) {
	db, err := getDbConn(getConnString(cfg))
	if err != nil {
		return nil, err
	}
	for _, q := range userRolePostgresQueries {
		_, err := db.Exec(q)
		if err != nil {
			log.Println(err)
		}
	}
	db.SetMaxOpenConns(10)
	store := &userRoleStore{db: db}
	return store, nil
}

func (u *userRoleStore) Create(userRole *UserRole) (*UserRole, error) {
	query := "insert into user_roles (id, role_id, user_id) values ($1, $2, $3)"
	result, err := u.db.Exec(query, userRole.Id, userRole.RoleId, userRole.UserId)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, ErrCreateUserRoleUnknown
	}
	return userRole, nil
}

func (u *userRoleStore) Get(id string) (*UserRole, error) {
	userRole := &UserRole{}
	query := "select id, role_id, user_id from user_roles where id = $1 limit 1"
	err := u.db.QueryRow(query, id).Scan(&userRole.Id, &userRole.RoleId, &userRole.UserId)
	if err == sql.ErrNoRows {
		return nil, ErrUserRoleNotFound
	} else if err != nil {
		return nil, err
	}
	return userRole, nil
}

func (u *userRoleStore) List(roleId, userId string) ([]UserRole, error) {
	userRoles := []UserRole{}
	query := "select id, role_id, user_id from user_roles "
	var values []interface{}
	cnt := 1
	if userId != "" {
		query = query + "where user_id = $" + strconv.Itoa(cnt)
		values = append(values, userId)
		cnt += 1
		if roleId != "" {
			query = query + "and role_id = $" + strconv.Itoa(cnt)
			values = append(values, roleId)
			cnt += 1
		}
	} else if roleId != "" {
		query = query + "where role_id = $" + strconv.Itoa(cnt)
		values = append(values, roleId)
		cnt += 1
	}
	rows, err := u.db.Query(query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		item := UserRole{}
		err = rows.Scan(&item.Id, &item.RoleId, &item.UserId)
		if err != nil {
			return nil, err
		}
		userRoles = append(userRoles, item)
	}
	return userRoles, nil
}

func (u *userRoleStore) Delete(id string) error {
	query := "delete from user_roles where id = $1"
	result, err := u.db.Exec(query, id)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n <= 0 {
		return ErrUserRoleNotFound
	}
	return nil
}
