package main

import (
    "fmt"
	"net/http"
	"encoding/json"
	"os"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Congratulations! Your Go application has been successfully deployed on Kubernetes.")
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/", EnvHandler)
    http.ListenAndServe(":3000", nil)
}

func EnvHandler(rw http.ResponseWriter, req *http.Request) {
	environment := make(map[string]string)
	for _, item := range os.Environ() {
		splits := strings.Split(item, "=")
		key := splits[0]
		val := strings.Join(splits[1:], "=")
		environment[key] = val
	}

	data, err := json.MarshalIndent(environment, "", "")
	if err != nil {
		data = []byte("Error marshalling env vars: " + err.Error())
	}

	rw.Write(data)
}