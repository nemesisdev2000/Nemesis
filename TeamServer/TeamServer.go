package main

import (
	"fmt"
	"net/http"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/nemesisdev2000/Nemesis/TeamServer/ApiClient"
	"github.com/nemesisdev2000/Nemesis/TeamServer/AuthenticationServer"
)

func main() {

	fmt.Println("TeamServer Started")
	mainRouter := mux.NewRouter()
	authRouter := mainRouter.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/signup", AuthenticationServer.SignupHandler).Methods("POST")

	authRouter.HandleFunc("/signin", AuthenticationServer.SigninHandler).Methods("GET")

	authRouter.HandleFunc("/signout", AuthenticationServer.SignoutHandler).Methods("GET")

	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:8000"}))

	//adding functionanlity to communicate with the API server
	apiRouter := mainRouter.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/listen", ApiClient.ApiStartListener).Methods("POST")
	apiRouter.HandleFunc("/stoplistener", ApiClient.ApiStopListener).Methods("POST")

	server := &http.Server{
		Addr:    "127.0.0.1:5555",
		Handler: ch(mainRouter),
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error Booting the Server")
	}
}
