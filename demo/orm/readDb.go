package main

import (
	"fmt"
	"github.com/bobby96333/commonLib/stackError"
	"github.com/bobby96333/goSqlHelper"
	"github.com/bobby96333/goSqlHelper/demo/orm/entity"
	"github.com/bobby96333/goSqlHelper/demo/util"
)

func main(){

	goSqlHelper.DefaultDebugModel=true
	helper:= util.OpenDb()
	var tb1 entity.Tb1Entity
	err:=helper.Auto().From("tb_tb1").Where("id=?").QueryOrm(&tb1,5)
	stackError.CheckExitError(err)

	fmt.Printf("%+v",tb1)
	fmt.Println("done")

}




