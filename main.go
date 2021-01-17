package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	content := getContent()
	identifier := identifySelf()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(content))
		if identifier != "" {
			writer.Write([]byte(fmt.Sprintf("\n\nServed by [%s]", identifier)))
		}
	})

	port := getPort()
	log.Printf("Listening on %v", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
	if err != nil {
		panic(err)
	}
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return port
}

func getContent() string {
	content := os.Getenv("CONTENT")
	if content == "" {
		content = "Default content (environment variable 'CONTENT' not set)"
	}
	return content
}

func identifySelf() string {
	if fileExists("/etc/podinfo/name") && fileExists("/etc/podinfo/uid") {
		name, err := ioutil.ReadFile("/etc/podinfo/name")
		if err != nil {
			log.Print("Error reading name: ", err.Error())
			return ""
		}
		uid, err := ioutil.ReadFile("/etc/podinfo/uid")
		if err != nil {
			log.Print("Error reading uid: ", err.Error())
			return ""
		}

		return fmt.Sprintf("%s: %s", name, uid)
	}
	return ""
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		log.Printf("Path [%s] does not exist", path)
		return false
	}

	log.Printf("Path [%s] exists", path)
	return true
}
