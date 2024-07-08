package main

import (
	"fmt"

	"github.com/AswinJoseOpen/Login-Auth/config"
	"github.com/AswinJoseOpen/Login-Auth/db"
	"github.com/AswinJoseOpen/Login-Auth/server"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		panic("Unable to load config")
	}
	fmt.Println("Test is starting")
	db, err := db.Init()
	if err != nil {
		fmt.Printf("unable to connect to DB, %v", err)
	}
	server.Start(db, config)
}
