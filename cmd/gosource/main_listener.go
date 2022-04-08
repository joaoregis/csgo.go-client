package main

import (
	"gosource/internal/global/configs"
	"gosource/internal/global/logs"
	"net/http"

	"github.com/gorilla/mux"
)

func initListener() {

	router := mux.NewRouter()

	router.HandleFunc("/orchestrator/request-action/{action}", requestRefreshConfig).Methods("GET")

	logs.Error(http.ListenAndServe(":61975", router))

}

func requestRefreshConfig(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	action := params["action"]
	switch action {
	case "reload-cfg":
		configs.Reload()
	default:
		w.Write([]byte("[\"invalid action\"]"))
	}
}
