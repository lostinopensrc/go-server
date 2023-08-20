package web

import (
	"fmt"
	"net/http"
)

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse form error: %v", err)
		return
	}
	fmt.Fprintf(w, "Post request success\n")
	Name := r.FormValue("Name")
	Address := r.FormValue("Address")
	fmt.Fprintf(w, "Name: %s\n", Name)
	fmt.Fprintf(w, "Address: %s\n", Address)

}

func DefaultHandler() http.Handler {
	FileServer := http.FileServer(http.Dir("./static"))
	return FileServer
}
