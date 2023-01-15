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

	//fmt.Println("----------------------------------- 1、ReadLineFile ------------------------------------")
	//ReadLineFile()
	//
	//fmt.Println("----------------------------------- 2、ReadBuffFile ------------------------------------")
	//ReadBuffFile()
	//
	//fmt.Println("----------------------------------- 3、ReadcolumnFile ----------------------------------")
	//ReadcolumnFile()
	//
	//fmt.Println("----------------------------------- 4、ReadZipFile -------------------------------------")
	//ReadZipFile()
	//
	//fmt.Println("----------------------------------- 5、ReadGZipFile ------------------------------------")
	//ReadGZipFile()
	//
	//fmt.Println("----------------------------------- 6、WriteFile ------------------------------------")
	//WriteFile()
	//
	//fmt.Println("----------------------------------- 7、AppendFile   ------------------------------------")
	//AppendFile()
	//
	//fmt.Println("----------------------------------- 8、CopyFile     -------------------------------------")
	//CopyFile()

	CopyFileWithBytes()


}

func ReadLineFile()  {

	filename := "/Users/yangjianwei/GolandProjects/src/inaction_02_getting_started_guide/chapter_12/004_input_output_review/readlinefile.txt"
	openfile, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer openfile.Close()

	reader := bufio.NewReader(openfile)

	for {
		line, err := reader.ReadString('\n')

		fmt.Printf("ReadLineFile method invoke, line = %v", line)

		if err == io.EOF {
			log.Fatal(err)
		}

	}

}

func ReadBuffFile()  {

	filename := "/Users/yangjianwei/GolandProjects/src/inaction_02_getting_started_guide/chapter_12/004_input_output_review/readbufffile.txt"
	openfile, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer openfile.Close()

	reader := bufio.NewReader(openfile)

	var buff = make([]byte, 1024)

	for {
		n, err := reader.Read(buff)

		if err == io.EOF {
			log.Fatal(err)
		}

		fmt.Printf("ReadBuffFile method invoke, n = %v, string(buff) = %v", n, string(buff[0:n]))
	}

}

func ReadcolumnFile()  {

	filename := "/Users/yangjianwei/GolandProjects/src/inaction_02_getting_started_guide/chapter_12/004_input_output_review/readcolumnfile.txt"
	openfile, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer openfile.Close()

	reader := bufio.NewReader(openfile)

	var col0, col1, col2 []string

	for {
		line, err := reader.ReadString('\n')

		fields := strings.Fields(line)
		col0 = append(col0, fields[0])
		col1 = append(col1, fields[1])
		col2 = append(col2, fields[2])

		if err == io.EOF {
			break
		}

	}

	fmt.Printf("col0 = %v \n", col0)
	fmt.Printf("col1 = %v \n", col1)
	fmt.Printf("col2 = %v \n", col2)

}

func ReadZipFile()  {

	filename := "/Users/yangjianwei/GolandProjects/src/inaction_02_getting_started_guide/chapter_12/004_input_output_review/readzipfile.zip"
	zipRCloser, err := zip.OpenReader(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer zipRCloser.Close()

	for _, file := range zipRCloser.File {

		if strings.HasPrefix(file.Name, "__MACOSX") {
			continue
		} else {
			fmt.Printf("ReadZipFile method invoke, file.Name = %v \n", file.Name)

			openRCloser, err := file.Open()

			if err != nil {
				log.Fatal(err)
			}

			defer openRCloser.Close()

			reader := bufio.NewReader(openRCloser)

			for {
				line, err := reader.ReadString('\n')

				if err != nil {
					break
				}

				fmt.Printf("line = %v", line)
			}
		}

	}

}

func ReadGZipFile()  {

	filename := "/Users/yangjianwei/GolandProjects/src/inaction_02_getting_started_guide/chapter_12/004_input_output_review/readgzipfile.tar.gz"
	openfile, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer openfile.Close()

	gzipReader, err := gzip.NewReader(openfile)

	var reader * bufio.Reader

	if err != nil {
		reader = bufio.NewReader(openfile)
	} else {
		reader = bufio.NewReader(gzipReader)
	}

	for {
		line, err := reader.ReadString('\n')

		fmt.Printf("line = %v", line)

		if err == io.EOF {
			break
		}
	}

}

func WriteFile()  {
	filename := "/Users/yangjianwei/GolandProjects/src/inaction_02_getting_started_guide/chapter_12/004_input_output_review/writefile.txt"
	openfile, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)

	if err != nil {
		log.Fatal(err)
	}

	defer openfile.Close()

	writer := bufio.NewWriter(openfile)

	for i := 0; i < 10; i++ {
		result := "i = " + strconv.Itoa(i) + "\n"
		writer.WriteString(result)
	}

	writer.Flush()
}

func AppendFile()  {

	filename := "/Users/yangjianwei/GolandProjects/src/inaction_02_getting_started_guide/chapter_12/004_input_output_review/appendfile.txt"
	openfile, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)

	if err != nil {
		log.Fatal(err)
	}

	defer openfile.Close()

	writer := bufio.NewWriter(openfile)

	for i := 0; i < 10; i++ {
		writer.WriteString("i = " + strconv.Itoa(i) + "\n")
		writer.Flush()
	}


}


func CopyFile()  {
	sourcename := "/Users/yangjianwei/GolandProjects/src/inaction_02_getting_started_guide/chapter_12/004_input_output_review/sourcefile.txt"
	targetname := "/Users/yangjianwei/GolandProjects/src/inaction_02_getting_started_guide/chapter_12/004_input_output_review/targetfile.txt"

	sourcefile, err := os.Open(sourcename)
	if err != nil {
		log.Fatal(err)
	}
	defer sourcefile.Close()
	reader := bufio.NewReader(sourcefile)

	targetfile, err := os.OpenFile(targetname, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer targetfile.Close()
	writer := bufio.NewWriter(targetfile)

	io.Copy(writer, reader)

}

func CopyFileWithBytes()  {

	sourcename := "/Users/yangjianwei/GolandProjects/src/inaction_02_getting_started_guide/chapter_12/004_input_output_review/copyfilewith_source.txt"
	targetname := "/Users/yangjianwei/GolandProjects/src/inaction_02_getting_started_guide/chapter_12/004_input_output_review/copyfilewith_target.txt"

	sourcefile, err := os.Open(sourcename)
	if err != nil {
		log.Fatal(err)
	}
	defer sourcefile.Close()

	targetfile, err := os.OpenFile(targetname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer targetfile.Close()

	var buff = make([]byte, 1024)

	reader := bufio.NewReader(sourcefile)
	writer := bufio.NewWriter(targetfile)

	for {
		n, err := reader.Read(buff)

		fmt.Println(string(buff[0:n]))

		if err == io.EOF {
			log.Fatal(err)
		}

		writer.Write(buff[0:n])
		writer.Flush()
	}

}
