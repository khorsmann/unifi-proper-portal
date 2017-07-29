package main

import (
	"github.com/gorilla/mux"
	"github.com/sequoiia/unifi-proper-portal/controller"
	"github.com/urfave/negroni"
	"net/http"
)

func main() {
	controller.LoadConfig() // Load config.yaml
	router := mux.NewRouter()

	router.Handle("/", negroni.New(
		negroni.Wrap(http.HandlerFunc(controller.Root)),
	))

	router.Handle("/users", negroni.New(
		negroni.Wrap(http.HandlerFunc(controller.UsersView)),
	))

	router.Handle("/users/auth/{clientid}", negroni.New(
		negroni.Wrap(http.HandlerFunc(controller.UsersAuthorisationApi)),
	))

	router.Handle("/users/delete/{clientid}", negroni.New(
		negroni.Wrap(http.HandlerFunc(controller.UsersDeleteApi)),
	))

	// fb
	router.Handle("/social/fb/auth", negroni.New(
		negroni.Wrap(http.HandlerFunc(controller.OAuthRedirect)),
	))

	n := negroni.New(negroni.NewRecovery())
	n.UseHandler(router)
	n.Run("0.0.0.0:8080")
}
