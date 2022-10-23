package main

import (
	"E-TamuAPI/api"
	"E-TamuAPI/seeders"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	godotenv.Load()
	rand.Seed(time.Now().Unix())
	db, err := sql.Open("sqlite3", "file:./database.db?_foreign_keys=true")

	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Database Connected!")

	seeders.MigrateDB(db)

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8000"
	}

	api := api.NewAPI(db)
	r := api.GetRouter()
	log.Println("Listening in port 8000")
	http.ListenAndServe(":"+PORT, r)
}
