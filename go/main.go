package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

var cfg *config

type config struct {
	DB DBConfig
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

type Data struct {
	ID        int64 `json:"id"`
	Col_texto string
	Col_dt    time.Time
}

func Load() error {
	cfg = new(config)

	cfg.DB = DBConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Pass:     "1234",
		Database: "postgres",
	}

	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func main() {
	Load()
	http.HandleFunc("/tb01", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var data Data
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		databaseConfig := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			GetDB().Host,
			GetDB().Port,
			GetDB().User,
			GetDB().Pass,
			GetDB().Database,
		)
		db, err := sql.Open("postgres", databaseConfig)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		_, err = db.Exec("INSERT INTO tb01 (col_texto, col_dt) VALUES ($1, $2)", data.Col_texto, time.Now())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintln(w, "Data received and inserted successfully")
	})

	fmt.Println("Server running at http://localhost:3001")
	log.Fatal(http.ListenAndServe(":3001", nil))
}
