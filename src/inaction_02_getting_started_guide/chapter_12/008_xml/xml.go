package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

func main() {

	xmlname := "/Users/yangjianwei/GolandProjects/src/inaction_02_getting_started_guide/chapter_12/008_xml/xmlfile.xml"

	xmlfile, err := os.Open(xmlname)

	if err != nil {
		log.Fatal(err)
	}

	defer xmlfile.Close()

	reader := bufio.NewReader(xmlfile)
	decoder := xml.NewDecoder(reader)

	for token, err := decoder.Token(); err == nil; token, err = decoder.Token() {
		switch token_type := token.(type) {
		case xml.StartElement:
			name := token_type.Name.Local
			fmt.Printf("token name = %v \n", name)

			for _, attr := range token_type.Attr {
				attrName := attr.Name.Local
				attrValue := attr.Value
				fmt.Printf("attrname = %v, attrvalue = %v \n", attrName, attrValue)
			}

		case xml.EndElement:
			fmt.Printf("end of token_type! \n")

		case xml.CharData:
			content := string([]byte(token_type))
			fmt.Printf("this is the content = %v \n", content)

		default:
			fmt.Printf("this is the default case! \n")
		}
	}

}

