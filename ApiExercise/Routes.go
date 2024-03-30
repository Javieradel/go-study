package main

import "log"

var Routess = []Route{
	{
		Path: "/{$}",
		Name: "Root",
		Handler: func(r Response) {
			log.Println("Root")
			r.SetBody(JSONApiResponse{Data: "Asdasdasdasd"}).Send()
		},
	},
	{
		Path: "/hello",
		Name: "Saludo",
		Handler: func(r Response) {
			log.Println("Saludo")
			r.SetBody(JSONApiResponse{Data: "Hola"}).Send()
		},
	},
}
