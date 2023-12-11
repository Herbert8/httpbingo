package service

func registerHandler() {
	handleWrapperFunc("/anything", ProcAnything)
	handleWrapperFunc("/anything/", ProcAnything)
	handleWrapperFunc("/cookies", ProcGetCookies)
	handleWrapperFunc("/cookies/set", ProcSetCookies)
	handleWrapperFunc("/cookies/set-detail/", ProcSetCookieDetail)
	handleWrapperFunc("/cookies/delete", ProcDelCookies)
	handleWrapperFunc("/redirect-to", ProcRedirectTo)
	handleWrapperFunc("/basic-auth/", ProcBasicAuth)
	handleWrapperFunc("/delay/", ProcDelay)
	handleWrapperFunc("/base64", ProcBase64)
	handleWrapperFunc("/base64/", ProcBase64)
	handleWrapperFunc("/data", ProcData)
	handleWrapperFunc("/download", ProcDownload)
	handleWrapperFunc("/detect", ProcDetect)
	handleWrapperFunc("/status", ProcStatus)
	handleWrapperFunc("/status/", ProcStatus)
	handleWrapperFunc("/response-headers", ProcResponseHeader)
	handleWrapperFunc("/encoding/utf8", ProcUTF8)
	handleWrapperFunc("/html", ProcHTML)
	handleWrapperFunc("/json", ProcJSON)
	handleWrapperFunc("/xml", ProcXML)
	handleWrapperFunc("/gzip", ProcGzip)
	handleWrapperFunc("/image/jpeg", ProcJPEG)
	handleWrapperFunc("/image/png", ProcPNG)
	handleWrapperFunc("/image/svg", ProcSVG)
	handleWrapperFunc("/image/webp", ProcWebP)
	handleWrapperFunc("/image/gif", ProcGif)
	handleWrapperFunc("/dump/request", ProcDumpRequest)
}

func init() {
	registerHandler()
}
