package controllers

import (
	"net/http"
	"text/template"
)

func AccueilHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "accueil", nil)
}

func AchatHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "achat", nil)
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact", nil)
}

func EvenementHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "evenement", nil)
}

func CompteHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "compte", nil)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tpl, err := template.ParseFiles("views/base.html", "views/header.html", "views/footer.html", "views/"+tmpl+".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.ExecuteTemplate(w, "base", data)
}
