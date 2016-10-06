package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("You must specify the `server` or `worker` subcommand.")
	}

	cmd := args[0]

	switch cmd {
	case "server":
		for _, port := range []string{"80", "8080"} {
			port := port
			log.Printf("Starting on %s", port)
			go http.ListenAndServe(":"+port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				log.Printf("%s - %s", r.Method, r.URL)
				w.WriteHeader(200)
				fmt.Fprintf(w, "Hello from port %s\n", port)
			}))
		}
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
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
