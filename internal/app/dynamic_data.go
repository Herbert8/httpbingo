package service

import (
	_ "embed"
	"encoding/base64"
	"fmt"
	"httpbingo/internal/httputils"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func ProcBase64(w http.ResponseWriter, r *http.Request) {

	var base64Str string

	if strings.EqualFold("GET", r.Method) {
		paramArr := parsePathParams(r.URL.Path, 1)
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

	pathParams := parsePathParams(r.URL.Path, 1)

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

	// 解析表单
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		httputils.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 结果数据
	var retData []byte

	// 文本框内容
	sContent := r.FormValue("content")

	// 获取文件句柄
	if contentFile, _, err := r.FormFile("contentFile"); err != nil {
		// 如果读取文件报错，判断文本框内容是否为空
		// 如果文本框内容不为空，则使用文本框内容作为返回数据
		if sContent != "" {
			retData = []byte(sContent)
		} else {
			// 如果文本框内容也为空，则报错
			httputils.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		// 处理正常读取文件的情况
		defer func() {
			_ = contentFile.Close()
		}()
		// 读取文件内容
		if retData, err = io.ReadAll(contentFile); err != nil {
			httputils.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// 确定 Response Body
	responseBodyData := retData

	// Content-Type
	sContentType := r.FormValue("contentType")

	// Content-Type 默认值 application/octet-stream
	if sContentType == "unknown" {
		sContentType = "application/octet-stream"
	}

	// 如果 Content-Type 指定为 auto，或者没有指定，则根据返回内容自动检测
	if sContentType == "" || sContentType == "auto" {
		sContentType = http.DetectContentType(responseBodyData)
	}

	// 指定 Response 的 Content-Type
	w.Header().Set("Content-Type", sContentType)

	// 判断是否启用下载
	if r.FormValue("downloadSwitch") != "" {

		// 处理 Content-Type 为 text/html; charset=utf8; 的情况
		// 只取分号前面的 MIME Type 部分
		sMIMEType, _, _ := strings.Cut(sContentType, ";")
		// MIME Type 对应的中文描述
		sMIMETypeDescription := mimeTypeDictionary[sMIMEType]

		// 遍历对应扩展名
		var sFileExt string
		for k, v := range fileExtNameDictionary {
			if v == sMIMEType {
				sFileExt = k
				break
			}
		}
		// 默认扩展名为 dat
		if sFileExt == "" {
			sFileExt = "dat"
		}

		// 获取用户指定的下载文件名
		sFilename := r.FormValue("downloadFilename")
		sFilename = strings.TrimSpace(sFilename)
		// 如果没有指定
		if sFilename == "" {
			// 则使用 MIME Type 说明作为主文件名
			if sMIMETypeDescription != "" {
				sFilename = sMIMETypeDescription
			} else {
				// 如果 sMIMETypeDescription 也为空，则使用默认值
				sFilename = "模拟数据"
			}
		}
		// 与默认扩展名组合
		sFilename = fmt.Sprintf("%s.%s", sFilename, sFileExt)

		// 文件名 url 编码
		sFilename = url.QueryEscape(sFilename)
		// 通过 Response Header 指定下载信息
		// 指定文件名，正常情况根据指定文件名及编码进行下载；对于不支持编码的情况，采用 ASCII 文件名 file.dat
		sContentDisposition := fmt.Sprintf("attachment; filename=\"file.dat\"; filename*=utf-8''%s", sFilename)

		// 指定 Response 的 Content-Disposition 头
		w.Header().Set("Content-Disposition", sContentDisposition)
	}
	_, _ = w.Write(responseBodyData)
}

//go:embed resource/web/content_uploader.html
var dataContentUploaderHtml []byte

func init() {
	http.HandleFunc("/data/config", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write(dataContentUploaderHtml)
	})
}
