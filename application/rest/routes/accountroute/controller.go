package accountroute

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
	accountService service.AccountService
	mapper         mapper.Mapper
}

func NewController(accountService service.AccountService, mapper mapper.Mapper) *Controller {
	once.Do(func() {
		instance = &Controller{
			accountService: accountService,
			mapper:         mapper,
		}
	})
	return instance
}

func (s *Controller) handleAddAccount(c echo.Context) error {

	input := viewmodel.AddAccount{}
	err := c.Bind(&input)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	err = input.Validate()
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	account := entity.Account{
		Name:   input.Name,
		CPF:    input.CPF,
		Secret: input.Secret,
	}

	err = s.accountService.CreateAccount(account)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}
	return routeutils.ResponseCreated(c)
}

func (s *Controller) handleGetAccounts(c echo.Context) error {

	accounts, err := s.accountService.GetAccounts()
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	response := []viewmodel.Account{}
	err = s.mapper.From(accounts).To(&response)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	return routeutils.ResponseAPIOK(c, response)
}

func (s *Controller) handleGetAccountByID(c echo.Context) error {

	accountUUID, err := routeutils.GetAndValidateParam(c, "account_id", "Invalid account_id")
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	account, err := s.accountService.GetAccountByUUID(accountUUID)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	response := viewmodel.Account{}
	err = s.mapper.From(account).To(&response)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	return routeutils.ResponseAPIOK(c, response)
}

func (s *Controller) handleGetAccountBalanceByID(c echo.Context) error {

	accountUUID, err := routeutils.GetAndValidateParam(c, "account_id", "Invalid account_id")
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	account, err := s.accountService.GetAccountByUUID(accountUUID)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	response := viewmodel.Account{
		Balance: account.Balance,
	}

	return routeutils.ResponseAPIOK(c, response)
}
