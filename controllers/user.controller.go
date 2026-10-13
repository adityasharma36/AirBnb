package controllers

import (
	"database/sql"
	
	"net/http"
	"strconv"

	"goMod/dto"
	"goMod/service"
	"goMod/utils/helper"
)

type UserController struct {
    UserService service.UserService
}

func NewUserController(
    userService service.UserService,
) *UserController {
    return &UserController{
        UserService: userService,
    }
}

func (u *UserController) CreateUser(
    w http.ResponseWriter,
    r *http.Request,
) {
    value := r.Context().Value(dto.UserCreateKey)
    user, ok := value.(*dto.CreateUserRequestDTO)

    if !ok || user == nil {
        helper.WriteJsonErrResponse(
            w,
            http.StatusBadRequest,
            "Invalid user request",
            nil,
        )
        return
    }

    created, err := u.UserService.CreateUser(
        user.Username,
        user.Email,
        user.Password,
    )

    if err != nil {
        if err == service.ErrUserAlreadyExists {
            helper.WriteJsonErrResponse(
                w,
                http.StatusConflict,
                "Email or username already exists",
                err,
            )
            return
        }

        helper.WriteJsonErrResponse(
            w,
            http.StatusInternalServerError,
            "Failed to create user",
            err,
        )
        return
    }

    helper.WriteJsonSucResp(
        w,
        http.StatusCreated,
        "User created successfully",
        created,
    )
}

func (u *UserController) LoginUser(
    w http.ResponseWriter,
    r *http.Request,
) {
    value := r.Context().Value(dto.PayloadKey)
    payload, ok := value.(*dto.LoginUserRequestDTO)

    if !ok || payload == nil {
        helper.WriteJsonErrResponse(
            w,
            http.StatusBadRequest,
            "Invalid login request",
            nil,
        )
        return
    }

    token, err := u.UserService.LoginUser(payload)
    if err != nil {
        helper.WriteJsonErrResponse(
            w,
            http.StatusUnauthorized,
            "Invalid email or password",
            err,
        )
        return
    }

    helper.WriteJsonSucResp(
        w,
        http.StatusOK,
        "Login successful",
        map[string]string{
            "token": token,
        },
    )
}

func (u *UserController) GetUser(
    w http.ResponseWriter,
    r *http.Request,
) {

    idValue := r.URL.Query().Get("id")

    id, err := strconv.ParseInt(idValue, 10, 64)
    if err != nil || id <= 0 {
        helper.WriteJsonErrResponse(
            w,
            http.StatusBadRequest,
            "Invalid user ID",
            err,
        )
        return
    }

    user, err := u.UserService.GetUser(id)
    if err != nil {
        if err == sql.ErrNoRows {
            helper.WriteJsonErrResponse(
                w,
                http.StatusNotFound,
                "User not found",
                err,
            )
            return
        }

        helper.WriteJsonErrResponse(
            w,
            http.StatusInternalServerError,
            "Failed to fetch user",
            err,
        )
        return
    }

    helper.WriteJsonSucResp(
        w,
        http.StatusOK,
        "User fetched successfully",
        user,
    )
}
