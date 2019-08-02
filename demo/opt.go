package main

import (
	"fmt"
	"github.com/bobby96333/goSqlHelper"
	"github.com/bobby96333/goSqlHelper/demo/util"
)

func main(){

	helper:= util.OpenDb()
	helper.OpenDebug()
	cnt,err:= helper.Exec("update tb_tb1 set val=concat(val,'W')")
	if err!=nil {
		panic(err)
	}
	fmt.Println("updates ",cnt," record")

	record:=goSqlHelper.HelperRow{
		"val":"8899",
	}
	id,err:=helper.Auto().Insert("tb_tb1").SetRow(&record).ExecInsert()
	if err!=nil {
		panic(err)
	}
	fmt.Printf("insert a record id:%d\n",id)



	updateCnt,err := helper.Auto().Update("tb_tb1").SetRow(&record).Where("id=?").ExecUpdateOrDel(5)
	if err!=nil {
		panic(err)
	}
	fmt.Printf("update record:%d\n",updateCnt)


	updateCnt,err = helper.Auto().Delete("tb_tb1").Where("id=?").ExecUpdateOrDel(id)
	if err!=nil {
		panic(err)
	}
	fmt.Printf("delete record:%d\n",updateCnt)



}