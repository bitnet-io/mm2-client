/*
statsrv is a very simple static file server in go using the defaul net/http package
Usage:
	-p 1717               port to serve on
	-d ./stocks     the directory of static files to host
Go to http://localhost:1717 - it will display the index.html or directory listing file.
Place any .json file and it will be served as application/json type document. :)
Example:
mkdir -p /var/tmp/stats
echo '[{"type": "apache", "stats": [{"requests": 123212}, {"200": 123200}, {"4xx": 11}, {"5xx": 1}]}]' > /var/tmp/stats/stats.json
go run statsrv.go
curl -v http://localhost:8080/stats.json
*****************
Caution:  Currently no logging or any other safeguards, so suggest to never use this on a public accessible server !!
*****************
*/


package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var http_dir *string
	var http_port *string

	http_dir = flag.String("d", "./stocks", "stats directory to serve from")
	http_port = flag.String("p", "1717", "http port to listen on")
	flag.Parse()

	log.Println("Starting stats server with root dir: ", *http_dir)
	fs := http.FileServer(http.Dir(*http_dir))
	http.Handle("/", fs)
	log.Println("Listening on: ", *http_port)
	http.ListenAndServe(":"+*http_port, nil)
}
