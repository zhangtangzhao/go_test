package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	_ "github.com/go-sql-driver/mysql"
)

func connect() (*sql.DB,error){

	db, _ := sql.Open("mysql", "root:123456@tcp(192.168.48.131:3306)/test")
	//设置数据库最大连接数
	db.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	db.SetMaxIdleConns(10)
	//验证连接
	if err := db.Ping(); err != nil {
		errors.Wrap(err,"db connect fail")
		return nil,err
	}
	fmt.Println("connnect success")

	return db,nil

}

func Query(db *sql.DB) (string, error){
	var name string
	err := db.QueryRow("select name from test_info where id = ?",2).Scan(&name)
	if(errors.Is(err,sql.ErrNoRows)){
		// 因为这个sql.ErrNoRows是查不到数据而触发，所以不用向上Wrap,直接在dao处理错误返回空数据
		fmt.Printf("query ErrNoRows id[%v] error[%v]\n",2,err)
		return "", nil
	}else if err != nil{
		errors.Wrap(err,"query sql fail")
	}
	return name,err
}

func main(){
	db, err := connect()
	if err != nil{
		fmt.Printf("sql connect error: %v \n",err)
		return
	}
	name,err := Query(db)
	if err != nil{
		fmt.Printf("sql query error : %v \n",err)
		return
	}
	fmt.Printf("data name[%v]",name)
	db.Close()
}
