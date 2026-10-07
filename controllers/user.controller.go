package controllers

import (
	"goMod/service"
	"net/http"
)


type UserController struct{
	UserService service.UserService
}

func NewUserController(_userService service.UserService) *UserController{
	return &UserController{
		UserService: _userService,
	}
}

func (u * UserController) GetUser(w http.ResponseWriter, r * http.Request){

	w.Write([]byte("this is User Contoller"))
	
	u.UserService.GetUser();


}