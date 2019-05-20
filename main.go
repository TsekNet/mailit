package main

import (
	"html/template"
	"log"
	//"mailit/smtp"
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
	//email.Send("mint ice cream with chocolate sprinkles")
}

func main() {

	// log.Println("Emailing..")
	// sender()

	http.Handle("/web/", //final url can be anything
		http.StripPrefix("/web/",
			http.FileServer(http.Dir("web")))) //Go looks in the relative "web" directory first using http.FileServer(), then matches it to a
	//url of our choice as shown in http.Handle("/web/"). This url is what we need when referencing our css files
	//once the server begins. Our html code would therefore be <link rel="stylesheet"  href="/web/stylesheet/...">
	//It is important to note the url in http.Handle can be whatever we like, so long as we are consistent.

	log.Println("Starting web server...")
	http.HandleFunc("/reddit/aww", handler)
	log.Fatal(http.ListenAndServe(":8089", nil))

}
