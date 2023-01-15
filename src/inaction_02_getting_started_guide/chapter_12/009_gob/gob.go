package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Mine struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

func main() {

	fmt.Println("---------------------------- 1、使用 bytes.Buffer 作为网络 实现 RPC ---------------------------")
	GobBuff()

	fmt.Println("---------------------------- 2、使用 gob.NewEncoder 写文件 -----------------------------------")
	GobWriteFile()

	fmt.Println("---------------------------- 3、使用 gob.NewDecoder 读文件 -----------------------------------")
	GobReadFile()

}

func GobBuff()  {

	// 用 bytes.Buffer 代表 网络传输的问题!
	var buff bytes.Buffer

	encoder := gob.NewEncoder(&buff)
	err := encoder.Encode(Mine{"zhangsan", "123456"})
	if err != nil {
		log.Fatal(err)
	}


	decoder := gob.NewDecoder(&buff)
	var mine Mine
	err = decoder.Decode(&mine)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("GobBuff method invoke, mine = %v \n", jsonParse(mine))

}

func GobWriteFile()  {

	filename := "/Users/yangjianwei/GolandProjects/src/inaction_02_getting_started_guide/chapter_12/009_gob/writefile.txt"

	openfile, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer openfile.Close()

	encoder := gob.NewEncoder(openfile)
	err = encoder.Encode(Mine{"xiaoming", "root"})
	if err != nil {
		log.Fatal(err)
	}

}

func GobReadFile()  {

	filename := "/Users/yangjianwei/GolandProjects/src/inaction_02_getting_started_guide/chapter_12/009_gob/writefile.txt"

	openfile, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer openfile.Close()

	var mine Mine
	decoder := gob.NewDecoder(openfile)
	err = decoder.Decode(&mine)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("GobReadFile method invoke, mine = %v \n", jsonParse(mine))
}

func jsonParse(i interface{}) string {

	marshal, err := json.Marshal(i)
	if err != nil {
		log.Fatal(err)
	}

	return string(marshal)

}
