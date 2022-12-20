package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func FindDB() *sql.DB {

	config := mysql.Config{
		User:                 "root",
		Passwd:               "Yang199001",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "recordings",
		AllowNativePasswords: true,
	}

	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("mysql is connected!")

	return db
}

/**
 * TODO: 首字母大写, 以便于在序列化时能够被访问!
 */
type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

/*
 * 查询数组
 */
func selectList(db *sql.DB, artist string) ([]Album, error) {

	var albums []Album

	rows, err := db.Query("select * from album where artist = ?", artist)

	if err != nil {
		return nil, fmt.Errorf("selectList method invoke query error, %q: %v", artist, err)
	}

	defer rows.Close()

	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("selectList method invoke scan error, %q: %v", artist, err)
		}
		albums = append(albums, alb)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("selectList method invoke rows error, %q: %v", artist, err)
	}

	return albums, nil
}

/*
 * 查询单行
 */
func selectById(db *sql.DB, id int64) (Album, error) {

	var album Album
	row := db.QueryRow("select * from album where id = ?", id)

	if err := row.Err(); err != nil {
		return album, fmt.Errorf("selectById method invoke query error, %q: %v", id, err)
	}

	if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
		return album, fmt.Errorf("selectById method invoke scan error, %q: %v", id, err)
	}

	return album, nil
}

/**
 * 添加单个
 */
func add(db *sql.DB, album Album) (int64, error) {

	result, err := db.Exec("insert into album(title, artist, price) values (?, ?, ?)", album.Title, album.Artist, album.Price)

	if err != nil {
		return -1, fmt.Errorf("add method invoke exec error, %q: %v", album, err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		return -1, fmt.Errorf("add method invoke last insert id error, %q: %v", album, err)
	}

	return id, nil
}

/**
 * 删除单个
 */
func deleteById(db *sql.DB, id int64) (int64, error) {

	result, err := db.Exec("delete from album where id = ?", id)
	if err != nil {
		return -1, fmt.Errorf("deleteById method invoke exec error, %q: %v", id, err)
	}

	affected, err := result.RowsAffected()

	if err != nil {
		return -1, fmt.Errorf("deleteById method invoke rows affected error, %q: %v", id, err)
	}

	return affected, nil
}

/**
 * 批量删除
 */
func deleteBatch(db *sql.DB, ids []int64) (int64, error) {

	sql := "delete from album where id in ("
	for index, id := range ids {
		if index != len(ids)-1 {
			sql += fmt.Sprintf("%v, ", id)
		} else {
			sql += fmt.Sprintf("%v)", id)
		}
	}

	result, err := db.Exec(sql)
	if err != nil {
		return -1, fmt.Errorf("deleteBatch method invoke exec error, %q: %v", ids, err)
	}

	affected, err := result.RowsAffected()

	if err != nil {
		return -1, fmt.Errorf("deleteBatch method invoke rows affected error, %q: %v", ids, err)
	}

	return affected, nil

}

/**
 * 修改数据
 */
func update(db *sql.DB, album Album) (int64, error) {

	result, err := db.Exec("update album set title = ?, artist = ?, price = ? where id = ?",
		album.Title, album.Artist, album.Price, album.ID)
	if err != nil {
		return -1, fmt.Errorf("update method invoke exec error, %q: %v", album, err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return -1, fmt.Errorf("update method invoke rows affected error, %q: %v", album, err)
	}

	return affected, nil

}
