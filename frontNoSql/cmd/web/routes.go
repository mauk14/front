package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/book/:id", app.showBook)
	router.HandlerFunc(http.MethodGet, "/ebook/:id", app.list)
	router.HandlerFunc(http.MethodGet, "/login", app.loginGet)
	router.HandlerFunc(http.MethodPost, "/login", app.login)
	router.HandlerFunc(http.MethodGet, "/reg", app.regGet)
	router.HandlerFunc(http.MethodPost, "/reg", app.reg)
	router.HandlerFunc(http.MethodPost, "/books/create", nil)

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(router)
}
