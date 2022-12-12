package lib

import (
	"bufio"
	"io"
	"log"
	"net/http"
)

func Post(req *http.Request, c chan<- string) error {
	// 创建一个 HTTP 客户端
	client := &http.Client{}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	defer close(c)
	reader := bufio.NewReader(resp.Body)
	// 循环读取响应流
	for {
		// 读取一行数据
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				// 如果已经读取到了响应的结尾，则退出循环
				break
			}
			log.Fatal(err)
		}
		if line == "" {
			continue
		}
		c <- line
	}

	return nil
}
