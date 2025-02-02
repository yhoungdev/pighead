package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {

	const Port int = 5500
	port := flag.Int("port", 5500, "Port to serve files on")
	dir := flag.String("dir", ".", "Directory to serve files from")
	flag.Parse()

	addr := fmt.Sprintf(":%d", *port)
	fmt.Printf("Serving files from %s http://localhost%s\n", *dir, addr)

	go func() {
		time.Sleep(5 * time.Millisecond)
	}()

}
