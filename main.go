package main

import (
	"test-d-2/connection"
	"test-d-2/handlers"
)

func main() {
	connection.Connect()

	handlers.HandleReq()
}
