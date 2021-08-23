package accountroute

import (
	"github.com/labstack/echo/v4"
)

const (
	rootRoute = "/"
)

type UserRouter struct {
	ctrl      *Controller
	routeName string
}

func NewRouter(ctrl *Controller, routeName string) *UserRouter {
	return &UserRouter{
		ctrl:      ctrl,
		routeName: routeName,
	}
}

func (r *UserRouter) RegisterRoutes(appGroup, privateGroup *echo.Group) {
	router := appGroup.Group(r.routeName)
	router.POST(rootRoute, r.ctrl.handleAddAccount)
}
