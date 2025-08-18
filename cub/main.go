package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Dominus Iesus Christus")

	dsn := "host=localhost user=postgres password=password dbname=concurrency port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panicf("Could not connect to database: %v\n", err)
	}
	db.AutoMigrate(&User{}, &Plan{})

	wg := sync.WaitGroup{}
	app := Bridge{
		DB:        db,
		WaitGroup: &wg,
	}
	go app.ListenForShutdown()

	mux := http.NewServeMux()

	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/register", app.Register)
	mux.HandleFunc("/activate", app.Activate)
	mux.HandleFunc("/login", app.Login)
	mux.HandleFunc("/logout", app.Logout)

	mux.HandleFunc("POST /login", app.PostLogin)
	mux.HandleFunc("POST /register", app.PostRegister)

	server := &http.Server{
		Addr:    ":1000",
		Handler: RecoveryMiddleware(mux),
	}
	server.ListenAndServe()
}
