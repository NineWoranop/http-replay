package main

import (
	"fmt"
	"http-replay/cmd"
	"http-replay/file"
	"http-replay/httpserver"
	"time"
)

func timeStampMillis() int {
	return int(time.Now().UnixNano() / 1000000)
}

func main() {
	var err error
	var myargs *cmd.CmdArguments
	myargs, err = cmd.ReadArguments()
	if err != nil {
		fmt.Printf("Failed to start %s\n", err)
	} else {
		var handlerConfig *file.HanlderConfig = file.ToHanlderConfig(myargs)
		var fileHandler file.FileHandler = file.NewFileHandler(handlerConfig)
		var server httpserver.HttpServer = httpserver.NewHttpServer(myargs)
		var responseStr string
		responseStr, err = fileHandler.Read()
		if err != nil {
			fmt.Printf("Failed to read first file %s\n", err)
		} else {
			fmt.Printf("%s have been loaded\n", fileHandler.Config.CurrentFilePath())
			server.UpdateResponse(responseStr)
			go func() {
				for {
					myargs.Sleep()
					work(&fileHandler, &server)
				}
			}()
			server.Serve()
		}
	}
}

func work(fileHandler *file.FileHandler, server *httpserver.HttpServer) {
	if fileHandler.HasNext() {
		fileHandler.Next()
		httpResponseStr, err2 := fileHandler.Read()
		if err2 != nil {
			fmt.Printf("Failed to read %s\n", err2)
		} else {
			fmt.Printf("%s have been loaded\n", fileHandler.Config.CurrentFilePath())
			server.UpdateResponse(httpResponseStr)
		}
	}
}
