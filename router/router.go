package router

import (
	"goMod/controllers"
	"goMod/middleware"

	"github.com/go-chi/chi/v5"
)


type Router interface{
	Register(r chi.Router);
}
func SetUpRouter(UserRouter Router) *chi.Mux {

	chiRouter := chi.NewRouter();
	chiRouter.Use(middleware.RateLimtingMiddleware)

	chiRouter.Get("/ping",controllers.PingHandler)

	UserRouter.Register(chiRouter);
	
	return chiRouter
}