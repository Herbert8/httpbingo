package service

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

var ServicePort int

type FileInfo struct {
	FileName   string      `json:"filename"`
	MIMEHeader VariableMap `json:"mime-header"`
	Size       int64       `json:"size"`
	Content    string      `json:"content"`
}

type RequestInfo struct {
	Args  VariableMap `json:"args"`
	Data  string      `json:"data"`
	Files VariableMap `json:"files"`
	Form  VariableMap `json:"form"`
	//MultiPart VariableMap `json:"multi_part"`
	Headers         VariableMap `json:"headers"`
	JSON            VariableMap `json:"json"`
	Method          string      `json:"method"`
	URL             string      `json:"url"`
	ServerEndpoints []string    `json:"server_endpoints"`
	Client          string      `json:"client"`
}

func fileHeader2FileInfo(fileHeader *multipart.FileHeader) FileInfo {
	var retFileInfo FileInfo
	retFileInfo.FileName = fileHeader.Filename
	retFileInfo.Size = fileHeader.Size
	if fileReader, err := fileHeader.Open(); err == nil {
		if content, errRead := io.ReadAll(fileReader); errRead == nil {
			retFileInfo.Content = string(content)
		}
	}
	retFileInfo.MIMEHeader = values2Map(url.Values(fileHeader.Header))
	return retFileInfo
}

func fileHeaderSlice2FileInfos(fileHeaderSlice []*multipart.FileHeader) any {
	fileHeaderSliceLen := len(fileHeaderSlice)
	retFileInfos := make([]FileInfo, fileHeaderSliceLen)
	for idx, fileHeader := range fileHeaderSlice {
		retFileInfos[idx] = fileHeader2FileInfo(fileHeader)
	}
	if fileHeaderSliceLen > 1 {
		return retFileInfos
	} else if fileHeaderSliceLen == 1 {
		return retFileInfos[0]
	} else {
		return nil
	}
}

func fileHeaderMap2FileInfoMap(fileHeaderMap map[string][]*multipart.FileHeader) VariableMap {
	retFileInfoMap := make(VariableMap)
	for k, v := range fileHeaderMap {
		retFileInfoMap[k] = fileHeaderSlice2FileInfos(v)
	}
	return retFileInfoMap
}

func ProcAnything(w http.ResponseWriter, r *http.Request) {

	// 创建收集 请求 信息的对象
	requestInfo := RequestInfo{}

	// URL 参数
	requestInfo.Args = values2Map(r.URL.Query())
	// 解析 Form
	//_ = r.ParseForm()
	_ = r.ParseMultipartForm(200)
	// 读取 Body
	bodyDataBytes, readErr := io.ReadAll(r.Body)
	if readErr == nil {
		requestInfo.Data = string(bodyDataBytes)
	}

	// 当 MultipartForm 和 MultipartForm.File 不为空，则认为有文件上传，显示文件信息
	if r.MultipartForm != nil && r.MultipartForm.File != nil {
		requestInfo.Files = fileHeaderMap2FileInfoMap(r.MultipartForm.File)
	} else {
		requestInfo.Files = make(VariableMap)
	}

	// 转换 Body
	requestInfo.Form = values2Map(r.PostForm)
	// 转换 Header
	requestInfo.Headers = values2Map(url.Values(r.Header))
	requestInfo.Headers["Host"] = r.Host

	// 方法
	requestInfo.Method = r.Method

	requestInfo.Client = r.RemoteAddr

	//fmt.Println("multiPart:", r.MultipartForm)

	// 如果 Body 是 JSON，则解析
	jsonMap := make(VariableMap)
	_ = json.Unmarshal(bodyDataBytes, &jsonMap)
	requestInfo.JSON = jsonMap

	// URL
	sUrl, urlErr := url.QueryUnescape(r.RequestURI)

	if urlErr == nil {
		requestInfo.URL = sUrl
	} else {
		requestInfo.URL = r.RequestURI
	}

	ipArr := ObtainIPs()
	for idx, val := range ipArr {
		ipArr[idx] = fmt.Sprintf("%s:%d", val, ServicePort)
	}
	requestInfo.ServerEndpoints = ipArr

	//r.RequestURI
	// 指定 Response 的 Content-Type
	w.Header().Set("Content-Type", "application/json")

	writeJSONResponse(requestInfo, w)
}
