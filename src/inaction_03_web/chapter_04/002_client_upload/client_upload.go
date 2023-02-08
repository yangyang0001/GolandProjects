package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	targetURL := "http://localhost:9090/upload"
	//filename := "./运算器实现乘法的过程.png"
	filename := "./控制器.png"
	ClientUpload(filename, targetURL)
}

func ClientUpload(filename string, targetURL string) error {

	bodyBuff := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuff)

	// 模拟 html 上传, 创建一个 临时的 writer
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		log.Fatal(err)
		return err
	}

	openfile, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer openfile.Close()

	// 将 file 写入到 上面创建的 writer
	_, err = io.Copy(fileWriter, openfile)
	if err != nil {
		log.Fatal(err)
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	response, err := http.Post(targetURL, contentType, bodyBuff)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer response.Body.Close()

	resp_body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("response status = %v \n", response.Status)
	fmt.Printf("response body   = %v \n", string(resp_body))
	return nil

}
