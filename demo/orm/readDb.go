package main

import (
	"fmt"
	"github.com/bobby96333/goSqlHelper/demo"
	"github.com/bobby96333/goSqlHelper/demo/orm/entity"
)

func main(){


	helper:=demo.OpenDb()
	var tb1 entity.Tb1
	err:=helper.Auto().From("tb_tb1").Where("id=?").QueryOrm(&tb1,2)
	if err!=nil {
		panic(err)
	}

	fmt.Printf("%+v",tb1)
	fmt.Println("done")

}




