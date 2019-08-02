package goSqlHelper

import (
	"fmt"
	"testing"
	"time"
)

func TestMysqlOpen(t *testing.T) {

	con,err:=MysqlOpen("root:123456@tcp(centos:3306)/db1")
	if err!=nil {
		panic(err)
	}

	time1:=time.Now()
	for i:=0;i<1000;i++{
		testTable(con)
	}
	time2:=time.Now()
	for i:=0;i<1000;i++{
		testObj(con)
	}
	time3:=time.Now()

	fmt.Println("table time:",time2.Sub(time1))

	fmt.Println("obj time:",time3.Sub(time2))

	fmt.Println("down")
}

func testTable(con *SqlHelper){

	row,err:= con.QueryRow("select * from tb_tb1 where id=2")
	if err!=nil {
		panic(err)
	}
	fmt.Println(row["val"])
}
func testObj(con *SqlHelper){


}


type tb1 struct{
	id int
	val string
}


