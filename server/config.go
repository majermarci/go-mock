package main

import (
	"fmt"
	"log/slog"
	"os"
	"regexp"

	"gopkg.in/yaml.v2"
)

type method struct {
	Headers map[string]string `yaml:"headers"`
	Body    string            `yaml:"body"`
	Status  int               `yaml:"status"`
}

type config struct {
	Endpoints map[string]map[string]method `yaml:",inline"`
}

func loadConfig(file string) (conf config, err error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return conf, err
	}

	// Unmarshal the YAML into the config struct
	err = yaml.UnmarshalStrict(data, &conf)
	if err != nil {
		return conf, err
	}

	// Define the regular expressions for validation
	pathRegexp := regexp.MustCompile(`^[a-zA-Z0-9/-]+$`)
	methodRegexp := regexp.MustCompile(`^(GET|POST|PUT|DELETE|HEAD|OPTIONS|PATCH|TRACE)$`)
	headerKeyRegexp := regexp.MustCompile(`^[a-zA-Z0-9-]+$`)
	headerValueRegexp := regexp.MustCompile(`^[a-zA-Z0-9/;=.\s-]+$`)

	reservedPaths := map[string]bool{
		"/healthz": true,
	}

	// Validate the path, methods, headers and status codes
	for path, methods := range conf.Endpoints {
		if !pathRegexp.MatchString(path) {
			return conf, fmt.Errorf("invalid path: %s", path)
		}

		if reservedPaths[path] {
			slog.Warn(fmt.Sprintf("Path %q in the configuration conflicts with a reserved path", path))
		}

		for method, endpoint := range methods {
			if !methodRegexp.MatchString(method) {
				return conf, fmt.Errorf("invalid method: %s", method)
			}

			for key, value := range endpoint.Headers {
				if !headerKeyRegexp.MatchString(key) || !headerValueRegexp.MatchString(value) {
					return conf, fmt.Errorf("invalid header: %s: %s", key, value)
				}
			}

			if endpoint.Status < 100 || endpoint.Status > 599 {
				return conf, fmt.Errorf("invalid status: %d! status must be between 100 and 599", endpoint.Status)
			}
		}
	}

	return conf, nil
}
