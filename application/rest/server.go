package rest

import (
	"fmt"
	"os"

	"github.com/diegoclair/bank-transfer/application/factory"
	"github.com/diegoclair/bank-transfer/application/rest/routes/authroute"
	"github.com/diegoclair/bank-transfer/application/rest/routes/pingroute"
	"github.com/diegoclair/go_utils-lib/v2/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	logger.Info(fmt.Sprintf("About to start the application on port: %s...", port))

	if err := server.Start(fmt.Sprintf(":%s", port)); err != nil {
		panic(err)
	}
}

func initServer() *echo.Echo {

	factory := factory.GetDomainServices()

	srv := echo.New()
	srv.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	pingController := pingroute.NewController()
	authController := authroute.NewController(factory.AuthService, factory.Mapper)

	pingRoute := pingroute.NewRouter(pingController, "ping")
	authRoute := authroute.NewRouter(authController, "auth")

	appRouter := &Router{}
	appRouter.addRouters(authRoute)
	appRouter.addRouters(pingRoute)

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
