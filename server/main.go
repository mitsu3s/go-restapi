package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Response struct {
	Column []struct {
		Message string `json:"message"`
	} `json:"Column"`
}

type Post struct {
	Message string `json:"message"`
}

func main() {
	var posts []Post
	url := "http://host.docker.internal:8000/hello"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp)

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		fmt.Println("Error: status code", resp.StatusCode)
	}

	if err := json.Unmarshal(body, &posts); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(posts)
}
