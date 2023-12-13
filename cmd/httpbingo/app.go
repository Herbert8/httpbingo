package main

import (
	_ "embed"
	"flag"
	"fmt"
	service "httpbingo/internal/app"
	"log"
	"net/http"
	"os"
)

func main() {

	//showSummary()
	// 设置日志格式，显示代码行号
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	bHelp := flag.Bool("h", false, "Display this help screen")
	nPort := flag.Int("p", 8080, "Bind and listen for incoming requests")

	flag.Parse()

	service.ListenPort = *nPort

	if *bHelp {
		//showUsage()
		service.ShowSummary()
		os.Exit(0)
	}

	// 启动服务
	sPort := fmt.Sprintf(":%d", *nPort)
	fmt.Printf("Starting server on port %d...\n", *nPort)
	startServeErr := http.ListenAndServe(sPort, nil)
	if startServeErr != nil {
		log.Println(startServeErr)
	}

}
