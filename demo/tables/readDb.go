package main

import (
	"fmt"
	"github.com/bobby96333/goSqlHelper/demo"
)

func main(){
	helper:=demo.OpenDb()
	row,err:=helper.QueryRow("select * from tb_tb1 limit 1")
	if err!=nil {
		panic(err)
	}
	fmt.Printf("%+v\n",row)

	//equal to
	row,err=helper.Auto().Select("*").From("tb_tb1").Limit(1).QueryRow()
	if err!=nil {
		panic(err)
	}
	fmt.Printf("%+v\n",row)
	fmt.Println("done\n")

}
