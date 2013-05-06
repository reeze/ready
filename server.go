package ready

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Request struct {
	http.Request
}

type Response struct {
	http.Response
}

type Server struct {
	host    string
	port    int
	docRoot string
}

func (s Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// TODO Move static request to static handler
	uri := req.RequestURI

	if uri == "/" {
		uri = "/index.html"
	}
	file := fmt.Sprintf("%s%s", s.docRoot, uri)

	contents, err := ioutil.ReadFile(file)

	fmt.Printf("Serve uri: %s for file: %s\n", uri, file)

	if err != nil {
		http.NotFound(w, req)
	} else {
		var ctnType = "text/plain"
		if strings.HasSuffix(file, ".png") {
			ctnType = "image/png"
		} else if strings.HasSuffix(file, ".css") {
			ctnType = "text/css"
		} else if strings.HasSuffix(file, ".js") {
			ctnType = "text/javascript"
		} else if strings.HasSuffix(file, ".html") {
			ctnType = "text/html"
		}
		w.Header().Set("Content-Type", ctnType)
		w.Write(contents)
	}
}

func NewServer(host string, port int, docRoot string) Server {
	server := Server{
		host:    host,
		port:    port,
		docRoot: docRoot,
	}

	return server
}

func (s Server) run() {
	addr := fmt.Sprintf("%s:%d", s.host, s.port)

	err := http.ListenAndServe(addr, s)

	if err != nil {
		log.Fatal("Failed to serve request")
	}
}
