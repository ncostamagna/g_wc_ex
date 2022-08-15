package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	port := ":3333"
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/posts", getPosts)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		fmt.Println(err)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /user request\n")
	io.WriteString(w, "This is my user endpoint!\n")
}
func getPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /posts request\n")
	io.WriteString(w, "This is my post endpoint!\n")
}
