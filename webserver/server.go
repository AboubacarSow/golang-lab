package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./static")))
	port := ":8080"
	fmt.Print("Server starts running on port" + port)
	fmt.Println(". Follow link: http://localhost" + port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}

}
