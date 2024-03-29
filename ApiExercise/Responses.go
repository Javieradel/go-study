package main

type OkResponse struct {
	Message string `json:"message"`
	Status  uint8  `json:"status"`
	Data    any    `json:"Data"`
}
