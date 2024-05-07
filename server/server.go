package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
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

			slog.Info("Request mocked", "method", method, "path", path)

			w.WriteHeader(endpoint.Status)
			if endpoint.Body != "" {
				w.Write([]byte(endpoint.Body))
			}
		} else {
			slog.Error("Not found", "method", method, "path", path)
			http.NotFound(w, r)
		}
	})

	_, err := strconv.Atoi(*port)
	if err != nil {
		slog.Error("Invalid port:"+*port, "details", err)
		os.Exit(1)
	}

	slog.Info("Starting server on port " + *port)
	err = http.ListenAndServe(":"+*port, nil)
	if err != nil {
		slog.Error("Error with server", "details", err)
		os.Exit(1)
	}
}
