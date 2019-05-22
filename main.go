package main

import (
	"html/template"
	"log"
	"mailit/smtp"
	"mailit/web"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	data, err := reddit.GetTopPosts()
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	t := template.Must(template.ParseFiles("web/html/index.html"))
	//If errors show an internal server error message
	//I also pass the welcome struct to the welcome-template.html file.
	if err := t.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func sender() {
	email.Send("test")
}

func main() {

	log.Println("Emailing..")
	sender()

	http.Handle("/web/",
		http.StripPrefix("/web/",
			http.FileServer(http.Dir("web"))))

	log.Println("Starting web server...")
	http.HandleFunc("/reddit/aww", handler)
	log.Fatal(http.ListenAndServe(":8089", nil))

}
