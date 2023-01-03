package main

import "fmt"

func main() {

	fmt.Println("-------------------------- 1、练习题1  ----------------------------")
	s := make([]byte, 5)
	fmt.Printf("len(s) = %v, cap(s) = %v \n", len(s), cap(s))
	s = s[2:4]
	fmt.Printf("len(s) = %v, cap(s) = %v \n", len(s), cap(s))

	fmt.Println("-------------------------- 2、练习题2  ----------------------------")
	s1 := []byte{'p', 'o', 'e', 'm'}
	s2 := s1[2:]
	fmt.Printf("s2 = %v \n", string(s2)) // ['e', 'm']
	s2[1] = 't'
	fmt.Printf("s1 = %v \n", string(s1)) // ['p', 'o', 'e', 't']
	fmt.Printf("s2 = %v \n", string(s2)) // ['e', 't']

}
