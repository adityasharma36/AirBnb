package service

import db "goMod/db/repositories"


type UserService interface{
	GetUser() error
}

type UserServiceImpl struct{
	userRepository db.UserRepository
}

func NewService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	};
}

func (u *UserServiceImpl) GetUser() error{
	println("this is inside user Create Server");
	u.userRepository.Create();
	return nil;
}