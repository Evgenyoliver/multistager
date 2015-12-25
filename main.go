package main

import(
	"log"
	"net/http"
)

var serverConfig *ServerConfig = &ServerConfig{}

func main() {
	client, err := newConsulClient()
	if err != nil {
		log.Fatal("Failed to connect consul", err)
	}

	err = serverConfig.load(client)
	if err != nil {
		log.Fatal("Failed to fetch configuration from consul", err)
	}

	go consulWorker(client)

	http.HandleFunc("/v1/container", containerHandler)
	http.ListenAndServe(":8080", nil)
}
