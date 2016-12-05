package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Config struct {
	Server    Server
	Redirects map[string]map[string]string
}

type Server struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

func main() {
	c, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatal("Error reading config: ", err)
	}

	config := Config{}
	err = yaml.Unmarshal(c, &config)
	if err != nil {
		log.Fatal("Error marshalling YAML: ", err)
	}

	// http.Request.Host
	// RedirectHandler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		to := config.Redirects[r.Host]["to"]
		with, _ := strconv.Atoi(config.Redirects[r.Host]["with"])

		http.Redirect(w, r, to, with)
	})

	fmt.Printf("Listening on %v:%v\n", config.Server.Address, config.Server.Port)
	err = http.ListenAndServe(fmt.Sprintf("%v:%v", config.Server.Address, config.Server.Port), nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}

}
