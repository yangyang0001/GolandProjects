package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

type MineTask struct {

	Name string `json:"name"`
	Pass string `json:"pass"`

}

func main() {

	//channel := make(chan MineTask)
	//
	//go SendTask(channel)
	//
	//go DoTask(channel)
	//
	//time.Sleep(time.Second * 2)


	channel := make(chan *MineTask)

	go MakeTask(channel)

	go DoneTask(channel)

	time.Sleep(time.Second * 2)

}

func SendTask(channel chan MineTask)  {
	for i := 0; i < 10; i++ {
		channel <- MineTask{"name-" + strconv.Itoa(i), "pass-" + strconv.Itoa(i)}
	}
}

func DoTask(channel chan MineTask)  {
	for task := range channel {
		fmt.Printf("task = %v \n", jsonParse(task))
	}
}

func MakeTask(channel chan *MineTask) {
	for i := 0; i < 10; i++ {
		channel <- &MineTask{"name_" + strconv.Itoa(i), "pass_" + strconv.Itoa(i)}
	}
}

func DoneTask(channel chan *MineTask)  {
	for task := range channel {
		fmt.Printf("task = %v \n", jsonParse(task))
	}
}

func jsonParse(obj interface{}) string {

	marshal, err := json.Marshal(obj)

	if err != nil {
		log.Fatal(err)
	}

	return string(marshal)

}
