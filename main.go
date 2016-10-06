package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	var (
		_ = flag.String("admin-port", env("PORT", "8080"), "The port for the admin api")
	)
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("You must specify the `server` or `worker` subcommand.")
	}

	cmd := args[0]

	switch cmd {
	case "server":
		port := "8080"
		log.Println("Starting on port", port)
		log.Fatal(http.ListenAndServe(":"+port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s - %s", r.Method, r.URL)
			w.WriteHeader(200)
			fmt.Fprintf(w, "Hello from port %s\n", port)
		})))
	case "worker":
		for {
			<-time.After(1 * time.Second)
			fmt.Printf("Hard work %d...\n", rand.Int())
		}
	default:
		log.Fatalf("Unknown subcommand: %s", cmd)
	}
}

func env(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		value = fallback
	}
	return value
}
