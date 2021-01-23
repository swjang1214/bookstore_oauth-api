package mysql_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client   *sql.DB
	username = "root"
	password = "ksj0208##"
	host     = "127.0.0.1"
	schema   = "users_db"
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host,
		schema,
	)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("mysql connection successfull")
}

/*
func Get(id interface{}, sql string) (*sql.Row, *errors.RestError) {
	stmt, err := Client.Prepare(sql)
	if err != nil {
		logger.Error("error when trying to prepare Get statement", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(id)
	if result == nil {
		logger.Error(fmt.Sprintf("error when trying to Query by %s", id), nil)
		return nil, errors.NewInternalServerError("database error")
	}
	return result, nil
}

func Create(data interface{}, sql string) (sql.Result, *errors.RestError) {
	stmt, err := Client.Prepare(sql)
	if err != nil {
		logger.Error("error when trying to prepare Create statement", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result, err := stmt.Exec(token.AccessToken, token.UserId, token.ClientId, token.Expires)
	if err != nil {
		logger.Error("error when trying to save token", err)
		return errors.NewInternalServerError("database error")
	}
	return result, nil
}

update
stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		logger.Error("error when trying to update user", err)
		return errors.NewInternalServerError("database error")
	}

delete

stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID)
	if err != nil {
		logger.Error("error when trying to delete user", err)
		return errors.NewInternalServerError("database error")
	}


*/
