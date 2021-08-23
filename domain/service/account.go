package service

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/diegoclair/bank-transfer/domain/entity"
	"github.com/diegoclair/bank-transfer/util/errors"
	"github.com/diegoclair/go_utils-lib/v2/resterrors"
	"github.com/labstack/gommon/log"
	"github.com/twinj/uuid"
)

type accountService struct {
	svc *Service
}

func newAccountService(svc *Service) AccountService {
	return &accountService{
		svc: svc,
	}
}

func (s *accountService) CreateAccount(account entity.Account) (err error) {

	log.Info("CreateAccount: Process Started")
	defer log.Info("CreateAccount: Process Finished")

	account.CPF, err = s.svc.cipher.Encrypt(account.CPF)
	if err != nil {
		log.Error("CreateAccount: ", err)
		return err
	}
	fmt.Println(s.svc.dm.MySQL())
	_, err = s.svc.dm.MySQL().Account().GetAccountByDocument(account.CPF)
	if err != nil && !errors.SQLNotFound(err.Error()) {
		log.Error("CreateAccount: ", err)
		return err
	} else if err == nil {
		log.Error("CreateAccount: The document number is already in use")
		return resterrors.NewConflictError("The cpf is already in use")
	}

	hasher := md5.New()
	hasher.Write([]byte(account.Secret))
	account.Secret = hex.EncodeToString(hasher.Sum(nil))

	account.UUID = uuid.NewV4().String()

	err = s.svc.dm.MySQL().Account().CreateAccount(account)
	if err != nil {
		log.Error("CreateAccount: ", err)
		return err
	}

	return nil
}
