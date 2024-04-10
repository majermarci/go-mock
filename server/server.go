package main

import (
	"log"
	"net/http"
)

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

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
