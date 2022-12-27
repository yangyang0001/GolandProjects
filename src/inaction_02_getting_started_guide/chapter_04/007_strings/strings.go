package main

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("--------------------------------- 1、strings 前缀和后缀 ------------------------------")
	string0 := "this is go language strings example!"
	fmt.Printf("string0.HasPrefix('this') = %v \n", strings.HasPrefix(string0, "this"))
	fmt.Printf("string0.HasSuffix('zge!') = %v \n", strings.HasSuffix(string0, "zge!"))

	// 字符串包含关系!
	fmt.Println("--------------------------------- 2、字符串包含关系 ----------------------------------")
	fmt.Printf("strings.Contains(string0, 'is') = %v \n", strings.Contains(string0, "s"))

	fmt.Println("-------------------- 3、判断子字符串或字符在父字符串中出现的位置（索引）--------------------")
	index := strings.Index(string0, "s")
	lastIndex := strings.LastIndex(string0, "s")
	fmt.Printf("index = %v, lastIndex = %v \n", index, lastIndex)

	fmt.Println("--------------------------------- 4、 字符串替换 -------------------------------------")
	replace := strings.Replace(string0, "s", "S", -1)
	fmt.Printf("replace string = %v \n", replace)

	fmt.Println("--------------------------------- 5、 统计字符串出现次数 ------------------------------")
	count := strings.Count(string0, "s")
	fmt.Printf("the count = %v \n", count)

	fmt.Println("--------------------------------- 6、 重复字符串 -------------------------------------")
	repeat := strings.Repeat("M", 10)
	fmt.Printf("the repeat = %v \n", repeat)

	fmt.Println("--------------------------------- 7、 修改字符串大小写 --------------------------------")
	lower := strings.ToLower(string0)
	upper := strings.ToUpper(string0)
	fmt.Printf("the lower = %v, the upper = %v \n", lower, upper)

	fmt.Println("--------------------------------- 8、 去除空字符串 -----------------------------------")
	string1 := " zhangsan "
	trimspace := strings.TrimSpace(string1)
	fmt.Printf("string1 = %v, trimspace = %v \n", string1, trimspace)
	trimleft := strings.TrimLeft(string1, " z")
	fmt.Printf("string1 = %v, trimleft = %v \n", string1, trimleft)

	fmt.Println("--------------------------------- 9、 分割字符串 -------------------------------------")
	string2 := "welcome to the go language!"
	split := strings.Split(string2, " ")
	fmt.Printf("string2 = %v, split = %v \n", string2, split)
	for index, str := range split {
		fmt.Printf("index = %v, str = %v \n", index, str)
	}

	fmt.Println("--------------------------------- 10、 拼接 join 的使用 ------------------------------")
	array := []string{"abc", "def", "ghi"}
	join := strings.Join(array, "-")
	fmt.Printf("the join string = %v \n", join)

	fmt.Println("--------------------------------- 11、 从字符串中读取内容 -----------------------------")
	reader := strings.NewReader(string2)
	var buffer = make([]byte, 8)
	for {
		count, err := reader.Read(buffer)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("readstring = %v \n", string(buffer[0:count]))
	}

	fmt.Println("--------------------------------- 12、 字符串与其它类型的转换 --------------------------")
	// 这时就需要 使用 strconv
	result, _ := strconv.Atoi("100")
	newstr := strconv.Itoa(88)
	fmt.Printf("result = %v, newstr = %v \n", result, newstr)

}
