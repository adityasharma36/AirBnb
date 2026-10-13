package service

import (
	"errors"
	"fmt"

	"goMod/config"
	db "goMod/db/repositories"
	"goMod/dto"
	"goMod/utils/helper"

	"goMod/model"

	"github.com/golang-jwt/jwt/v5"
)

var ErrUserAlreadyExists = db.ErrUserAlreadyExists

type UserService interface {
    GetUser(id int64) (*model.User, error)
    CreateUser(username, email, password string) (*model.User, error)
    LoginUser(payload *dto.LoginUserRequestDTO) (string, error)
}

type UserServiceImpl struct {
    userRepository db.UserRepository
}

func NewUserService(
    userRepository db.UserRepository,
) UserService {
    return &UserServiceImpl{
        userRepository: userRepository,
    }
}

func (u *UserServiceImpl) GetUser(
    id int64,
) (*model.User, error) {

    user, err := u.userRepository.GetUserById(id)
    if err != nil {
        return nil, err
    }

    return user, nil
}

func (u *UserServiceImpl) CreateUser(
    username string,
    email string,
    password string,
) (*model.User, error) {

    hashedPassword, err := helper.HashPassword(password)
	fmt.Println("this is Create User Service , ", hashedPassword, err)
    if err != nil {
        return nil, fmt.Errorf("hash password: %w", err)
    }

    user, err := u.userRepository.Create(
        username,
        email,
        hashedPassword,
    )
	fmt.Println("this is inside user Service ", user, err)
    if err != nil {
        return nil, err
    }

    return user, nil
}

func (u *UserServiceImpl) LoginUser(
    payload *dto.LoginUserRequestDTO,
) (string, error) {

    user, err := u.userRepository.GetUserByEmail(payload.Email)
    if err != nil {
        return "", errors.New("invalid email or password")
    }

    passwordMatches := helper.CheckPasswordHash(
        payload.Password,
        user.Password,
    )

    if !passwordMatches {
        return "", errors.New("invalid email or password")
    }

    claims := jwt.MapClaims{
        "id":    user.Id,
        "email": user.Email,
    }

    token := jwt.NewWithClaims(
        jwt.SigningMethodHS256,
        claims,
    )

    jwtKey := []byte(
        config.GetString("JWT_STRING", "myString"),
    )

    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", fmt.Errorf("create JWT token: %w", err)
    }

    return tokenString, nil
}
