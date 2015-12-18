package main

import "github.com/fsouza/go-dockerclient"

type(
	Container struct {
		Key    string `json:"key"`
		Branch string `json:"branch"`
		Image  string `json:"image"`
		ID     string `json:"id"`
	}

	ContainersList []Container
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

	err = dockerClient.StartContainer(c.ID, &docker.HostConfig{})
	if err != nil {
		return err
	}

	return nil
}

func (container Container) Remove(dockerClient *docker.Client) error {
	removeOpts := docker.RemoveContainerOptions {
		ID: container.ID,
		Force: true,
	}

	return dockerClient.RemoveContainer(removeOpts)
}

func (container Container) List(dockerClient *docker.Client) (ContainersList, error) {
	return ContainersList{}, nil
}
