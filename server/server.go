package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

var port = flag.String("p", "8080", "Port where the server will listen (default: 8080)")

func serve(c config) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		method := r.Method

		if endpoint, ok := c.Endpoints[path][method]; ok {
			w.Header().Set("System", "mock-server")

			for key, value := range endpoint.Headers {
				w.Header().Set(key, value)
			}

			log.Printf("Request mocked: %s %s\n", method, path)

			w.WriteHeader(endpoint.Status)
			if endpoint.Body != "" {
				w.Write([]byte(endpoint.Body))
			}
		} else {
			log.Printf("Not found: %s %s\n", method, path)
			http.NotFound(w, r)
		}
	})

	_, err := strconv.Atoi(*port)
	if err != nil {
		log.Fatalf("Invalid port: %v", err)
	}

	log.Printf("Starting server on :%s...", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
