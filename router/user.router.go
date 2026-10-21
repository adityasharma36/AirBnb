package router

import (
	"goMod/controllers"
	"goMod/middleware"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct{
	UserController *controllers.UserController
}

func NewUserRouter(_userController *controllers.UserController) Router{

	return &UserRouter{
		UserController: _userController,
	}

}

func (ur * UserRouter) Register(r chi.Router){
	// r.Post("/signup",ur.UserController.RegisterUser)
	r.With(middleware.LoginMiddleware).Post("/login",ur.UserController.LoginUser)
	r.With(middleware.CreateUserMiddleware).Post("/create",ur.UserController.CreateUser)
	r.With(middleware.JWTMiddlewareAuth).Get("/profile",ur.UserController.GetUser);
	

}