package url

import (
	"log"
	"github.com/ant0ine/go-json-rest/rest"
	"ufleet/cluster/apiserver"
)

func Register() (*rest.Api, error){

	restapi := rest.NewApi()
	restapi.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get(apiserver.ClusterPath, apiserver.GetClusters),
	)

	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	restapi.SetApp(router)
	return restapi, nil
}