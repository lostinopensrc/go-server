package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lostinopensrc/go-server/web"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)              // using inbuilt FileServer as handler to directs to index.html under static directory
	http.HandleFunc("/form", web.FormHandler) // using cusotom FormHandler function defined in web package

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
