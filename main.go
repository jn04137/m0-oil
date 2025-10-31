package main

import (
	"log"
	"net/http"

	"thdr/m0-oil/car"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Connect("mysql", "m0_user:password@(localhost:3306)/m0_database")
	if err != nil {
		log.Fatalf("Failed to connect to db: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Ping to DB failed: %v", err)
	}

	carRepo := car.NewCarRepository(db);
	carRouter := car.NewCarRouter(carRepo)

	mux := http.NewServeMux()
	mux.Handle("/car/", http.StripPrefix("/car", carRouter))

	http.ListenAndServe(":8080", enabledCors(mux))
}

func enabledCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

