package entity

import (
	"database/sql"
	"github.com/bobby96333/goSqlHelper"
	"strings"
)

type Tb1Entity struct{
	goSqlHelper.IEntity
	Id int
	Val sql.NullString
}

func(this *Tb1Entity) MapFields(columns []string) []interface{}{

	ret:=make([]interface{},len(columns))
	for i,col :=range columns{
		realCol:=col
		if index:=strings.Index(col,".");index != -1 {
			realCol=col[index+1:]
		}
		switch(realCol){
		case "id":ret[i]=&this.Id
		case "val":ret[i]=&this.Val
		default: panic("reflect it")
		}
	}
	return ret
}
func(this *Tb1Entity) MapColumn() map[string]interface{}{

	ret:=make(map[string]interface{})
	ret["id"]=this.Id
	ret["val"]=this.Val
	return ret
}
func(this *Tb1Entity) PrimaryKeys() []string{
	return []string{"id"}
}
func(this *Tb1Entity) TableName()string{

	return "tb_tb1"
}