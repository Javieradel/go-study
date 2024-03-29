package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

type Bootstrap struct {
	Port    string
	AppName string
}

func (b *Bootstrap) getEnv() bool {
	var env map[string]string
	var err error
	env, err = godotenv.Read()

	b.AppName = env["APP_NAME"]
	b.Port = env["PORT"]

	if env["APP_NAME"] == "" {
		println("Starting as Anonimous app")
		b.AppName = "Anonimous app"
	}

	if env["PORT"] == "" {
		b.Port = "8080"
		println("Port default is: ", b.Port)
	}

	if err != nil {
		println("error", err)
		log.Fatal(err)
		return false
	}
	return true
}

func (b *Bootstrap) Start() {
	isLoadedEnvs := b.getEnv()
	fmt.Printf("Starting app %s\n", b.AppName)

	router := new(Router)

	if !isLoadedEnvs {
		log.Fatal("Error loading .env file", isLoadedEnvs)
		return
	}

	println("Loading endpoints")
	for key, route := range Routess {
		fmt.Printf("%d.- %s loaded\n", key, route.Name)
		router.registerRoute(route)
	}

	url := fmt.Sprintf(":%s", b.Port)
	log.Printf("Server listening on port %s %s", b.Port, url)
	log.Fatal(http.ListenAndServe(url, nil))
}
