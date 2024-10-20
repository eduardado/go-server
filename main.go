package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	m := http.NewServeMux() // multiplexor, routes incoming http request to our handlers in code

	m.HandleFunc("/", handlePage) // multiplexor handles any request to the root with our function

	const addr = ":8000" // address we are serving on

	srv := http.Server{ // server
		Handler:      m,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	fmt.Println("Server started on port:", addr)
	err := srv.ListenAndServe() // it will block, serving forever, in this line until something goes wrong
	log.Fatal(err)
}
func handlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	const page = "<html><body><p>Hello from the Go Server</p></body></html>"
	w.WriteHeader(200)
	w.Write([]byte(page))
}
