package dto

import "goMod/config"

type payloadKey string
type userCreateKey string


var PayloadKey payloadKey = payloadKey(config.GetString("PayloadKey","loadKey"))
var UserCreateKey userCreateKey = userCreateKey(config.GetString("UserCreateKey","loadKey"))
type LoginUserRequestDTO struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
}

type CreateUserRequestDTO struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=8"`
    Username string `json:"username" validate:"required"`
}
