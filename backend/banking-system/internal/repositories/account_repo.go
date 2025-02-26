package repositories

import (
	"database/sql"
	"fmt"
	"northern-bank/internal/entities"
	"northern-bank/pkg"
)

type AccountRepositoryDB struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) AccountRepository {
	return &AccountRepositoryDB{db: db}
}

func (a *AccountRepositoryDB) CreateAccount(acc_req *entities.Account) (*entities.Account, error) {
	var account entities.Account
	row := a.db.QueryRow(`
		insert into accounts (user_id, acc_no, bank_code, balance)
		values ($1, $2, $3, $4)
		RETURNING acc_id, user_id, acc_no, bank_code, balance;
	`, acc_req.UserID, acc_req.AccNo, acc_req.BankCode, acc_req.Balance)

	err := row.Scan(&account.AccID, &account.UserID, &account.AccNo, &account.BankCode, &account.Balance)
	if err != nil {
		fmt.Printf("Error on %v => %v\n", pkg.GetCallerInfo(), err)
		return nil, err
	}

	return &account, nil
}
