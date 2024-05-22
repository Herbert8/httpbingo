package main

import (
	_ "embed"
	"flag"
	"fmt"
	service "httpbingo/internal/app"
	"httpbingo/internal/httputils/middleware"
	"log"
	"net/http"
	"os"
)

func main() {
	//showSummary()
	// 设置日志格式，显示代码行号
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	var nPort int
	var bHelp bool
	flag.BoolVar(&bHelp, "h", false, "Display this help screen")
	flag.IntVar(&nPort, "p", 8080, "Bind and listen for incoming requests")

	flag.Parse()

	service.ListenPort = nPort

	if bHelp {
		//showUsage()
		service.ShowSummary()
		os.Exit(0)
	}

	logMiddleware := middleware.LogMiddleware{}

	// 启动服务
	sPort := fmt.Sprintf(":%d", nPort)
	fmt.Printf("Starting server on port %d...\n", nPort)

	httpServer := http.Server{
		Addr:    sPort,
		Handler: &logMiddleware,
	}
	startServeErr := httpServer.ListenAndServe()
	//startServeErr := http.ListenAndServe(sPort, &logMiddleware)
	if startServeErr != nil {
		log.Println(startServeErr)
	}

}
