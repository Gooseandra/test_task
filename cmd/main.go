package main

import (
	"database/sql"
	"fmt"
	"github.com/go-yaml/yaml"
	_ "github.com/lib/pq"
	"log"
	"os"
	app2 "test_task/internal/app"
	"test_task/internal/protocols/fiber"
	"test_task/internal/servises"
	"test_task/internal/storages/psql"
)

var app app2.App

func main() {
	var settings Config
	var filename string

	switch len(os.Args) {
	case 1:
		filename = "db.yml"
	case 2:
		filename = os.Args[1]
	default:
		log.Fatal("Invalid argument count")
	}

	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Read file: %s", err.Error())
	}
	err = yaml.Unmarshal([]byte(bytes), &settings)
	if err != nil {
		log.Fatalf("Unmarshal file: %s", err.Error())
	}
	db, err := sql.Open(settings.Database.Type, settings.Database.Arguments)
	if err != nil {
		log.Fatalf("Open database: %s", err.Error())
	}

	subscribe := psql.NewPsqlSubscribe(db)
	fetch := psql.NewPsqlFetch(db)
	app.Subscribe = servises.NewSubscribe(subscribe)
	app.Fetch = servises.NewFetch(fetch)
	app.Binance = servises.NewBinance("https://api.binance.com/api/v3/ticker/price?symbol=", fetch, subscribe)
	app.Fib = fiber.Init(app)

	err = app.Binance.Init()

	if err != nil {
		fmt.Println(err.Error())
	}

	log.Fatal(app.Fib.Listen(":3000"))
}
