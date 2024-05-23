package service

import (
	_ "embed"
	"httpbingo/internal/httputils"
	"net/http"
)

//go:embed resource/file/word.doc
var word97Data []byte

//go:embed resource/file/word.docx
var wordData []byte

//go:embed resource/file/excel.xls
var excel97Data []byte

//go:embed resource/file/excel.xlsx
var excelData []byte

//go:embed resource/file/powerpoint.ppt
var ppt97Data []byte

//go:embed resource/file/powerpoint.pptx
var pptData []byte

//go:embed resource/file/pdf.pdf
var pdfData []byte

type fileDataInfo struct {
	mimeType string
	fileData []byte
}

var fileDataMap = map[string]fileDataInfo{
	"word97": {
		mimeType: "application/msword",
		fileData: word97Data,
	},
	"word": {
		mimeType: "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		fileData: wordData,
	},
	"excel97": {
		mimeType: "application/excel",
		fileData: excel97Data,
	},
	"excel": {
		mimeType: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		fileData: excelData,
	},
	"ppt97": {
		mimeType: "application/mspowerpoint",
		fileData: ppt97Data,
	},
	"ppt": {
		mimeType: "application/vnd.openxmlformats-officedocument.presentationml.presentation",
		fileData: pptData,
	},
	"pdf": {
		mimeType: "application/pdf",
		fileData: pdfData,
	},
}

func ProcFile(responseWriter http.ResponseWriter, r *http.Request) {
	pathParams := parsePathParams(r.URL.Path, 1)
	var fileType string
	if len(pathParams) > 0 {
		fileType = pathParams[0]
	}
	if fileInfo := fileDataMap[fileType]; fileInfo.fileData != nil {
		_, _ = writeByteSliceToResponseWithContentType(responseWriter, fileInfo.fileData, fileInfo.mimeType)
	} else {
		httputils.ErrorNotFound(responseWriter)
	}
}

func init() {
	handleWrapperFunc("/file/", ProcFile)
}
