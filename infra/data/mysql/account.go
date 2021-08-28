package mysql

import (
	"github.com/diegoclair/bank-transfer/domain/entity"
	"github.com/diegoclair/go_utils-lib/v2/mysqlutils"
)

type accountRepo struct {
	db connection
}

func newAccountRepo(db connection) *accountRepo {
	return &accountRepo{
		db: db,
	}
}

const querySelectBase string = `
		SELECT 
			ta.account_id,
			ta.account_uuid,
			ta.name,
			ta.cpf,
			ta.balance,
			ta.secret
		
		FROM tab_account 				ta
		`

func (r *accountRepo) parseAccount(row scanner) (account entity.Account, err error) {

	err = row.Scan(
		&account.ID,
		&account.UUID,
		&account.Name,
		&account.CPF,
		&account.Balance,
		&account.Secret,
	)

	if err != nil {
		return account, err
	}

	return account, nil
}
func (r *accountRepo) AddTransfer(transfer entity.Transfer) (err error) {
	query := `
		INSERT INTO tab_transfer (
			account_origin_id,
			account_destination_id,
			amount
		) 
		VALUES (?, ?, ?);
	`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		transfer.AccountOriginID,
		transfer.AccountDestinationID,
		transfer.Amount,
	)
	if err != nil {
		return mysqlutils.HandleMySQLError(err)
	}

	return nil
}

func (r *accountRepo) CreateAccount(account entity.Account) (err error) {
	query := `
		INSERT INTO tab_account (
			account_uuid,
			name,
			cpf,
			secret
		) 
		VALUES (?, ?, ?, ?);
	`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		account.UUID,
		account.Name,
		account.CPF,
		account.Secret,
	)
	if err != nil {
		return mysqlutils.HandleMySQLError(err)
	}

	return nil
}

func (r *accountRepo) GetAccountByDocument(encryptedCPF string) (account entity.Account, err error) {

	query := querySelectBase + `
		WHERE  	ta.cpf 	= ?
	`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return account, mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(encryptedCPF)
	if err != nil {
		return account, mysqlutils.HandleMySQLError(err)
	}

	account, err = r.parseAccount(row)
	if err != nil {
		return account, mysqlutils.HandleMySQLError(err)
	}

	return account, nil
}

func (r *accountRepo) GetAccounts() (accounts []entity.Account, err error) {

	query := querySelectBase

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return accounts, mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return accounts, mysqlutils.HandleMySQLError(err)
	}
	account := entity.Account{}
	for rows.Next() {
		account, err = r.parseAccount(rows)
		if err != nil {
			return accounts, mysqlutils.HandleMySQLError(err)
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (r *accountRepo) GetAccountByUUID(accountUUID string) (account entity.Account, err error) {

	query := querySelectBase + `
		WHERE ta.account_uuid = ?
	`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return account, mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(accountUUID)
	if err != nil {
		return account, mysqlutils.HandleMySQLError(err)
	}

	account, err = r.parseAccount(row)
	if err != nil {
		return account, mysqlutils.HandleMySQLError(err)
	}

	return account, nil
}

func (r *accountRepo) GetTransfersByAccountID(accountID int64) (transfers []entity.Transfer, err error) {
	query := ` 
		SELECT 
			tt.transfer_id,
			tt.account_origin_id,
			tt.account_destination_id,
			tt.amount,
			tt.created_at
		
		FROM 	tab_transfer 			tt
		WHERE	tt.account_origin_id 	= 	?
	`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return transfers, mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(accountID)
	if err != nil {
		return transfers, mysqlutils.HandleMySQLError(err)
	}

	transfer := entity.Transfer{}
	for rows.Next() {
		err = rows.Scan(
			&transfer.ID,
			&transfer.AccountOriginID,
			&transfer.AccountDestinationID,
			&transfer.Amount,
			&transfer.CreateAt,
		)
		if err != nil {
			return transfers, err
		}

		transfers = append(transfers, transfer)
	}

	return transfers, nil
}

func (r *accountRepo) UpdateAccountBalance(account entity.Account) (err error) {

	query := `
		UPDATE 	tab_account
		
		SET 	balance 	= ?

		WHERE  	account_id 	= ?
	`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(account.Balance, account.ID)
	if err != nil {
		return mysqlutils.HandleMySQLError(err)
	}
	return nil
}
