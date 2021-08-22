package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type VariableMap map[string]interface{}

func values2Map(values url.Values) VariableMap {

	ret := make(VariableMap)

	for k, v := range values {
		if len(v) > 1 {
			ret[k] = v
		} else {
			ret[k] = v[0]
		}
	}
	return ret
}

func parsePathParams(fullPathStr string, basePath string) []string {
	paramPathStr := strings.Replace(fullPathStr, basePath, "", 1)
	originalParamArr := strings.Split(paramPathStr, "/")

	var retParamArr []string
	for _, param := range originalParamArr {
		if param != "" {
			retParamArr = append(retParamArr, param)
		}
	}
	return retParamArr
}


func writeAccessControl(responseWriter http.ResponseWriter) {
	responseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	responseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
}

func writeJSONResponse(data interface{}, responseWriter http.ResponseWriter) {
	// 指定 Response 的 Content-Type
	responseWriter.Header().Set("Content-Type", "application/json")

	// 收集到的信息转 JSON 返回给客户端
	jsonBytesRet, err := json.MarshalIndent(data, "", "  ")

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

type HTTPRequestHandler func (http.ResponseWriter, *http.Request)

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