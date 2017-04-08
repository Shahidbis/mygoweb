package controllers

import(
	"net/http"
	"text/template"
	"viewmodels"
)

type categoriesController struct {
	template *template.Template
}

func (this *categoriesController) get(w http.ResponseWriter, req *http.Request){
	vm := viewmodels.GetGategories()

	w.Header().Add("Content-type", "text/html")
	this.template.Execute(w, vm)
}
