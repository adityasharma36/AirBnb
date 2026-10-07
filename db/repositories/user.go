package db

import (
	"database/sql"
	"fmt"
	"goMod/model"
)


type UserRepository interface{
	GetUserById() (*model.User,error)
	Create() error
	GetAll() ([]*model.User,error)
	DeleteById(id int64) error
}

type UserRepositoryImpl struct{
	db *sql.DB
}
func NewUserRepository(_db *sql.DB) UserRepository{
	return &UserRepositoryImpl{
		db: _db,
	}
}

func (u* UserRepositoryImpl) GetAll() ([]*model.User,error){
	return nil,nil;
}
func (u* UserRepositoryImpl) DeleteById(id int64) error{
	return nil;
}
func (u* UserRepositoryImpl) Create() error{
	query:= "insert into user (username,email,password) values (?,? ,?)";
	row,err := u.db.Exec(query,"test","test@test.com","password123");

	if err != nil{
		fmt.Println("error occuring : ", err);
		return err;
	}
	rowAffected,rowErr := row.RowsAffected();


	if rowErr != nil{
		fmt.Println("Error Getting rows affected: ", rowErr);
		return err;
	}
	if rowAffected == 0{
		fmt.Println("row Affected Error",rowAffected);
		return err;
	}
	fmt.Println("User Created Successfully :  ", rowAffected);
	return nil;
}
func (u * UserRepositoryImpl) GetUserById() (*model.User,error) {

	query := "select id , username , email , created_at ,updated_at from user where id = ? "
	fmt.Println("this is inside User repositories");

	row := u.db.QueryRow(query,1)

	user := &model.User{}

	err :=row.Scan(&user.Id,&user.Username,&user.Email,&user.Created_at,&user.Updated_at);

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No User Found with given Id");
			return nil,err;
		} else {
			fmt.Println("Error Scanner error : ", err);
			return nil,err;
		}
	}
	fmt.Println("User fetch Successfully : ", user);
	return user,nil;
	
}