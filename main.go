package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", indexHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

type Resp struct {
	ReceivedTime string `json:"receivedTime"`
	Digest string `json:"digest"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	receivedTime := time.Now().Format("2006/01/02 15:04:05.000")
	digest := sha256.Sum256([]byte(receivedTime))

	resp := Resp{
		ReceivedTime: receivedTime,
		Digest:       fmt.Sprintf("%x", digest),
	}

	output, err := json.Marshal(&resp)
	if err != nil {
		log.Fatal(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}