package main

import (
	"donPass/backend/internal/config"
	"fmt"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	//TODO: init storage

	//TODO: init router

	//TODO: run server

}