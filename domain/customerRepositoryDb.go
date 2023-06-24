package domain

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/vietbui1502/RestAPIGolang/logger"
)

type CustomerRepositoryDb struct {
	db *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	customers := make([]Customer, 0)

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	err := d.db.Select(&customers, findAllSql)

	if err != nil {
		logger.Error("Error when quering customer table" + err.Error())
		return nil, err
	}

	return customers, nil
}

func (d CustomerRepositoryDb) FindCustomerbyID(id string) (*Customer, error) {
	findSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	var c Customer

	err := d.db.Get(&c, findSql, id)

	if err != nil {
		logger.Error("Error when quering customer table" + err.Error())
		return nil, err
	}

	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sqlx.Open("mysql", "root:codecamp@tcp(localhost:3306)/banking")
	if err != nil {
		logger.Error(err.Error())
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{db: client}
}
