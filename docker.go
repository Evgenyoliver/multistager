package main

import(
	"flag"
	"errors"
	"runtime"

	"github.com/fsouza/go-dockerclient"
)

func newClient() (*docker.Client, error) {
	switch runtime.GOOS {
	case "darwin":
		return docker.NewClientFromEnv()
	case "linux":
		dockerEndpoint := flag.String("e", "unix:///var/run/docker.sock", "Docker endpoint")
		flag.Parse()
		return docker.NewClient(*dockerEndpoint)
	}

	return nil, errors.New("Can't create docker client")
}
