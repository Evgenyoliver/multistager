package main

import "github.com/fsouza/go-dockerclient"

type(
	Container struct {
		Key    string `json:"key"`
		Branch string `json:"branch"`
		Image  string `json:"image"`
	}
)

func (container Container) Run(dockerClient *docker.Client) error {
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

	err = dockerClient.StartContainer(c.ID, &docker.HostConfig{PublishAllPorts: true})
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

func (container Container) List(dockerClient *docker.Client) ([]docker.APIContainers, error) {
	apiContainers, err := dockerClient.ListContainers(docker.ListContainersOptions{})
	if err != nil {
		return apiContainers, err
	}

	for id, cont := range apiContainers {
		if cont.Image != container.Image {
			apiContainers[id] = apiContainers[len(apiContainers)-1]
			apiContainers = apiContainers[:len(apiContainers)-1]
		}
	}

	return apiContainers, nil
}
