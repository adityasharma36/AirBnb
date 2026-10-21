package router

import (
	"goMod/controllers"
	"goMod/middleware"
	"goMod/utils/helper"

	"github.com/go-chi/chi/v5"
)


type Router interface{
	Register(r chi.Router);
}
func SetUpRouter(UserRouter Router) *chi.Mux {

	chiRouter := chi.NewRouter();

	chiRouter.Use(middleware.RateLimtingMiddleware)

	chiRouter.Handle("/api/v1/products",helper.ProxyToService("https://fakestoreapi.com","/api/v1"))

	chiRouter.Handle("/api/v1/products/*",helper.ProxyToService("https://fakestoreapi.com","/api/v1",))

	chiRouter.Get("/ping",controllers.PingHandler)

	UserRouter.Register(chiRouter);
	
	return chiRouter
}