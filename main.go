package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
)

// HTML response
func handler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}

	/* for i, addr := range addrs {
	    fmt.Printf("%d %v\n", i, addr)
	}*/

	log.Printf("%s %s - %s", r.Method, r.URL.Path, r.UserAgent())
	fmt.Fprintf(w, html, version, hostname, addrs[1])
}

func cleanup() {
	log.Println("Got CTRL-C...cleaning up now.")
}

func main() {
	log.Println("Starting up...")
	log.Printf("You´re running version %s of this program.", version)

	// Catch CRTL-C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		cleanup()
		os.Exit(0)
	}()

	// Call handler func (create html) and start serving
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
