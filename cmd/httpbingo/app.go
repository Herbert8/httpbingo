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

//go:embed doc/html/manual.html
var sHelp string

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

	helpFunc := func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte(sHelp))
	}
	http.HandleFunc("/", helpFunc)
	http.HandleFunc("/help", helpFunc)

	// 启动服务
	sPort := fmt.Sprintf(":%d", *nPort)
	fmt.Printf("Starting server on port %d...\n", *nPort)
	startServeErr := http.ListenAndServe(sPort, nil)
	if startServeErr != nil {
		log.Println(startServeErr)
	}

}
