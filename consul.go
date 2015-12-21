package main

import(
	"log"
	"time"
	"net/http"

	consul "github.com/hashicorp/consul/api"
)

func consulWorker() {
	consulConfig := &consul.Config{
		Address:    "localhost:8500",
		Scheme:     "http",
		HttpClient: http.DefaultClient,
	}

	client, err := consul.NewClient(consulConfig)
	if err != nil {
		log.Println("Failed to connect consul", err)
		return
	}

	agent := client.Agent()

	check := &consul.AgentServiceCheck{
		TTL: "30s",
	}

	service := &consul.AgentServiceRegistration{
		ID: "multistager1",
		Name: "multistager",
		Tags: []string{"1"},
		Port: 8080,
		Check: check,
	}

	err = agent.ServiceRegister(service)
	if err != nil {
		log.Println(err)
	}

	tick := time.Tick(time.Duration(15 * time.Second))
	for range tick {
		err = agent.PassTTL("service:multistager1", "Internal TTL ping")
		if err != nil {
			log.Println(err)
		}
	}
}