package main

import (
	"io"
	"log"
	"net/http"
)

var Routess = []Route{
	{
		Path: "/hola",
		Name: "Root",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			log.Println("HHHH")
			//response := OkResponse{Message: "ok", Status: 200, Data: nil}
			io.WriteString(w, "Hello, world!\n")
			//w.Header().Set("Content-Type", "application/json")
			//json.NewEncoder(w).Encode(response)
		},
	},
}
