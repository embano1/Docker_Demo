package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var html = `
<html>
<h1>Hi there!</h1>
<p>My version is: %s</p>
<p>My hostname is: %s</p>
</html>
`

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, html, version, hostname)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
