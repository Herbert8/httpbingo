package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)
import _ "embed"

//go:embed doc/html/manual.html
var helpData []byte

//go:embed doc/html/redoc.html
var redocData []byte

//go:embed doc/html/swagger.html
var swaggerData []byte

//go:embed doc/html/rapidoc/*
var docData embed.FS

// 此时 docData 映射到 doc 所在目录，即 doc 的上层

func init() {
	helpFunc := func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write(helpData)
	}
	//http.HandleFunc("/", helpFunc)
	http.HandleFunc("/help", helpFunc)

	redocFunc := func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write(redocData)
	}
	http.HandleFunc("/redoc", redocFunc)

	swaggerFunc := func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write(swaggerData)
	}
	http.HandleFunc("/swagger", swaggerFunc)

	// 获取 rapiDoc 的 FS
	rapiDocFS, err := fs.Sub(docData, "doc/html/rapidoc")
	if err != nil {
		log.Fatalln(err.Error())
	}
	// rapiDocFS 转 httpFS
	httpFS := http.FS(rapiDocFS)
	// httpFS 转 HTTP File Server Handler
	rapiDocHandler := http.FileServer(httpFS)
	http.Handle("/", http.StripPrefix("/", rapiDocHandler))
}
