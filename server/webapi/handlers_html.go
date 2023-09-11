package webapi

import (
	"fmt"
	"net/http"
	"text/template"
)

type WebPage struct {
	Title      string
	Servername string
}

func handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := &WebPage{Title: "home",
			Servername: "MTGO"}

		tmpl, err := template.New("").ParseFiles("webtemplates/home.html", "webtemplates/base.html")
		// check your err
		err = tmpl.ExecuteTemplate(w, "base", p)
		if err != nil {
			fmt.Errorf("Cannot load template. err=%s", err)
		}
	}
}

func handleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		p := &WebPage{Title: "login",
			Servername: "MTGO"}

		tmpl, err := template.New("").ParseFiles("webtemplates/account/login.html", "webtemplates/base.html")
		// check your err
		err = tmpl.ExecuteTemplate(w, "base", p)
		if err != nil {
			fmt.Errorf("Cannot load template. err=%s", err)
		}
	}
}

func handleRegister() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("").ParseFiles("webtemplates/account/register.html", "webtemplates/base.html")
		// check your err
		err = tmpl.ExecuteTemplate(w, "base", nil)
		if err != nil {
			fmt.Errorf("Cannot load template. err=%s", err)
		}
	}
}

func handleSettings() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("").ParseFiles("webtemplates/account/settings.html", "webtemplates/base.html")
		// check your err
		err = tmpl.ExecuteTemplate(w, "base", nil)
		if err != nil {
			fmt.Errorf("Cannot load template. err=%s", err)
		}
	}
}
