package main

import (
	"log"
	"net/http"

	"github.com/Nezent/Go-Net/config"
	"github.com/Nezent/Go-Net/routes"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)



func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config.Connect()
	router := &http.Server{
		Addr: ":8080",
		Handler: routes.PageRoute(config.DB),
	}
	router.ListenAndServe()


}