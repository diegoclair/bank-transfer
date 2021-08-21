package service

import (
	"context"

	"github.com/diegoclair/bank-transfer/domain/entity"
)

type authService struct {
	svc *Service
}

func newAuthService(svc *Service) AuthService {
	return &authService{
		svc: svc,
	}
}

func (s *authService) Login(appContext context.Context, email, password string) (retVal entity.Authentication, err error) {

	return retVal, nil
}
