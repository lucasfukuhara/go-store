package main

import (
	"net/http"
	"text/template"
)

var templateList = template.Must(template.ParseGlob("templates/*.html")) // load all files with .html from templates folder

func main() {
	http.HandleFunc("/", index)       // will link the address / with the func index
	http.ListenAndServe(":8000", nil) //start the server on port 8000
}

func index(w http.ResponseWriter, r *http.Request) {
	/* will send the template index as response, the last parameter is to send
	some information to the template.
	The Index is the anotation {{define}} we added on index.html
	*/
	templateList.ExecuteTemplate(w, "Index", nil)
}
