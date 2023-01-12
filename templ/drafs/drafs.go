package main

import (
	"fmt"
	// "io/fs"
	// "text/template"
	"html/template"
	"net/http"
)

type User struct {
	Name                  string
	Age                   uint16
	Money                 int16
	Avg_grades, Happiness float64
	Hobbies               []string
}

func homepage(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -400, 4.2, 0.8, []string{"Volleyball", "Football", "Skate"}}
	// fmt.Fprintf(w, bob.Name)
	templ, err := template.ParseFiles("templ/home.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	templ.Execute(w, bob)
}
func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	fmt.Println("Hello")
	http.HandleFunc("/", homepage)
	http.ListenAndServe(":2020", nil)
}
