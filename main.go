package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/njosefbeck/janet.inputsenpai.com/views"
)

var homeView *views.View

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func main() {
	homeView = views.NewView("layout", "views/home.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	http.ListenAndServe(":3000", r)
}
