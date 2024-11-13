package main

import (
	"fmt"
	"rimo_bot_server/routers"
)

func main() {
	r := router.SetupRouter()

	fmt.Println("Starting server on :8080")
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
