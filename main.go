package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	podName := os.Getenv("POD_NAME")
	if podName == "" {
		podName = "unknown-pod"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := fmt.Sprintf("Hello from Kubernetes!\n\nPod Name: %s\nHost: %s\n",
			podName,
			r.Host)
		fmt.Fprint(w, message)
		log.Printf("Request from %s", r.RemoteAddr)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	})

	port := "8080"
	log.Printf("Server starting on port %s...", port)
	log.Printf("Pod name: %s", podName)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
