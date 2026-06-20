package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./static")))
	port := ":8080"
	fmt.Println("Server starts running on port" + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}

}
