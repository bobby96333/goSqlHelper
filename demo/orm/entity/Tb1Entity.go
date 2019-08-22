package entity

import (
	"github.com/bobby96333/commonLib/sqlTypes"
	"strings"
)

type Tb1Entity struct{
	Id int
	Val sqlTypes.NullString
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

	return map[string]interface{}{
		"id":&this.Id,
		"val":&this.Val,
	}

}
func(this *Tb1Entity) PrimaryKeys() []string{
	return []string{"id"}
}
func(this *Tb1Entity) TableName()string{

	return "tb_tb1"
}