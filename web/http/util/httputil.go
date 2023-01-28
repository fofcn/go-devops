package util

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResp)
}

func WriteJsonString(w http.ResponseWriter, data string) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(data))
}

func WriteJsonBytes(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
