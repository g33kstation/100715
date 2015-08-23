package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl map[string]*template.Template

func init() {
	// Pré-rendu du template
	if tmpl == nil {
		tmpl = make(map[string]*template.Template)
	}
}

// Data est un struct qui imite le resultat en provenance de la BDD
//TODO: remplacer le struct par la fonction `getPerson(id)` (@LeMeteore)
type Data struct {
	PageTitle, ArticleTitle string
	ArticleBody             []byte
}

var patt = Data{
	"About LeMeteore",
	"It's all about a gr8 pythonistas",
	[]byte("lorem ipsum dolor sit amet..."),
}

func runServer() {
	log.Println("Server started on :4400...")
	http.HandleFunc("/", homepage)
	http.ListenAndServe(":4400", nil)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	tmpl["index"] = template.Must(template.ParseFiles("tpl/layout.html", "tpl/index.html"))

	//w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//createTable()
	//insertPerson()
	if err := tmpl["index"].Execute(w, patt); err != nil {
		log.Fatal("Tpl: ", err)
	}
}
