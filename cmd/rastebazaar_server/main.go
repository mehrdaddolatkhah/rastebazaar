package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	infra "rastebazaar/pkg/infra"

	mongodb "rastebazaar/pkg/repository/mongo"

	"rastebazaar/pkg/transport/rest"
)

func main() {

	router := rest.RouteHandler(initInfrastructure())
	log.Fatal(http.ListenAndServe(":8080", router))

}

func initInfrastructure() *infra.Infrastructure {

	// connect to MongoDB
	client, err := mongodb.ConnectToDatabase()

	if err != nil {
		fmt.Println("Failed to connect MongoDB", err)
	}

	tmpl, err := template.ParseGlob("templates/*.html")

	if err != nil {
		panic(err)
	}

	infra := infra.Infrastructure{
		MongoClient: client,
		Template:    tmpl,
	}

	return &infra
}
