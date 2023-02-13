
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=Yang199001 dbname=test sslmode=disable")
	checkErr(err)
	defer db.Close()

	//插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo(username, department, created) VALUES($1,$2,$3) RETURNING uid")
	checkErr(err)
	res, err := stmt.Exec("xiaohong", "研发部门", "2012-12-09")
	checkErr(err)

	// pq 不支持这个函数, 因为他没有类似MySQL的自增ID
	// id, err := res.LastInsertId()
	// checkErr(err)
	// fmt.Println(id)

	var lastInsertId int
	err = db.QueryRow("INSERT INTO userinfo(username, department, created) VALUES($1, $2, $3) returning uid;", "xiaoming", "研发部门", "2012-12-09").Scan(&lastInsertId)
	checkErr(err)
	fmt.Printf("lastInsertId = %v \n", lastInsertId)


	//更新数据
	stmt, err = db.Prepare("update userinfo set username = $1 where uid = $2")
	checkErr(err)
	res, err = stmt.Exec("rootupdate", 1)
	checkErr(err)
	upd_count, err := res.RowsAffected()
	checkErr(err)
	fmt.Printf("upd_count = %v \n", upd_count)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Printf("uid = %v, username = %v, department = %v, created = %v", uid, username, department, created)
	}

	//删除数据
	stmt, err = db.Prepare("delete from userinfo where uid = $1")
	checkErr(err)
	res, err = stmt.Exec(1)
	checkErr(err)
	del_count, err = res.RowsAffected()
	checkErr(err)

	fmt.Printf("del_count = %v \n", del_count)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}