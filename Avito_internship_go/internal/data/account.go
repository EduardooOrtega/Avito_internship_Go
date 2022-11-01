package data

import (
	"database/sql"
	"errors"
	"fmt"
)

func (a AccountModel) Create() {
	query := `
		INSERT INTO account (account_cash, account_reserved_cash)
		VALUES (0, 0)`

	a.DB.QueryRow(query).Scan()
}

func (a AccountModel) CreateId(id int64) {
	query := `
		INSERT INTO account (account_id, account_cash, account_reserved_cash)
		VALUES ($1, 0, 0)`

	a.DB.QueryRow(query, id).Scan()
}

func (a AccountModel) Get(id int64) (*Account, error) {
	if id < 1 {
		return nil, errors.New("incorrect id")
	}

	query := `
			SELECT account_id, account_cash, account_reserved_cash
			FROM account
			WHERE account_id = $1
		`

	var account Account
	err := a.DB.QueryRow(query, id).Scan(
		&account.AccountId,
		&account.AccountCash,
		&account.AccountReservedCash,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("")
		default:
			return nil, err
		}
	}

	return &account, nil
}

func (a AccountModel) Update(account *Account) error {

	query := `
		UPDATE account
		SET account_cash = $2
		WHERE account_id = $1;
`
	args := []interface{}{
		account.AccountId,
		account.AccountCash,
		account.AccountReservedCash,
	}

	return a.DB.QueryRow(query, args...).Scan()
}

func (a AccountModel) UpdateFull(account *Account) error {
	fmt.Println("UpdateFull")
	//accountCashTemp, err := strconv.ParseInt(account.AccountCash, 10, 64)
	query := `
		UPDATE account
		SET account_cash = $2, account_reserved_cash = $3
		WHERE account_id = $1;
`
	args := []interface{}{
		account.AccountId,
		account.AccountCash,
		account.AccountReservedCash,
	}

	return a.DB.QueryRow(query, args...).Scan()
}
