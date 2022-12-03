package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type thumbnailRequest struct {
	Url string `json:"url"`
}

type screenshotAPIRequest struct {
	Token          string `json:"token"`
	Url            string `json:"url"`
	Output         string `json:"output"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	ThumbnailWidth int    `json:"thumbnail_width"`
}

func thumbnailHandler(w http.ResponseWriter, r *http.Request) {
	var decoded thumbnailRequest
	err := json.NewDecoder(r.Body).Decode(&decoded)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	apiRequest := screenshotAPIRequest{
		Token:          "QX33FDK-34944RG-PVSFV74-2PVDP0X",
		Url:            decoded.Url,
		Output:         "json",
		Width:          1920,
		Height:         1080,
		ThumbnailWidth: 300,
	}
	jsonString, err := json.Marshal(apiRequest)
	checkError(err)
	req, err := http.NewRequest("POST", "https://shot.screenshotapi.net/screenshot?token=QX33FDK-34944RG-PVSFV74-2PVDP0X&url=https%3A%2F%2Fgoogle.com&output=image&file_type=png&wait_for_event=load", bytes.NewBuffer(jsonString))
	checkError(err)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)
	checkError(err)
	defer response.Body.Close()
	type screenshotAPIResponse struct {
		Screenshot string `json:"screenshot"`
	}
	var apiResponse screenshotAPIResponse
	err = json.NewDecoder(response.Body).Decode(&apiResponse)
	checkError(err)
	_, err = fmt.Fprintf(w, `{ "screenshot" : %s }`, apiResponse.Screenshot)
	checkError(err)
	fmt.Printf("Got the following url : %s\n", decoded.Url)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	http.HandleFunc("/api/thumbnail", thumbnailHandler)
	fs := http.FileServer(http.Dir("../frontend/dist"))
	http.Handle("/", fs)
	fmt.Println("Server listening on port 3000...")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}
