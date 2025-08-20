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
	email_error_channel := make(chan error)

	app := Bridge{
		DB:              db,
		WaitGroup:       &wg,
		EmailErrChannel: email_error_channel,
	}

	go app.ListenForShutdown()
	go ListenForEmailErrors(email_error_channel)

	mux := http.NewServeMux()

	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/register", app.Register)
	mux.HandleFunc("/activate/{hash}", app.Activate)
	mux.HandleFunc("/login", app.Login)
	mux.HandleFunc("/logout", app.Logout)
	mux.HandleFunc("/subscribe", app.ChooseSubscription)
	mux.HandleFunc("/subscribe/{plan_id}", app.SubscribeUser)

	mux.HandleFunc("POST /register", app.PostRegister)
	mux.HandleFunc("POST /login", app.PostLogin)

	server := &http.Server{
		Addr:    ":1000",
		Handler: RecoveryMiddleware(mux),
	}
	server.ListenAndServe()
}
