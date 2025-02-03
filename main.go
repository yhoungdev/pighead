package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.Int("port", 5500, "Port to serve files on")
	dir := flag.String("dir", "./htmx", "Directory to serve files from")
	flag.Parse()

	addr := fmt.Sprintf(":%d", *port)
	fmt.Printf("Serving files from %s at http://localhost%s\n", *dir, addr)

	http.Handle("/", http.FileServer(http.Dir(*dir)))

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
