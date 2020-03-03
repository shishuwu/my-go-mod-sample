package main

import (
	"io"
	"net/http"
)
import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	logSomething()

	port := kingpin.Flag("port", "Port Number to listen").Default("9999").Short('p').String()
	kingpin.Parse()
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	//r.HandleFunc("/products", ProductsHandler)
	//r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)
	logrus.Info(*port)
	http.ListenAndServe("0.0.0.0:"+*port, r)
}

func logSomething() {
	log.Info("This is a log line")
	log.Warn("Another log line")
	// log.Fatal("This is really bad, exiting...")
}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello World!")
	logrus.Info(request.RemoteAddr)
}
