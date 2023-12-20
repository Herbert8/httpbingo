package service

import (
	_ "embed"
	"encoding/json"
	"net/http"
)

//go:embed resource/data/utf8.html
var utf8Data []byte

//go:embed resource/data/html.html
var htmlData []byte

//go:embed resource/data/json.json
var jsonData []byte

//go:embed resource/data/xml.xml
var xmlData []byte

//go:embed resource/data/jpeg.jpg
var jpegData []byte

//go:embed resource/data/dragon.png
var pngData []byte

//go:embed resource/data/svg.svg
var svgData []byte

//go:embed resource/data/webp.webp
var webpData []byte

//go:embed resource/data/gif.gif
var gifData []byte

func ProcUTF8(responseWriter http.ResponseWriter, r *http.Request) {
	_, _ = writeByteSliceToResponse(responseWriter, utf8Data)
}

func ProcHTML(responseWriter http.ResponseWriter, r *http.Request) {
	_, _ = writeByteSliceToResponse(responseWriter, htmlData)
}

func ProcJSON(responseWriter http.ResponseWriter, r *http.Request) {
	_, _ = writeByteSliceToResponseWithContentType(responseWriter, jsonData, "application/json")
}

func ProcXML(responseWriter http.ResponseWriter, r *http.Request) {
	_, _ = writeByteSliceToResponse(responseWriter, xmlData)
}

func ProcJPEG(responseWriter http.ResponseWriter, r *http.Request) {
	_, _ = writeByteSliceToResponse(responseWriter, jpegData)
}

func ProcPNG(responseWriter http.ResponseWriter, r *http.Request) {
	_, _ = writeByteSliceToResponse(responseWriter, pngData)
}

func ProcWebP(responseWriter http.ResponseWriter, r *http.Request) {
	_, _ = writeByteSliceToResponse(responseWriter, webpData)
}

func ProcGif(responseWriter http.ResponseWriter, r *http.Request) {
	_, _ = writeByteSliceToResponse(responseWriter, gifData)
}

func ProcSVG(responseWriter http.ResponseWriter, r *http.Request) {
	_, _ = writeByteSliceToResponseWithContentType(responseWriter, svgData, "image/svg+xml")
}

func ProcGzip(responseWriter http.ResponseWriter, r *http.Request) {
	type GzippedRequestInfo struct {
		Gzipped bool `json:"gzipped"`
		RequestInfo
	}

	requestInfo := ProcAnythingInfo(r)
	gzippedRequestInfo := GzippedRequestInfo{
		RequestInfo: requestInfo,
		Gzipped:     true,
	}
	gzippedAnythingJsonData, err := json.MarshalIndent(gzippedRequestInfo, "", "  ")
	if err != nil {
		return
	}
	compressedData, _ := GzipCompress(gzippedAnythingJsonData)
	responseWriter.Header().Set("Content-Encoding", "gzip")
	_, _ = writeByteSliceToResponseWithContentType(responseWriter, compressedData, "application/json")
}
