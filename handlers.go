package main

import(
	"net/http"
	"io/ioutil"
	"encoding/json"

	"github.com/unrolled/render"

)

func parseContainerParams(w http.ResponseWriter, r *http.Request) (Container, error) {
	container := Container{}

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return container, err
	}

	err = json.Unmarshal(body, &container)
	if err != nil {
		return container, err
	}

	return container, nil
}

func containerHandler(w http.ResponseWriter, r *http.Request) {
	rndr := render.New()

	dockerClient, err := newClient()
	if err != nil {
		http.Error(w, "Failed to create new docker client", http.StatusBadRequest)
	}

	switch r.Method {
		case "GET":
			container := Container{Image: r.URL.Query().Get("image")}
			containersList, err := container.List(dockerClient)
			if err != nil {
				http.Error(w, "Failed to get containers list", http.StatusBadRequest)
				return
			}

			rndr.JSON(w, http.StatusOK, containersList)
			return

		case "POST":
			container, err := parseContainerParams(w, r)
			if err != nil {
				http.Error(w, "Failed to read request body", http.StatusBadRequest)
				return
			}

			err = container.Run(dockerClient)
			if err != nil {
				http.Error(w, "Failed to start container: " + err.Error(), http.StatusBadRequest)
				return
			}

			rndr.JSON(w, http.StatusOK, container)
			return

		case "DELETE":
			container, err := parseContainerParams(w, r)
			if err != nil {
				http.Error(w, "Failed to read request body", http.StatusBadRequest)
				return
			}

			err = container.Remove(dockerClient)
			if err != nil {
				http.Error(w, "Failed to remove container", http.StatusBadRequest)
				return
			}

			rndr.JSON(w, http.StatusOK, true)
			return

		case "PUT":
			container, err := parseContainerParams(w, r)
			if err != nil {
				http.Error(w, "Failed to read request body", http.StatusBadRequest)
				return
			}

			err = container.Remove(dockerClient)
			if err != nil {
				http.Error(w, "Failed to remove container", http.StatusBadRequest)
				return
			}

			err = container.Run(dockerClient)
			if err != nil {
				http.Error(w, "Failed to start container", http.StatusBadRequest)
				return
			}

			rndr.JSON(w, http.StatusOK, true)
			return
	}

	http.Error(w, "Action not found", http.StatusNotFound)
}
