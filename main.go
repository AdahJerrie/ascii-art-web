package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func homeHandler(w http.ResponseWriter, req *http.Request) {
	//w is how you send respnses
	//req is how requests come in
	//Fmt.Fprintln - writes to the browser instead of terminal
	/*fmt.Fprintln(w, `
	<!DOCTYPE html>
	<html>
		<head>
			<title>ASCII Art</title>
		</head>
		<body>
			<h1>ASCII Art Generator</h1>
			<p>Welcome to the ASCII art web app</p>
		</body>
	</html>
	`)
	*/

	templ, err := template.ParseFiles("template/index.html")
	if err != nil {
		return
	}
	templ.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", homeHandler)
	fmt.Println("server running on port 8080")

	http.ListenAndServe(":8080", nil)
}
