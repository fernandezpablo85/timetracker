package main

import (
	"fmt"
	"net/http"
	"os"
	"html/template"
	"path"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))
  http.HandleFunc("/", serveTemplate)

	port := os.Getenv("PORT")
	fmt.Printf("listening on port %s...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
  main := path.Join("templates", "main.html")

  tmpl, _ := template.ParseFiles(main)
  tmpl.Execute(w, nil)
}