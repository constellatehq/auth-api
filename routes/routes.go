package routes

import (
	"github.com/constellatehq/auth-api/routes/auth"
)

func init() {
	auth.initAuthRoutes()
}