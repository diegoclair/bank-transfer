package service

import (
	"context"

	"github.com/diegoclair/bank-transfer/domain/entity"
	"github.com/diegoclair/bank-transfer/infra/auth"
	"github.com/labstack/gommon/log"
)

type transferService struct {
	svc *Service
}

func newTransferService(svc *Service) TransferService {
	return &transferService{
		svc: svc,
	}
}

func (s *transferService) CreateTransfer(appContext context.Context, transfer entity.Transfer) (err error) {

	log.Info("CreateTransfer: Process Started")
	defer log.Info("CreateTransfer: Process Finished")

	return nil
}

func (s *transferService) GetTransfers(appContext context.Context) (transfers []entity.Transfer, err error) {

	log.Info("GetTransfers: Process Started")
	defer log.Info("GetTransfers: Process Finished")
	loggedAccountUUID := appContext.Value(auth.AccountUUIDKey)

	account, err := s.svc.dm.MySQL().Account().GetAccountByUUID(loggedAccountUUID.(string))
	if err != nil {
		log.Error("GetTransfers: ", err)
		return transfers, err
	}

	transfers, err = s.svc.dm.MySQL().Account().GetTransfersByAccountID(account.ID)
	if err != nil {
		log.Error("GetTransfers: ", err)
		return transfers, err
	}

	return transfers, nil
}
