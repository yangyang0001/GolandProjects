package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Mine struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

type Pattern struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func main() {

	fmt.Println("*************************************** 1、为 struct 类型添加方法, 类似 java 中的 get 方法 ****************")
	mine0 := Mine{"zhangsan", "zhang123456"}
	mine1 := Mine{"wangwu", "wang123456"}

	fmt.Printf("mine0 json is : %v \n", jsonParse(mine0))
	fmt.Printf("mine1 json is : %v \n", jsonParse(mine1))

	fmt.Printf("mine0.Name =  %v, mine0.Pass = %v \n", mine0.GetName(), mine0.GetPass())
	fmt.Printf("mine1.Name =  %v, mine1.Pass = %v \n", mine1.GetName(), mine1.GetPass())

	fmt.Println("*************************************** 2、使用 函数 ***************************************************")
	fmt.Printf("mine0.Name =  %v, mine0.Pass = %v \n", GetName(mine0), GetPass(mine0))
	fmt.Printf("mine1.Name =  %v, mine1.Pass = %v \n", GetName(mine1), GetPass(mine1))

	fmt.Println("*************************************** 3、指针接收者方法 类似 java 中的 set 方法 *************************")
	mine0.SetName("maosan")
	mine0.SetPass("mao123")
	mine1.SetName("gousi")
	mine1.SetPass("gou12")

	fmt.Printf("mine0.Name =  %v, mine0.Pass = %v \n", mine0.GetName(), mine0.GetPass())
	fmt.Printf("mine1.Name =  %v, mine1.Pass = %v \n", mine1.GetName(), mine1.GetPass())

	fmt.Println("*************************************** 4、指针接收者函数 ***********************************************")
	SetName(&mine0, "xiaoming")
	SetPass(&mine0, "ming1234")
	fmt.Printf("mine0.Name =  %v, mine0.Pass = %v \n", mine0.GetName(), mine0.GetPass())

	fmt.Println("**************************************** 5、使用 struct Pattern, 模拟 java 中的 set 和 get method *******")
	pattern := Pattern{1, "Factory", "工厂模式"}
	fmt.Printf("pattern json is : %v \n", jsonParse(pattern))

	pattern.SetId(2)
	pattern.SetName("Chain")
	pattern.SetDesc("责任链模式")
	fmt.Printf("pattern.Id = %v, \"pattern.Name = %v, \"pattern.Desc = %v \n", pattern.GetId(), pattern.GetName(), pattern.GetDesc())

}

/*************************************** 1、为 struct 类型添加方法, 类似 java 中的 get 方法 ***************/
func (m Mine) GetName() string {
	return m.Name
}

func (m Mine) GetPass() string {
	return m.Pass
}

/*************************************** 2、使用 函数 **************************************************/
func GetName(m Mine) string {
	return m.Name
}

func GetPass(m Mine) string {
	return m.Pass
}

/*************************************** 3、指针接收者方法 类似 java 的 set 方法 **************************/
func (m *Mine) SetName(name string) {
	m.Name = name
}

func (m *Mine) SetPass(pass string) {
	m.Pass = pass
}

/*************************************** 4、指针接收者函数 ***********************************************/
func SetName(m *Mine, name string) {
	m.Name = name
}

func SetPass(m *Mine, pass string) {
	m.Pass = pass
}

/*************************************** 5、使用 struct Pattern, 模拟 java 中的 set 和 get method *******/
func (pattern *Pattern) GetId() int {
	return pattern.Id
}

func (pattern *Pattern) GetName() string {
	return pattern.Name
}

func (pattern *Pattern) GetDesc() string {
	return pattern.Desc
}

func (pattern *Pattern) SetId(id int) {
	pattern.Id = id
}

func (pattern *Pattern) SetName(name string) {
	pattern.Name = name
}

func (pattern *Pattern) SetDesc(desc string) {
	pattern.Desc = desc
}

func jsonParse(obj any) string {

	bytes, err := json.Marshal(obj)

	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}
