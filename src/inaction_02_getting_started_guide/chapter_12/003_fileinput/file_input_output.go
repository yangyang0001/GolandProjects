package main

import (
	"archive/zip"
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	//fmt.Println("----------------------------------- 1、reader.ReadString -----------------------------------")
	//readLineFile()
	//
	//fmt.Println("----------------------------------- 2、reader.Read(....) ----------------------------------")
	//readBuffFile()
	//
	//fmt.Println("----------------------------------- 3、readColumnFile -------------------------------------")
	//readColumnFile()
	//
	//fmt.Println("----------------------------------- 4、readZipFile案例 -------------------------------------")
	//readZipFile()
	//
	//fmt.Println("----------------------------------- 5、readGZipFile案例 ------------------------------------")
	//readGZipFile()
	//
	//fmt.Println("----------------------------------- 5、writeFile  案例 ------------------------------------")
	//writeFile()
	//
	//fmt.Println("----------------------------------- 6、appendFile 案例 -------------------------------------")
	//appendFile()

	fmt.Println("----------------------------------- 7、copyFile   案例 -------------------------------------")
	copyFile()

}


func fileExist(path string) bool {
	_, err := os.Stat(path)

	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}

	return true
}

/**
	reader 使用 reader.ReadString() 方法
 */
func readLineFile()  {
	filename := "/Users/yangjianwei/GolandProjects/src/inaction_02_getting_started_guide/chapter_12/003_fileinput/read.txt"

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		fmt.Printf("line = %v", line)

		if err == io.EOF {
			fmt.Println("read file end!")
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}

}

/**
	reader 使用 reader.Read(buff []byte) 方法
*/
func readBuffFile() {

	filename := "/Users/yangjianwei/GolandProjects/src/inaction_02_getting_started_guide/chapter_12/003_fileinput/read.txt"

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var buff = make([]byte, 1024)

	reader := bufio.NewReader(file)

	for {
		n, err := reader.Read(buff)

		fmt.Printf("n = %v, result = %v \n", n, string(buff[0:n]))

		if err == io.EOF {
			fmt.Println("read file end!")
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}

/**
	columnfile.txt 中 分列读取!
 */
func readColumnFile()  {

	columnname := "/Users/yangjianwei/GolandProjects/src/inaction_02_getting_started_guide/chapter_12/003_fileinput/columnfile.txt"
	columnfile, err := os.Open(columnname)

	if err != nil {
		log.Fatal(err)
	}

	defer columnfile.Close()

	reader := bufio.NewReader(columnfile)

	var col0, col1, col2 []string

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		fields := strings.Fields(line)
		col0 = append(col0, fields[0])
		col1 = append(col1, fields[1])
		col2 = append(col2, fields[2])

	}

	for col0_idx, col0_val := range col0 {
		fmt.Printf("col0_val = %v, col0_idx = %v \n", col0_val, col0_idx)
	}

	for col1_idx, col1_val := range col1 {
		fmt.Printf("col1_val = %v, col1_idx = %v \n", col1_val, col1_idx)
	}

	for col2_idx, col2_val := range col2 {
		fmt.Printf("col2_val = %v, col2_idx = %v \n", col2_val, col2_idx)
	}

}


/**
	TODO:
		go 语言对zip方法定义和使用, 不如 java 设计的精确, 更有哲学层次感! 看来 装饰者模式 不仅 网络原理中用, 在其他语言中还是要遵循的, 不是实现了就完了!
		虽然 这里也符合, 这个方法名称 真的体现不出 发明语言者 该具备的水平!
*/
func readZipFile()  {

	filename := "/Users/yangjianwei/GolandProjects/src/inaction_02_getting_started_guide/chapter_12/003_fileinput/input.zip"

	// 打开一个zip格式文件
	zipRCloser, err := zip.OpenReader(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer zipRCloser.Close()

	var buff = make([]byte, 1024)

	// 迭代压缩文件中的文件, 打印出文件中的内容
	for _, file := range zipRCloser.File {
		if strings.HasPrefix(file.Name, "__MACOSX") {
			continue
		} else {
			fmt.Printf("read zip file method invoke, filename = %v \n", file.Name)
			openRCloser, err := file.Open()
			if err != nil {
				log.Fatal(err)
			}
			defer openRCloser.Close()

			for {
				/**
					TODO: 理论上, 按照 java 的方式, 就直接使用 当前接口中的 Read() 方法读取就可以了, 但是这里不行!
					n, err := openRCloser.Read(buff)
					if err != nil {
						log.Fatal(err.Error())
					}
				 */

				reader := bufio.NewReader(openRCloser)
				n, err := reader.Read(buff)
				if err == io.EOF {
					break
				}

				fmt.Printf("n = %v \n", n)
				fmt.Printf(string(buff[0:n]))
			}
		}
	}

}


/**
	TODO: go 语言就不是处理上层动作的语言, 这的 tar.gz 文件就不能是过个文件的, 否则读取会有问题! 可以读取单个文件!
 */
func readGZipFile()  {
	filename := "/Users/yangjianwei/GolandProjects/src/inaction_02_getting_started_guide/chapter_12/003_fileinput/readgzipfile.tar.gz"
	gzipfile, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer gzipfile.Close()

	var reader * bufio.Reader

	gzReader, err := gzip.NewReader(gzipfile)

	if err != nil {
		reader = bufio.NewReader(gzipfile, )
	} else {
		reader = bufio.NewReader(gzReader)
	}

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}
		fmt.Printf("line = %v \n", line)
	}

}

func writeFile() {

	outname := "/Users/yangjianwei/GolandProjects/src/inaction_02_getting_started_guide/chapter_12/003_fileinput/write.txt"
	outfile, err := os.OpenFile(outname, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	defer outfile.Close()

	writer := bufio.NewWriter(outfile)

	for i := 0; i < 10; i++ {
		writer.WriteString("i = " + strconv.Itoa(i) + "\n")
	}
	writer.Flush()

}


func appendFile()  {

	inputname := "/Users/yangjianwei/GolandProjects/src/inaction_02_getting_started_guide/chapter_12/003_fileinput/input.txt"
	outpuname := "/Users/yangjianwei/GolandProjects/src/inaction_02_getting_started_guide/chapter_12/003_fileinput/output.txt"

	inputfile, err := os.Open(inputname)

	if err != nil {
		log.Fatal(err)
	}

	defer inputfile.Close()

	reader := bufio.NewReader(inputfile)

	var outpufile * os.File

	if fileExist(outpuname) {
		outpufile, err = os.OpenFile(outpuname, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("open outpufile error is ", err)
			return
		}
	} else {
		outpufile, err = os.Create(outpuname)
		if err != nil {
			fmt.Println("create outpufile error is ", err)
			return
		}
	}

	defer outpufile.Close()

	var bytes = make([]byte, 1024)
	writer := bufio.NewWriter(outpufile)

	for {
		read_num, err := reader.Read(bytes)

		if err == io.EOF {
			break
		}

		_, err = writer.Write(bytes[0:read_num])

		if err != nil {
			break
		}

		writer.Flush()
	}

}


func copyFile() {
	sourcename := "/Users/yangjianwei/GolandProjects/src/inaction_02_getting_started_guide/chapter_12/003_fileinput/sourcefile.txt"
	targetname := "/Users/yangjianwei/GolandProjects/src/inaction_02_getting_started_guide/chapter_12/003_fileinput/targetfile.txt"

	sourcefile, err := os.Open(sourcename)
	if err != nil {
		log.Fatal(err)
	}
	defer sourcefile.Close()

	targetfile, err := os.OpenFile(targetname, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer targetfile.Close()

	io.Copy(targetfile, sourcefile)
}



