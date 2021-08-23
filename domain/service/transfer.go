package service

import (
	"github.com/diegoclair/bank-transfer/domain/entity"
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

func (s *transferService) CreateTransfer(transfer entity.Transfer) (err error) {

	log.Info("CreateTransfer: Process Started")
	defer log.Info("CreateTransfer: Process Finished")

	return nil
}

func (s *transferService) GetTransfers() (transfers []entity.Transfer, err error) {

	log.Info("GetTransfers: Process Started")
	defer log.Info("GetTransfers: Process Finished")

	return transfers, nil
}
