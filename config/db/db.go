package config

import (
	"database/sql"
	"fmt"
	"goMod/config"

	"github.com/go-sql-driver/mysql"
)



func ConectToDb() (*sql.DB,error){

	cfg := mysql.NewConfig();

	cfg.User = config.GetString("DB_User" , "root");
	cfg.Passwd = config.GetString("DB_Passwd","root");
	cfg.Net = config.GetString("DB_NET","tcp");
	cfg.DBName = config.GetString("DB_Name","auth_dev");
	cfg.Addr = config.GetString("DB_Addr","127.0.0.1:3306");


	fmt.Println("Connected to Db",cfg.DBName);

	db, err :=sql.Open("mysql",cfg.FormatDSN())

	if err != nil{
		fmt.Println("failed to connected Db",err)
		return nil, err;
	}

	fmt.Println("connecting to DB ....")

	ping := db.Ping();

	if ping != nil{
		fmt.Println("ping error occured ");
		return nil,err;
	}

	fmt.Println("successfully connected to db : ", cfg.DBName)
	return  db,nil;

}