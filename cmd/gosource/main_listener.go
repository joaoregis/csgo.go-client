package main

import (
	"encoding/json"
	"gosource/internal/global/configs"
	"gosource/internal/global/logs"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func initListener() {

	router := mux.NewRouter()

	router.HandleFunc("/orchestrator/request-action/{action}", requestRefreshConfig_GET).Methods("GET")
	router.HandleFunc("/orchestrator/request-action/{action}", requestRefreshConfig_PUT).Methods("PUT")

	logs.Error(http.ListenAndServe(":61975", router))

}

func requestRefreshConfig_GET(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	action := params["action"]
	switch action {
	case "reload-cfg":
		configs.Reload()
	default:
		w.Write([]byte("[\"invalid action\"]"))
	}
}

func requestRefreshConfig_PUT(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	action := params["action"]
	switch action {
	case "update-option":

		bs, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(bs, &configs.G)
		logs.Info("config option updated")

	default:
		w.Write([]byte("[\"invalid action\"]"))
	}
}
