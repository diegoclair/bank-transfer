package authroute

import (
	"sync"

	"github.com/IQ-tech/go-mapper"
	"github.com/diegoclair/bank-transfer/application/rest/routeutils"
	"github.com/diegoclair/bank-transfer/application/rest/viewmodel"
	"github.com/diegoclair/bank-transfer/domain/entity"
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
	err = input.Validate()
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	auth, err := s.authService.Login(input.DocumentNumber, input.Secret)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	return routeutils.ResponseAPIOK(c, auth)
}

func (s *Controller) handleSignup(c echo.Context) error {

	input := viewmodel.Signup{}
	err := c.Bind(&input)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}
	err = input.Validate()
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	user := entity.User{}
	err = s.mapper.From(input).To(&user)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}
	user.DocumentNumber = input.DocumentNumber
	user.Password = input.Secret

	err = s.authService.CreateUser(user)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	return routeutils.ResponseCreated(c)
}
