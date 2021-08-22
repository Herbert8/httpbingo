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

	bHelp := flag.Bool("h", false, "Display this help screen")
	nPort := flag.Int("p", 8080, "Bind and listen for incoming requests")

	flag.Parse()

	service.ServicePort = *nPort

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
	http.HandleFunc("/anything", service.ProcessorWrapper(service.ProcAnything))
	http.HandleFunc("/anything/", service.ProcessorWrapper(service.ProcAnything))
	http.HandleFunc("/cookies", service.ProcessorWrapper(service.ProcGetCookies))
	http.HandleFunc("/cookies/set", service.ProcessorWrapper(service.ProcSetCookies))
	http.HandleFunc("/cookies/set-detail/", service.ProcessorWrapper(service.ProcSetCookieDetail))
	http.HandleFunc("/cookies/delete", service.ProcessorWrapper(service.ProcDelCookies))
	http.HandleFunc("/redirect-to", service.ProcessorWrapper(service.ProcRedirectTo))
	http.HandleFunc("/basic-auth/", service.ProcessorWrapper(service.ProcBasicAuth))
	http.HandleFunc("/delay/", service.ProcessorWrapper(service.ProcDelay))
	http.HandleFunc("/base64", service.ProcessorWrapper(service.ProcBase64))
	http.HandleFunc("/base64/", service.ProcessorWrapper(service.ProcBase64))
	http.HandleFunc("/data", service.ProcessorWrapper(service.ProcData))
	http.HandleFunc("/download", service.ProcessorWrapper(service.ProcDownload))
	http.HandleFunc("/detect", service.ProcessorWrapper(service.ProcDetect))
	http.HandleFunc("/status", service.ProcessorWrapper(service.ProcStatus))
	http.HandleFunc("/status/", service.ProcessorWrapper(service.ProcStatus))
	http.HandleFunc("/response-headers", service.ProcessorWrapper(service.ProcResponseHeader))


	// 启动服务
	sPort := fmt.Sprintf(":%d", *nPort)
	fmt.Printf("Starting server on port %d...\n", *nPort)
	startServeErr := http.ListenAndServe(sPort, nil)
	if startServeErr != nil {
		log.Println(startServeErr)
	}
//TODO: 增加 /status/{codes}
//TODO: response-headers
//TODO: encoding/utf8
//TODO: gzip
}
