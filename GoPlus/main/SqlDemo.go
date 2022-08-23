package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error) {
	dsn := "root:xuan0623@tcp(127.0.0.1:3306)/contract?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func insertData() {
	sqlStr := "insert into test(username,password) values (?,?)"
	ret, err := db.Exec(sqlStr, "李四", "456")
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d\n", theID)
}

type User struct {
	username string
	password int
}

func queryOneRow() {
	s := "select * from test where username = ?"
	var u User
	err := db.QueryRow(s, "张三").Scan(&u.username, &u.password)
	if err != nil {
		fmt.Printf("query failed, the error is %v\n", err)
	} else {
		fmt.Println("查询成功！")
		fmt.Printf("姓名:%v\n密码:%v", u.username, u.password)
	}
}

func queryManyRow() {
	s := "select * from test"
	r, err := db.Query(s)
	var u User
	defer r.Close()
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		for r.Next() {
			r.Scan(&u.username, &u.password)
			fmt.Printf("姓名:%v 密码:%v\n", u.username, u.password)
		}
	}
}

func update() {
	s := "update test set username=?,password=? where id=?"
	r, err := db.Exec(s, "big kite", "456789", 2)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		i, _ := r.RowsAffected()
		fmt.Printf("i:%v\n", i)
	}
}

func delete() {
	s := "delete from test where id=?"
	r, err := db.Exec(s, 1)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		i, _ := r.RowsAffected()
		fmt.Printf("i:%v\n", i)
	}
}

func showAddress(an []int) {
	fmt.Printf("tempAddress is %p\n", &an)
	fmt.Printf("tempContent is %p\n", an)
}
func main() {
	dbLink := linkDatabase()
	result := Result{}
	GormOriginQuerySingle(dbLink, &result)
	/*
		ans := []int{1, 2, 3, 4, 5, 6}
		fmt.Printf("trueAddress is %p\n", &ans)
		fmt.Printf("trueContent is %p\n", ans)
		showAddress(ans)
	*/
	//del01()
	//del()
	// db, err := sql.Open("mysql", "root:xuan0623@/contract")
	// if err != nil {
	// 	fmt.Println("连接失败！")
	// 	panic(err)
	// }
	// print(db)
	// //最大连接时长
	// db.SetConnMaxIdleTime(time.Minute * 6)
	// //最大连接数
	// db.SetMaxOpenConns(10)
	// //空闲连接数
	// db.SetMaxIdleConns(10)
	// err := initDB()
	// if err != nil {
	// 	fmt.Printf("err:%v\n", err)
	// } else {
	// 	fmt.Println("连接成功！")
	// }
	// insertData()
	//queryOneRow()
	//queryManyRow()
	//update() //当更新后的数据和更新前的数据一样时不会进行更新操作
	//delete()
	//GormCreate()
	//GormInsert()
	//GormQuery()
	//GormQuery01()
	// arr := [5]int{1, 1, 1, 1, 1}
	// fmt.Printf("%p", &arr)
	// /*
	// 	聚合类型（数组和结构体）的新感悟
	// 	声明一个变量（比如arr和str）来表示这整个数组（或结构体），
	// 	这个变量的类型为数组或结构体，
	// 	当对这个变量进行取址操作时，
	// 	返回的均为该数组（或结构体）的第一个元素的地址（相当于对第一个元素取址）。
	// 	数组和结构体的区别在于：
	// 	数组中每个元素的类型都一样，通过下标来访问
	// 	结构体中元素的类型不相同（导致长度不同），通过属性名称来访问
	// */
	// str := struct {
	// 	add    int
	// 	name   string
	// 	time   int
	// 	height int
	// }{
	// 	111,
	// 	"mark",
	// 	123,
	// 	456,
	// }
	// fmt.Printf("%p\n", &str)
	// //fmt.Println(&str) //等价于fmt.Printf("%v\n", &str)(%v占位符表示默认类型，默认情况下会输出&{111 mark 123 456}即输出结构体的所有值)
	// // fmt.Printf("%p\n", &str.add)
	// // fmt.Printf("%p\n", &str.name)
	// // fmt.Printf("%p\n", &str.time)
	// // fmt.Printf("%p\n", &str.height)
	// LoadingTable(dbLink)
	/*
		result := []Result{}
		fmt.Println("result's capacity is", cap(result))
		fmt.Printf("result's address is %p\n", result)

		result01 := append(result, Result{1, "ks", "333"})
		fmt.Printf("result01's capacity is %v, result01's address is %p\n", cap(result01), result01)
		result01 = append(result01, Result{1, "ks", "333"})
		fmt.Printf("result01's capacity is %v, result01's address is %p\n", cap(result01), result01)
		result01 = append(result01, Result{1, "ks", "333"})
		fmt.Printf("result01's capacity is %v, result01's address is %p\n", cap(result01), result01)
		result01 = append(result01, Result{1, "ks", "333"})
		fmt.Printf("result01's capacity is %v, result01's address is %p\n", cap(result01), result01)
		result01 = append(result01, Result{1, "ks", "333"})
		fmt.Printf("result01's capacity is %v, result01's address is %p\n", cap(result01), result01)
		fmt.Printf("%p\n", result)
		fmt.Printf("%p\n", result01)
		fmt.Println(result01[0])
	*/
	//fmt.Println(result[0])
	//GormOriginQueryMany(dbLink, result)
	//Sresult := Result{}
	//GormOriginQuerySingle(dbLink, &Sresult)

}
