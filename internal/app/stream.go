package service

import (
	"fmt"
	"net/http"
	"time"
)

func ProcStream(w http.ResponseWriter, r *http.Request) {

	// 设置 Content-Type 为 text/plain 或其他适当的值
	w.Header().Set("Content-Type", "text/plain")

	// 开始向客户端发送数据
	for i := 0; i < 10; i++ {
		// 模拟一些数据
		data := fmt.Sprintf("Data %d\n", i)

		// 将数据写入响应流
		_, err := fmt.Fprint(w, data)
		if err != nil {
			// 处理错误，例如连接中断等
			fmt.Println("Error:", err)
			return
		}

		// 强制刷新响应流
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}

		// 模拟延迟，以便观察流式传输效果
		time.Sleep(1 * time.Second)
	}

}
