package main

import (
	"github.com/IlhamRamadhan-IR/api-team-management-system/config"
	"github.com/IlhamRamadhan-IR/api-team-management-system/router"
)

func init() {
	config.SetupConfiguration()
}

func main() {
	db := config.InitDatabase()
	router.NewRouter(db)
}
