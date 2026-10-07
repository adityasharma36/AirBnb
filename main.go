package main

import (
	"fmt"
	"goMod/app"
	"goMod/config"
)

func main(){
	
	config.Load()
	cfg := app.NewConfig()
	app := app.NewApplication(cfg);

	

	app.Run()
	fmt.Println("Completed here")
}
