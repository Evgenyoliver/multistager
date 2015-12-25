package main

import(
	"errors"
	"strings"

	"github.com/fsouza/go-dockerclient"
)

type(
	Container struct {
		Key    string `json:"key,omitempty"`
		Branch string `json:"branch"`
		Image  string `json:"image"`
	}
)

func (container Container) Run(dockerClient *docker.Client) error {
	containersList, err := container.List(dockerClient)
	if err != nil {
		return err
	}

	if len(containersList) >= serverConfig.StagesLimit {
		return errors.New("Containers limit exceeded")
	}

	containerOpts := docker.CreateContainerOptions{
		Name: container.Branch,
		Config: &docker.Config{
			Env: []string{
				"SERVICE_TAGS=" + container.Branch,
				"KEY=" + container.Key,
			},
			Image: container.Image,
		},
	}

	c, err := dockerClient.CreateContainer(containerOpts)
	if err != nil {
		return err
	}

	hostConfig := &docker.HostConfig{
		PublishAllPorts: true,
		Links: serverConfig.Links,
	}

	err = dockerClient.StartContainer(c.ID, hostConfig)
	if err != nil {
		return err
	}

	return nil
}

func (container Container) Remove(dockerClient *docker.Client) error {
	removeOpts := docker.RemoveContainerOptions {
		ID: container.Branch,
		Force: true,
	}

	return dockerClient.RemoveContainer(removeOpts)
}

func (container Container) List(dockerClient *docker.Client) ([]Container, error) {
	apiContainers, err := dockerClient.ListContainers(docker.ListContainersOptions{})
	if err != nil {
		return []Container{}, err
	}

	containersList := []Container{}

	for _, cont := range apiContainers {
		if cont.Image == container.Image {
			containerItem := Container{
				Branch: strings.Replace(cont.Names[0], "/", "", -1),
				Image: cont.Image,
			}

			containersList = append(containersList, containerItem)
		}
	}

	return containersList, nil
}
