package main

import (
	"fmt"
	"log"
)

func main() {

	db := FindDB()

	// 查询数组
	albums, err := selectList(db, "John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("albums is : %v \n", albums)

	// 查询单个
	album, err := selectById(db, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("album is  : %v \n", album)

	// 插入
	addalbum := Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	}
	id, err := add(db, addalbum)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("last insert id is : %v \n", id)

	// 修改
	updateAlbum := Album{
		ID:     4,
		Title:  "zhangsan",
		Artist: "zhangsan",
		Price:  100.99,
	}
	updated, err := update(db, updateAlbum)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("updated is : %v \n", updated)

	// 删除单个
	affected, err := deleteById(db, 4)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("affected is : %v \n", affected)

	// 批量删除
	ids := []int64{1, 2, 3}
	affecteds, err := deleteBatch(db, ids)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("affecteds is : %v \n", affecteds)

}
