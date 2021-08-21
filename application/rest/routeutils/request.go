package routeutils

import (
	"context"

	"github.com/diegoclair/bank-transfer/infra/auth"
	"github.com/labstack/echo/v4"
)

// GetContext returns a fulled appcontext
func GetContext(ctx echo.Context) (appContext context.Context) {
	appContext = context.WithValue(context.Background(), auth.UserIDKey, ctx.Get(auth.UserIDKey.String()))
	return appContext
}
