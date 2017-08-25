package main

import (
	"net/http"
	"fmt"
	"os"
)

func main() {
	containerID, _ := os.Hostname()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world! I'm container: %s.\n", containerID)
	})

	http.ListenAndServe(":8080", nil)
}
