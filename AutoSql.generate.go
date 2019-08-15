package goSqlHelper

import (
	"fmt"
	"strconv"
)

func (this *AutoSql) generateSelectSql () string{
	field:="*"
	if this.fieldSql != ""{
		field=this.fieldSql
	}
	sql:="SELECT "+field
	if this.tbname!= "" {
		sql+=" FROM "+this.tbname
	}
	for _,join :=range this.joins {
		sql+=" "+join
	}
	if this.where!=""{
		sql+=" WHERE "+this.where
	}
	if this.groupBy!=""{
		sql+=" GROUP BY "+this.groupBy
	}
	if this.having!=""{
		sql+=" HAVING "+this.having
	}
	if this.orderby!=""{
		sql+=" ORDER BY "+this.orderby
	}
	if this.limit!=0{
		sql+=" LIMIT "+ strconv.Itoa(this.limit)
	}
	return sql
}

func (this *AutoSql) generateUpdateSql () string{

	sql:=fmt.Sprintf("UPDATE %s set %s ",this.tbname,this.set)
	if this.where!=""{
		sql+=" WHERE "+this.where
	}
	if this.orderby!=""{
		sql+=" ORDER BY "+this.orderby
	}
	if this.limit!=0{
		sql+=" LIMIT "+ strconv.Itoa(this.limit)
	}
	return sql
}

func (this *AutoSql) generateDeleteSql () string{

	sql:=fmt.Sprintf("DELETE FROM %s ",this.tbname)
	if this.where!=""{
		sql+=" WHERE "+this.where
	}
	if this.orderby!=""{
		sql+=" ORDER BY "+this.orderby
	}
	if this.limit!=0{
		sql+=" LIMIT "+ strconv.Itoa(this.limit)
	}
	return sql
}
func (this *AutoSql) generateInsertSql () string{

	sql:=fmt.Sprintf("INSERT INTO %s ",this.tbname)
	if this.set!=""{
		sql+=" SET "+this.set
	}
	return sql
}
