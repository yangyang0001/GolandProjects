package main

import "fmt"

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev := Reverse(input)
	doubleRev := Reverse(rev)
	fmt.Printf("original : %q \n", input)
	fmt.Printf("reversed : %q \n", rev)
	fmt.Printf("doublerev: %q \n", doubleRev)
}

func Reverse(s string) string {

	r := []byte(s)

	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		//fmt.Printf("r[%v] is %v, r[%v] is %v \n", i, r[i], i, r[j])
		r[i], r[j] = r[j], r[i]
		//fmt.Printf("r[%v] is %v, r[%v] is %v \n", i, r[i], i, r[j])
		//fmt.Println("--------------------------------------------------")
	}

	return string(r)
}
