package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	postsURL := "https://jsonplaceholder.typicode.com/posts"
	var httpClient = &http.Client{Timeout: 10 * time.Second}

	response, error := httpClient.Get(postsURL)

	if error != nil {
		panic(error.Error())
	}

	var posts []Post

	error = json.NewDecoder(response.Body).Decode(&posts)

	if error != nil {
		panic(error.Error())
	}

	fmt.Println(posts[0])
}
