package service

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func ProcBase64(w http.ResponseWriter, r *http.Request) {

	var base64Str string

	if strings.EqualFold("GET", r.Method) {
		paramArr := parsePathParams(r.URL.Path, "/base64")
		if len(paramArr) > 0 {
			base64Str = paramArr[0]
		}
	}

	if strings.EqualFold("POST", r.Method) {
		_ = r.ParseForm()
		base64Str = r.FormValue("base64")
	}

	dataBytes, decodeErr := base64.StdEncoding.DecodeString(base64Str)

	if decodeErr != nil {
		w.Header().Set("Content-Type", "text/html")
		_, _ = fmt.Fprintf(w, "Incorrect Base64 data. <p>Try: <pre>SFRUUEJJTl9HTyBpcyBhd2Vzb21l</pre>")
	}

	sContentType := http.DetectContentType(dataBytes)
	w.Header().Set("Content-Type", sContentType)
	if sContentType == "application/octet-stream" {
		w.Header().Set("Content-Disposition", "attachment; filename=\"data.bin\"")
	}

	_, _ = w.Write(dataBytes)
}

func ProcDelay(w http.ResponseWriter, r *http.Request) {

	pathParams := parsePathParams(r.URL.Path, "/delay/")

	var sParam string
	if len(pathParams) > 0 {
		sParam = pathParams[0]
	} else {
		sParam = "3"
	}

	delaySeconds, convErr := strconv.Atoi(sParam)
	if convErr != nil {
		delaySeconds = 3
	}
	if delaySeconds > 10 {
		delaySeconds = 10
	}

	time.Sleep(time.Duration(delaySeconds) * time.Second)

	ProcAnything(w, r)
}



func ProcData(w http.ResponseWriter, r *http.Request) {

	// 内容
	sContent := r.FormValue("content")

	// 指定通过文件获取内容
	sDataFile := r.FormValue("content-file")

	// 如果指定了 content-file，则从指定文件读取内容
	var fileData []byte
	if sDataFile != "" {
		tmpFileData, readFileErr := ioutil.ReadFile(sDataFile)
		if readFileErr != nil {
			http.Error(w, readFileErr.Error(), http.StatusNotFound)
			return
		} else {
			fileData = tmpFileData
		}
	}

	// 确定 Response Body
	var responseBodyData []byte
	if len(fileData) == 0 {
		responseBodyData = []byte(sContent)
	} else {
		responseBodyData = fileData
	}

	// Content-Type
	sContentType := r.FormValue("content-type")
	// Content-Type 默认值 application/octet-stream
	if sContentType == "" {
		sContentType = "application/octet-stream"
	}
	// 如果 Content-Type 指定为 auto，则根据返回内容自动检测
	if sContentType == "auto" {
		sContentType = http.DetectContentType(responseBodyData)
	}

	w.Header().Set("Content-Type", sContentType)

	_, _ = w.Write(responseBodyData)
}


func ProcDownload(w http.ResponseWriter, r *http.Request) {

	// 指定默认文件名
	sFilename := r.FormValue("filename")
	if sFilename == "" {
		sFilename = "测试.dat"
	}

	// 文件名 url 编码
	sFilename = url.QueryEscape(sFilename)
	// 通过 Response Header 指定下载信息
	// 指定文件名，正常情况根据指定文件名及编码进行下载；对于不支持编码的情况，采用 ASCII 文件名 file.dat
	sContentDisposition := fmt.Sprintf("attachment; filename=\"file.dat\"; filename*=utf-8''%s", sFilename)

	w.Header().Set("Content-Disposition", sContentDisposition)

	ProcData(w, r)
}



func ProcDetect(w http.ResponseWriter, r *http.Request) {
	type DataInfo struct {
		Size        int    `json:"size"`
		ContentType string `json:"Content-Type"`
		Content     string `json:"content"`
	}
	// 读取 body 数据
	dataBytes, _ := ioutil.ReadAll(r.Body)

	// 检测或指定数据类型
	sContentType := http.DetectContentType(dataBytes)

	var sContent string
	var displayBytes []byte
	if len(dataBytes) > 50 {
		displayBytes = dataBytes[:50]
		sContent = string(displayBytes) + "..."
	} else {
		sContent = string(dataBytes)
	}

	dataInfo := DataInfo{
		Size: len(dataBytes),
		ContentType: sContentType,
		Content: sContent,
	}

	writeJSONResponse(dataInfo, w)
}
