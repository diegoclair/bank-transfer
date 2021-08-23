package viewmodel

import "github.com/diegoclair/go_utils-lib/v2/validstruct"

type AddAccount struct {
	Name string `json:"name,omitempty" validate:"required,min=3"`
	Login
}

func (u *AddAccount) Validate() error {
	err := u.Login.Validate()
	if err != nil {
		return err
	}

	err = validstruct.ValidateStruct(u)
	if err != nil {
		return err
	}
	return nil
}
