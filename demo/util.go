package demo

import "github.com/bobby96333/goSqlHelper"

func OpenDb() *goSqlHelper.SqlHelper{

	con,err:= goSqlHelper.MysqlOpen("root:123456@tcp(centos:3306)/db1")
	if err!=nil {
		panic(err)
	}
	return con
}
