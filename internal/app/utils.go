package service

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type VariableMap map[string]interface{}

func stringSlice2Value(strSlice []string) any {
	strSliceLen := len(strSlice)
	if strSliceLen > 1 {
		return strSlice
	} else if strSliceLen == 1 {
		return strSlice[0]
	} else {
		return ""
	}
}

func values2Map(values url.Values) VariableMap {
	ret := make(VariableMap)
	for k, v := range values {
		ret[k] = stringSlice2Value(v)
	}
	return ret
}

func parsePathParams(fullPathStr string, basePathComponentCount int) []string {
	basePathComponentCountWithRootPath := basePathComponentCount + 1
	// 拆分路径
	originalParamArr := strings.Split(fullPathStr, "/")
	if basePathComponentCountWithRootPath > len(originalParamArr)-1 {
		return nil
	}

	return originalParamArr[basePathComponentCountWithRootPath:]
}

func writeAccessControl(responseWriter http.ResponseWriter) {
	responseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	responseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
}

func writeJSONResponse(data interface{}, responseWriter http.ResponseWriter) {
	// 指定 Response 的 Content-Type
	responseWriter.Header().Set("Content-Type", "application/json")

	// 收集到的信息转 JSON 返回给客户端
	jsonBytesRet, err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(-1)
	}

	// 返回数据
	jsonStrRet := string(jsonBytesRet)
	jsonStrRet = strings.ReplaceAll(jsonStrRet, "\\u0026", "&")

	log.Printf("=============================================\n%s\n\n", jsonStrRet)

	_, _ = fmt.Fprint(responseWriter, jsonStrRet)
}

type HTTPRequestHandler func(http.ResponseWriter, *http.Request)

func handleWrapperFunc(pattern string, handler HTTPRequestHandler) {
	http.HandleFunc(pattern, ProcessorWrapper(handler))
}

func ProcessorWrapper(processor HTTPRequestHandler) HTTPRequestHandler {
	return func(writer http.ResponseWriter, request *http.Request) {
		writeAccessControl(writer)
		processor(writer, request)
	}
}

// ObtainIPs 获取本机 IP
func ObtainIPs() []string {
	var retIPs []string
	addrArr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return retIPs
	}
	for _, address := range addrArr {
		// 检查ip地址判断是否回环地址
		if netIP, ok := address.(*net.IPNet); ok && !netIP.IP.IsLoopback() {
			if netIP.IP.To4() != nil {
				retIPs = append(retIPs, netIP.IP.String())
			}
		}
	}
	return retIPs
}

func writeByteSliceToResponseWithContentType(responseWriter http.ResponseWriter, data []byte, contentType string) (int, error) {
	responseWriter.Header().Set("Content-Type", contentType)
	responseWriter.Header().Set("Content-Length", strconv.Itoa(len(data)))
	return responseWriter.Write(data)
}

func writeByteSliceToResponse(responseWriter http.ResponseWriter, data []byte) (int, error) {
	sContentType := http.DetectContentType(data)
	return writeByteSliceToResponseWithContentType(responseWriter, data, sContentType)
}

func GzipCompress(input []byte) ([]byte, error) {
	var compressedData bytes.Buffer

	// 创建一个 Gzip Writer，将数据写入其中
	writer := gzip.NewWriter(&compressedData)
	_, err := writer.Write(input)
	if err != nil {
		return nil, err
	}

	// 关闭 Gzip Writer，确保所有数据都被写入
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	return compressedData.Bytes(), nil
}
