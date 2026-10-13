package config

import (
	"database/sql"
	"fmt"
	"goMod/config"

	"github.com/go-sql-driver/mysql"
)



func ConectToDb() (*sql.DB,error){

	cfg := mysql.Config{
    User:             config.GetString("DB_User" , "root"),
    Passwd:           config.GetString("DB_Passwd","root"),
    Net:              config.GetString("DB_NET","tcp"),
    Addr:             config.GetString("DB_Addr","127.0.0.1:3306"),
    DBName:           config.GetString("DB_Name","auth_dev"),
    ParseTime:        true,
    AllowNativePasswords: true, // Often needed depending on your MySQL version
}

	
	
	db, err :=sql.Open("mysql",cfg.FormatDSN())
	
	if err != nil{
		fmt.Println("failed to connected Db",err)
		return nil, err;
	}



	ping := db.Ping();

	if ping != nil{
		fmt.Println("ping error occured ");
		return nil,err;
	}

	fmt.Println("successfully connected to db : ", cfg.DBName)
	return  db,nil;

}