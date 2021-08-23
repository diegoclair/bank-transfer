package rest

import (
	"fmt"
	"os"

	"github.com/diegoclair/bank-transfer/application/factory"
	"github.com/diegoclair/bank-transfer/application/rest/routes/accountroute"
	"github.com/diegoclair/bank-transfer/application/rest/routes/authroute"
	"github.com/diegoclair/bank-transfer/application/rest/routes/pingroute"
	"github.com/diegoclair/bank-transfer/application/rest/routes/transferroute"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

// IRouter interface for routers
type IRouter interface {
	RegisterRoutes(appGroup, privateGroup *echo.Group)
}

type Router struct {
	routers []IRouter
}

func StartRestServer() {
	server := initServer()

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	log.Info(fmt.Sprintf("About to start the application on port: %s...", port))

	if err := server.Start(fmt.Sprintf(":%s", port)); err != nil {
		panic(err)
	}
}

func initServer() *echo.Echo {

	factory := factory.GetDomainServices()

	srv := echo.New()
	srv.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	pingController := pingroute.NewController()
	accountController := accountroute.NewController(factory.AccountService, factory.Mapper)
	authController := authroute.NewController(factory.AuthService, factory.Mapper)
	transferController := transferroute.NewController(factory.TransferService, factory.Mapper)

	pingRoute := pingroute.NewRouter(pingController, "ping")
	accountRoute := accountroute.NewRouter(accountController, "accounts")
	authRoute := authroute.NewRouter(authController, "auth")
	transferRoute := transferroute.NewRouter(transferController, "transfers")

	appRouter := &Router{}
	appRouter.addRouters(accountRoute)
	appRouter.addRouters(authRoute)
	appRouter.addRouters(pingRoute)
	appRouter.addRouters(transferRoute)

	return appRouter.registerAppRouters(srv)
}

func (r *Router) addRouters(router IRouter) {
	r.routers = append(r.routers, router)
}

func (r *Router) registerAppRouters(srv *echo.Echo) *echo.Echo {

	appGroup := srv.Group("/")
	privateGroup := appGroup.Group("")

	for _, appRouter := range r.routers {
		appRouter.RegisterRoutes(appGroup, privateGroup)
	}

	return srv
}
