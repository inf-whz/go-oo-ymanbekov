package main

import (
	"fmt"
	"net/http"
	"html/template"

)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html",
		"templates/registration_form.html", "templates/footer.html", "templates/login.html")
	if err != nil{
		fmt.Fprint(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)

}

func main(){
	fmt.Println("port 3000")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":3000", nil)

}