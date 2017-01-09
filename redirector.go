package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
)

type Config struct {
	Server    Server
	Redirects map[string]Redirect
}

type Server struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

type Redirect struct {
	To   string `yaml:"to"`
	With int    `yaml:"with"`
}

func main() {
	// Read in config file
	c, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatal("Error reading config: ", err)
	}

	// Unmarshal YAML
	config := Config{}
	err = yaml.Unmarshal(c, &config)
	if err != nil {
		log.Fatal("Error marshalling YAML: ", err)
	}

	// Setup handler for redirects
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, found := config.Redirects[r.Host]

		if !found {
			log.Printf("Client %s %s -> Not Found (404)", r.RemoteAddr, r.Host)
			http.NotFound(w, r)
			return
		}

		To := config.Redirects[r.Host].To
		With := config.Redirects[r.Host].With

		log.Printf("Client: %s %s -> %s (%d)\n", r.RemoteAddr, r.Host, To, With)
		http.Redirect(w, r, To, With)
		return
	})

	// Start HTTP Server
	ServerAddress := config.Server.Address
	ServerPort := config.Server.Port

	log.Printf("Listening on %v:%v\n", ServerAddress, ServerPort)
	err = http.ListenAndServe(fmt.Sprintf("%v:%v", ServerAddress, ServerPort), nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
