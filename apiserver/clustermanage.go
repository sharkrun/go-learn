package apiserver

import (
	"net/http"
	"github.com/ant0ine/go-json-rest/rest"

	"ufleet/cluster/worker"
	"log"
)

// return all clusters list

func GetClusters(w rest.ResponseWriter, r *rest.Request){
	log.Println("GetClusters...")
	clusters, found := worker.Getclusters()
	if !found {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteJson(ErrorResponse(404, "Not Found!"))
		return
	}
	w.WriteJson(clusters)
}