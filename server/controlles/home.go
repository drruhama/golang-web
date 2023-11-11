package controlles

import (
	"go-web-template/server/utils"
	"html/template"
	"log"
	"net/http"
	"path"
)

type HomeController interface {
	Index(w http.ResponseWriter, r *http.Request)
}

type homeController struct{}

func NewHomeController() *homeController {
	return &homeController{}
}

func (*homeController) Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("static", "pages/employees/index.html"), utils.LayoutMaster)
	if err != nil {
		log.Panicln(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
