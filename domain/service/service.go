package service

import (
	"github.com/IQ-tech/go-crypto-layer/datacrypto"
	"github.com/diegoclair/bank-transfer/contract"
	"github.com/diegoclair/bank-transfer/domain/entity"
	"github.com/diegoclair/bank-transfer/util/config"
)

type Service struct {
	dm     contract.DataManager
	cfg    *config.EnvironmentVariables
	cipher datacrypto.Crypto
}

func New(dm contract.DataManager, cfg *config.EnvironmentVariables, cipher datacrypto.Crypto) *Service {
	svc := new(Service)
	svc.dm = dm
	svc.cfg = cfg
	svc.cipher = cipher

	return svc
}

type Manager interface {
	AuthService(svc *Service) AuthService
}

type PingService interface {
}

type AuthService interface {
	CreateUser(newUser entity.User) (err error)
	Login(email, password string) (retVal entity.Authentication, err error)
}

type serviceManager struct {
}

func NewServiceManager() Manager {
	return &serviceManager{}
}

func (s *serviceManager) AuthService(svc *Service) AuthService {
	return newAuthService(svc)
}
