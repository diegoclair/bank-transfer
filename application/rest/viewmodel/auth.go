package viewmodel

import (
	"github.com/diegoclair/bank-transfer/util/validator"
	"github.com/diegoclair/go_utils-lib/v2/resterrors"
	"github.com/diegoclair/go_utils-lib/v2/validstruct"
)

type Login struct {
	DocumentNumber string `json:"document_number,omitempty" validate:"required,min=11,max=11"`
	Secret         string `json:"secret,omitempty" validate:"required,min=8"`
}

func (l *Login) Validate() error {

	l.DocumentNumber = validator.CleanNumber(l.DocumentNumber)
	err := validstruct.ValidateStruct(l)
	if err != nil {
		return err
	}

	validDocument := validator.IsValidCPF(l.DocumentNumber)
	if !validDocument {
		return resterrors.NewUnprocessableEntity("Invalid cpf document")
	}

	return nil
}
