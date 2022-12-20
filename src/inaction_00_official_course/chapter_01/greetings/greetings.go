package greetings

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func Hello(name string) string {

	message := fmt.Sprintf("Hello method, Hello %v. Welcome!", name)

	return message
}

func HelloERR(name string) (string, error) {

	if name == "" {
		err := errors.New("name is empty")
		return "", err
	} else {
		message := fmt.Sprintf("HelloERR method, Hello %v. ", name)
		return message, nil
	}
}

func Hellos(names []string) (map[string]string, error) {

	resultMap := make(map[string]string)

	for _, name := range names {
		message, err := HelloERR(name)
		if err != nil {
			log.Fatal(err)
			return nil, err
		} else {
			resultMap[name] = message
		}
	}

	return resultMap, nil
}

func RandomHello(names []string) (map[string]string, error) {

	resultMap := make(map[string]string)

	for index, name := range names {

		if name == "" {
			errorinfo := fmt.Sprintf("names index is %v", index, ", the value is empty!")
			err := errors.New(errorinfo)
			return nil, err
		} else {
			message := fmt.Sprintf(formatHello(), name)
			resultMap[name] = message
		}
	}

	return resultMap, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func formatHello() string {
	formats := []string{
		"Hi %v. Welcome!",
		"Greate to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
