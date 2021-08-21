package authroute

import (
	"context"
	"sync"

	"github.com/IQ-tech/go-mapper"
	"github.com/diegoclair/bank-transfer/application/rest/routeutils"
	"github.com/diegoclair/bank-transfer/application/rest/viewmodel"
	"github.com/diegoclair/bank-transfer/domain/service"

	"github.com/labstack/echo/v4"
)

var (
	instance *Controller
	once     sync.Once
)

type Controller struct {
	authService service.AuthService
	mapper      mapper.Mapper
}

func NewController(authService service.AuthService, mapper mapper.Mapper) *Controller {
	once.Do(func() {
		instance = &Controller{
			authService: authService,
			mapper:      mapper,
		}
	})
	return instance
}

func (s *Controller) handleLogin(c echo.Context) error {

	input := viewmodel.Login{}
	err := c.Bind(&input)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	appContext := context.Background()
	auth, err := s.authService.Login(appContext, input.DocumentNumber, input.Secret)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	return routeutils.ResponseAPIOK(c, auth)
}
