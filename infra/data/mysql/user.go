package mysql

import (
	"github.com/diegoclair/bank-transfer/domain/entity"
	"github.com/diegoclair/go_utils-lib/v2/mysqlutils"
)

type userRepo struct {
	db connection
}

func newUserRepo(db connection) *userRepo {
	return &userRepo{
		db: db,
	}
}

const querySelectBase string = `
		SELECT 
			tu.user_id,
			tu.user_uuid,
			tu.name,
			tu.document_number,
			tu.password
		
		FROM tab_user 				tu
		`

func (r *userRepo) parseUser(row scanner) (retVal entity.User, err error) {

	err = row.Scan(
		&retVal.ID,
		&retVal.UUID,
		&retVal.Name,
		&retVal.DocumentNumber,
		&retVal.Password,
	)

	if err != nil {
		return retVal, err
	}

	return retVal, nil
}

func (r *userRepo) GetUserByDocument(encryptedDocumentNumber string) (user entity.User, err error) {

	query := querySelectBase + `
		WHERE  	tu.document_number 	= ?
	`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return user, mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(encryptedDocumentNumber)
	if err != nil {
		return user, mysqlutils.HandleMySQLError(err)
	}

	user, err = r.parseUser(row)
	if err != nil {
		return user, mysqlutils.HandleMySQLError(err)
	}

	return user, nil
}
