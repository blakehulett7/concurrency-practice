package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func (app *Bridge) ListenForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	app.Shutdown()
	os.Exit(0)
}

func (app *Bridge) Shutdown() {
	app.WaitGroup.Wait()
	fmt.Println("Et Spiritus Sancti...")
}
