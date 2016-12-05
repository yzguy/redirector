package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Server    Server
	Redirects []Redirect
}

type Server struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

type Redirect struct {
	From string `yaml:"from"`
	To   string `yaml:"to"`
	With int    `yaml:"with"`
}

func main() {
	c, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	config := Config{}
	err = yaml.Unmarshal(c, &config)
	if err != nil {
		panic(err)
	}

	for _, redirect := range config.Redirects {
		fmt.Printf("From: %v, To: %v, With: %v\n", redirect.From, redirect.To, redirect.With)
	}
}
