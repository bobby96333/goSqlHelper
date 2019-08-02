package entity

import (
	"github.com/bobby96333/goSqlHelper"
	"strings"
)

type Tb1 struct{
	goSqlHelper.IOrm
	Id int
	Val string
}

func(this *Tb1) MapFields(columns []string) []interface{}{

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