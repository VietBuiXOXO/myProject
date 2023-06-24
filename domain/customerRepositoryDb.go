package domain

import (
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

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{db: dbClient}
}
