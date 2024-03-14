package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-yaml/yaml"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type PriceData struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price,string"`
}

func getFromBinance(symbol string, db *sql.DB) {
	for {
		resp, err := http.Get("https://api.binance.com/api/v3/ticker/price?symbol=" + symbol)
		if err != nil {
			fmt.Println("Ошибка при выполнении запроса:", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Ошибка при чтении данных из ответа:", err)
			return
		}

		var price PriceData
		err = json.Unmarshal(body, &price)
		if err != nil {
			fmt.Println("Ошибка при парсинге JSON:", err)
			return
		}
		fmt.Printf("Символ: %s\nЦена: %f USDT\n", price.Symbol, price.Price)
		err = AddPostgres(db, price)
		if err != nil {
			log.Println(err.Error())
		}
		time.Sleep(time.Minute)
	}
}

func main() {

	var settings Settings
	bytes, fail := os.ReadFile("db.yml")
	if fail != nil {
		log.Println(fail.Error())
		log.Panic(fail.Error())
	}
	fail = yaml.Unmarshal([]byte(bytes), &settings)
	if fail != nil {
		log.Panic(fail.Error())
	}
	db, err := sql.Open(settings.Database.Type, settings.Database.Arguments)
	if err != nil {
		log.Println(err.Error())
	}

	symbols, err := GetWatchingCoins(db)
	if err != nil {
		log.Println(err.Error())
	}

	for _, item := range symbols {
		go getFromBinance(item, db)
	}

	app := fiber.New()
	fetch(app, db)
	add_ticker(app, db)

	log.Fatal(app.Listen(":3000"))
}
