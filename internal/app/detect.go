package service

import (
	_ "embed"
	"io"
	"net/http"
	"strings"
)

type DataInfo struct {
	Size        int    `json:"size"`
	ContentType string `json:"content-type"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

func data2DataInfo(dataBytes []byte) *DataInfo {
	// 检测或指定数据类型
	sContentType := http.DetectContentType(dataBytes)

	var sContent string
	var displayBytes []byte
	const MaxDataLen = 100
	if len(dataBytes) > MaxDataLen {
		displayBytes = dataBytes[:MaxDataLen]
		sContent = string(displayBytes) + "..."
	} else {
		sContent = string(dataBytes)
	}

	retDataInfo := new(DataInfo)
	retDataInfo.Size = len(dataBytes)
	retDataInfo.ContentType = sContentType
	if sNewContentType, _, found := strings.Cut(sContentType, ";"); found {
		sContentType = sNewContentType
	}
	retDataInfo.Description = mimeTypeDictionary[sContentType]
	retDataInfo.Content = sContent

	return retDataInfo
}

func detectFormUrlencoded(r *http.Request) (*DataInfo, error) {
	if !strings.Contains(r.Header.Get("Content-Type"), "application/x-www-form-urlencoded") {
		return nil, nil
	}

	// 读取 body 数据
	if dataBytes, err := io.ReadAll(r.Body); err == nil {
		return data2DataInfo(dataBytes), nil
	} else {
		return nil, err
	}
}

func detectMultipartFormData(r *http.Request) (*DataInfo, error) {
	if !strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data") {
		return nil, nil
	}

	// 解析表单数据，限制上传文件的大小
	err := r.ParseMultipartForm(10 << 20) // 10 MB限制
	if err != nil {
		return nil, err
	}

	// 获取文件句柄
	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = file.Close()
	}()

	if dataBytes, err := io.ReadAll(file); err == nil {
		return data2DataInfo(dataBytes), nil
	} else {
		return nil, err
	}
}

func ProcDetect(w http.ResponseWriter, r *http.Request) {
	var dataInfo *DataInfo
	var err error
	if dataInfo, err = detectMultipartFormData(r); dataInfo == nil {
		dataInfo, err = detectFormUrlencoded(r)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	writeJSONResponse(dataInfo, w)
}

//go:embed resource/web/detect_uploader.html
var detectUploader []byte

func init() {
	http.HandleFunc("/detect/uploader", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write(detectUploader)
	})
}
