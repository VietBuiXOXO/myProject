package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/vietbui1502/RestAPIGolang/logger"
)

type AccountRepositoryDb struct {
	db *sqlx.DB
}

func (repo AccountRepositoryDb) Save(a Account) (*Account, error) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"
	result, err := repo.db.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error when creating new account: " + err.Error())
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error when getting last insert of id: " + err.Error())
		return nil, err
	}
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{db: dbClient}
}
