package main

import (
	"log"
	"os"
	"strconv"

	api "sample-api/cmd/http-api"
)

func main() {
	port, _ := strconv.ParseInt(os.Getenv("PORT"), 10, 64)
	if port == 0 {
		port = 8080
	}

	logger := log.New(os.Stdout, " sample-api ", log.Default().Flags())
	api.Run(int(port), logger)
}
