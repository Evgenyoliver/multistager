package main

import(
	"net/http"
	"io/ioutil"
	"encoding/json"

	"github.com/unrolled/render"
)

func containerHandler(w http.ResponseWriter, r *http.Request) {
	rndr := render.New()

	dockerClient, err := newClient()
	if err != nil {
		http.Error(w, "Failed to connect to docker client", http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	container := Container{}
	err = json.Unmarshal(body, &container)
	if err != nil {
		http.Error(w, "Failed to unmarshal JSON file", http.StatusInternalServerError)
		return
	}

	switch r.Method {
		case "GET":
			containersList, err := container.List(dockerClient)
			if err != nil {
				http.Error(w, "Failed to get containers list", http.StatusBadRequest)
				return
			}

			rndr.JSON(w, http.StatusOK, containersList)
			return

		case "POST":
			err = container.Run(dockerClient)
			if err != nil {
				http.Error(w, "Failed to start container", http.StatusBadRequest)
				return
			}

			rndr.JSON(w, http.StatusOK, container)
			return

		case "DELETE":
			err = container.Remove(dockerClient)
			if err != nil {
				http.Error(w, "Failed to remove container", http.StatusBadRequest)
				return
			}

			rndr.JSON(w, http.StatusOK, true)
			return

		case "PUT":
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

func main() {
	http.HandleFunc("/v1/container", containerHandler)
	http.ListenAndServe(":8080", nil)
}
