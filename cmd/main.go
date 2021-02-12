package main

import "test-project/api"

func main() {
	apiServer := api.New()

	apiServer.Run(":8080")
}
