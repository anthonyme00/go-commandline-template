package sql

import (
	"cltest/configs"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySQLRepository struct {
	db *sqlx.DB
}

func NewMySQLRepository() (*MySQLRepository, error) {
	return (&MySQLRepository{}).connect()
}

func (r *MySQLRepository) connect() (*MySQLRepository, error) {
	globalConfig := configs.Global
	connectionString := "%s:%s@tcp(%s:%s)/%s"
	connectionString = fmt.Sprintf(connectionString, globalConfig.DB_USER, globalConfig.DB_PASS, globalConfig.DB_HOST, globalConfig.DB_PORT, globalConfig.DB_NAME)
	db, err := sqlx.Open("mysql", connectionString)
	if err == nil {
		err = db.Ping()
	}
	return &MySQLRepository{
		db: db,
	}, err
}

func (r *MySQLRepository) GetOne(dest interface{}, query string, args ...interface{}) error {
	return r.db.Get(dest, query, args...)
}

func (r *MySQLRepository) GetAll(dest func() interface{}, query string, args ...interface{}) error {
	rows, err := r.db.Queryx(query, args...)
	if err != nil {
		return err
	}

	i := 0
	for rows.Next() {
		rows.StructScan(dest())
		if err != nil {
			return err
		}
		i++
	}
	return nil
}

func (r *MySQLRepository) Clean() {
	r.db.Close()
}
