package main

import (
	"database/sql"
	"fmt"
	_ "log"
	_ "github.com/mattn/go-sqlite3"
)

type UserInfo struct {
	Uid int64 `json:"uid"`
	Username string `json:"username"`
	Department string `json:"department"`
	Created string `json:"created"`
}

func main() {

	db, err := sql.Open("sqlite3", "/Users/yangjianwei/test.db")
	checkErr(err)
	defer db.Close()

	//fmt.Println("------------------------------------------ Add method ------------------------------------------")
	//add0 := UserInfo{0, "xiaoming", "研发部", "2020-01-01"}
	//add1 := UserInfo{0, "xiaohong", "研发部", "2020-02-01"}
	//add2 := UserInfo{0, "zhangsan", "运营部", "2020-03-01"}
	//add3 := UserInfo{0, "zhangkai", "运维部", "2020-04-01"}
	//uid0, err := Add(db, add0)
	//uid1, _ := Add(db, add1)
	//uid2, _ := Add(db, add2)
	//uid3, _ := Add(db, add3)
	//fmt.Printf("uid0 = %v, uid1 = %v, uid2 = %v, uid3 = %v \n", uid0, uid1, uid2, uid3)

	fmt.Println("------------------------------------------ DeleteById method -----------------------------------")
	del_count, _ := DeleteById(db, 1)
	fmt.Printf("del_count = %v \n", del_count)


	fmt.Println("------------------------------------------ Update method ---------------------------------------")
	update := UserInfo{2, "xiaoling", "研发部", "2021-10-01"}
	upd_count, _ := Update(db, update)
	fmt.Printf("upd_count = %v \n", upd_count)


	fmt.Println("------------------------------------------ SelectById method -----------------------------------")
	select_userinfo, _ := SelectById(db, 2)
	fmt.Printf("select_userinfo = %v \n", select_userinfo)

	fmt.Println("------------------------------------------ SelectList method -----------------------------------")
	username := "zhang"
	list, _ := SelectList(db, username)
	for index, info := range list {
		fmt.Printf("index = %v, info = %v \n", index, info)
	}

}


/**
 * 添加单个
 */
func Add(db *sql.DB, userinfo UserInfo) (int64, error) {

	result, err := db.Exec("insert into userinfo(username, department, created) values(?, ?, ?)", userinfo.Username, userinfo.Department, userinfo.Created)
	if err != nil {
		return -1, fmt.Errorf("Add method invoke exec error, %q: %v", userinfo, err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("Add method invoke last insert id error, %q: %v", userinfo, err)
	}

	return id, nil
}


/**
 * 删除单个
 */
func DeleteById(db *sql.DB, uid int64) (int64, error) {

	result, err := db.Exec("delete from userinfo where uid = ?", uid)
	if err != nil {
		return -1, fmt.Errorf("DeleteById method invoke exec error, %q: %v", uid, err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return -1, fmt.Errorf("DeleteById method invoke rows affected error, %q: %v", uid, err)
	}

	return affected, nil
}


/**
 * 批量删除
 */
func DeleteBatch(db *sql.DB, uids []int64) (int64, error) {

	sql := "delete from userinfo where uid in ("
	for index, uid := range uids {
		if index != len(uids)-1 {
			sql += fmt.Sprintf("%v, ", uid)
		} else {
			sql += fmt.Sprintf("%v)", uid)
		}
	}

	result, err := db.Exec(sql)
	if err != nil {
		return -1, fmt.Errorf("DeleteBatch method invoke exec error, %q: %v", uids, err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return -1, fmt.Errorf("DeleteBatch method invoke rows affected error, %q: %v", uids, err)
	}

	return affected, nil
}


/**
 * 修改数据
 */
func Update(db *sql.DB, userinfo UserInfo) (int64, error) {

	result, err := db.Exec("update userinfo set username = ?, department = ?, created = ? where uid = ?",
		userinfo.Username, userinfo.Department, userinfo.Created, userinfo.Uid)
	if err != nil {
		return -1, fmt.Errorf("Update method invoke exec error, %q: %v", userinfo, err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return -1, fmt.Errorf("Update method invoke rows affected error, %q: %v", userinfo, err)
	}

	return affected, nil
}


/*
 * 查询单行
 */
func SelectById(db *sql.DB, uid int64) (UserInfo, error) {

	var userinfo UserInfo
	row := db.QueryRow("select * from userinfo where uid = ?", uid)

	if err := row.Err(); err != nil {
		return userinfo, fmt.Errorf("SelectById method invoke query error, %q: %v", uid, err)
	}

	if err := row.Scan(&userinfo.Uid, &userinfo.Username, &userinfo.Department, &userinfo.Created); err != nil {
		return userinfo, fmt.Errorf("SelectById method invoke scan error, %q: %v", uid, err)
	}

	return userinfo, nil
}


/*
 * 查询数组
 */
func SelectList(db *sql.DB, username string) ([]UserInfo, error) {

	var userinfos []UserInfo

	rows, err := db.Query("select * from userinfo where username like " + "\"" + username + "%\"")
	if err != nil {
		return nil, fmt.Errorf("SelectList method invoke query error, %q: %v", username, err)
	}

	defer rows.Close()

	for rows.Next() {
		var userinfo UserInfo
		if err := rows.Scan(&userinfo.Uid, &userinfo.Username, &userinfo.Department, &userinfo.Created); err != nil {
			return nil, fmt.Errorf("SelectList method invoke scan error, %q: %v", username, err)
		}
		userinfos = append(userinfos, userinfo)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("SelectList method invoke rows error, %q: %v", username, err)
	}

	return userinfos, nil
}


func checkErr(err error)  {
	if err != nil {
		fmt.Printf("error = %v \n", err)
		panic(err)
	}
}
