package main

import(
	"log"
	"time"
	"errors"
	"strconv"
	"strings"
	"net/http"

	consul "github.com/hashicorp/consul/api"
)

type ServerConfig struct {
	StagesLimit int
	Links       []string
}

func (serverConfig *ServerConfig) load(client *consul.Client) error {
	kv := client.KV()
	pair, _, _ := kv.Get("multistager/stages_limit", nil)
	if pair == nil {
		return errors.New("Failed to fetch key from consul")
	}

	limit, err := strconv.Atoi(string(pair.Value))
	if err != nil {
		return err
	}

	serverConfig.StagesLimit = limit

	pair, _, _ = kv.Get("multistager/links", nil)
	if pair == nil {
		return errors.New("Failed to fetch key from consul")
	}

	links := strings.Split(string(pair.Value), ",")
	if err != nil {
		return err
	}

	serverConfig.Links = links

	return nil
}

func newConsulClient() (*consul.Client, error) {
	consulConfig := &consul.Config{
		Address:    "localhost:8500",
		Scheme:     "http",
		HttpClient: http.DefaultClient,
	}

	return consul.NewClient(consulConfig)
}

func consulWorker(client *consul.Client) {
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

	err := agent.ServiceRegister(service)
	if err != nil {
		log.Println(err)
	}

	tick := time.Tick(time.Duration(15 * time.Second))
	for range tick {
		err = agent.PassTTL("service:multistager1", "Internal TTL ping")
		if err != nil {
			log.Println(err)
		}

		serverConfig.load(client)
	}
}
