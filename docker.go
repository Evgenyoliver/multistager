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
		return docker.NewClient("unix:///var/run/docker.sock")
	}

	return nil, errors.New("Can't create docker client")
}
