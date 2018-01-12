package main

import (
	"fmt"
	"log"
	"net/http"
	"ufleet/cluster/config"
	"os"
	"flag"
	"path"
	"github.com/ant0ine/go-json-rest/rest"
	"ufleet/cluster/utils"
	"strconv"
	//"ufleet/cluster/url"
	//"time"
	//"github.com/StephanDollberg/go-json-rest-middleware-jwt"
	"ufleet/cluster/url"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func handle_auth(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{"authed": r.Env["REMOTE_USER"].(string)})
}

func main() {
	etcdServer := flag.String("etcd", "http://" + config.ETCDHOST + ":" + strconv.Itoa(config.ETCDPORT), "etcd endpoint")
	serverPort := flag.String("server_port", strconv.Itoa(config.SERVER_PORT), "server port")

	// set log
	os.MkdirAll(path.Dir(config.LOGDIR), 0777)
	logPath := flag.String("log", config.LOGDIR, "log file path")
	logFile, err := os.OpenFile(*logPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
	log.SetFlags(config.GLogLevel)
	// parse params
	flag.Parse()

	// init log config and etcd server.
	utils.InitRunOption(*etcdServer)
	utils.Test()
	log.Println("start apiserver...")

	// register route
	restapi, err1 := url.Register()
	if err1 != nil {
		panic(err)
	}
	log.Println("Begin listening...")
	log.Fatal(http.ListenAndServe("localhost:" + *serverPort, restapi.MakeHandler()))
}