package httpserver

import (
	"fmt"
	"http-replay/cmd"
	"log"
	"net/http"
)

type Response struct {
	Body string
}

func (self *Response) UpdateBody(body string) {
	self.Body = body
}

type HttpServer struct {
	ListenAddress string
	MetricsPath   string
	Response      *Response
}

func NewHttpServer(myargs *cmd.CmdArguments) (server HttpServer) {
	server = HttpServer{ListenAddress: myargs.ListenAddress, MetricsPath: myargs.MetricsPath, Response: &Response{Body: ""}}
	return server
}

func (server HttpServer) Serve() {
	http.HandleFunc(server.MetricsPath, server.HelloHandler)
	fmt.Printf("Server started at port %s\n", server.ListenAddress)
	log.Fatal(http.ListenAndServe(server.ListenAddress, nil))
}

func (server HttpServer) HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, server.Response.Body)
}

func (server HttpServer) UpdateResponse(responseString string) {
	server.Response.UpdateBody(responseString)
}
