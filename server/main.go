package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main() {
	config := LoadConfig()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := OpenMySQL(ctx, config)
	if err != nil {
		log.Fatalf("connect mysql: %v", err)
	}
	defer db.Close()

	app := NewApp(config, NewMySQLStore(db))
	log.Printf("server listening on %s", config.Addr)
	if err := http.ListenAndServe(config.Addr, app); err != nil {
		log.Fatal(err)
	}
}
