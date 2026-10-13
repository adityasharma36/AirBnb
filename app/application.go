package app

import (
	"fmt"
	"goMod/config"
	"goMod/controllers"
	db "goMod/db/repositories"
	"goMod/router"
	"goMod/service"
	"net/http"
	"time"
	dbConfig "goMod/config/db"
)
 

type Config struct{
	Addr string
}

type Application struct{

	Config Config
	Store db.Storage

}

// this is Constructor to connect config to application
func NewApplication(c  Config) *Application{
	return &Application{
		Config: c,
		Store: *db.NewStorage(),
	}
}

// config Constructor 
func NewConfig() Config{
	port := config.GetString("PORT",":8080")
	return Config{
		Addr: port,
	}
}

func (app *Application) Run() error {

	database,err := dbConfig.ConectToDb()
	if err != nil{
		fmt.Println("failed to connected db")
		return err;
	}
	ur:= db.NewUserRepository(database)
	us:= service.NewUserService(ur);
	uc := controllers.NewUserController(us);
	uRouter := router.NewUserRouter(uc);
	
	server := &http.Server{
		Addr: app.Config.Addr,
		Handler: router.SetUpRouter(uRouter),
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10* time.Second,
	}

	fmt.Println("Starting server on  : ", app.Config.Addr )

	return server.ListenAndServe();

}