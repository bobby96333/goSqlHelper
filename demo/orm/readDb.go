package main

import (
	"fmt"
	"github.com/bobby96333/goSqlHelper/demo/orm/entity"
	"github.com/bobby96333/goSqlHelper/demo/util"
)

func main(){


	helper:= util.OpenDb()
	var tb1 entity.Tb1
	err:=helper.Auto().From("tb_tb1").Where("id=?").QueryOrm(&tb1,4)
	if err!=nil {
		panic(err)
	}

	fmt.Printf("%+v",tb1)
	fmt.Println("done")

}




