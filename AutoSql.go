package goSqlHelper

import "strconv"

func NewAutoSql(helper *SqlHelper) *AutoSql{
	var orm=new(AutoSql)
	orm.joins=make([]string,0)
	orm.sqlHelper=helper
	return orm
}

type AutoSql struct{
	sqlHelper *SqlHelper
	fieldSql string
	tbname string
	where string
	groupBy string
	orderby string
	having string
	limit int
	joins []string
}
func (this *AutoSql)Select(fieldSql string) *AutoSql{
	this.fieldSql=fieldSql
	return this
}
func (this *AutoSql)From(tbname string) *AutoSql{
	this.tbname=tbname
	return this
}
func (this *AutoSql)Where(where string) *AutoSql{
	this.where=where
	return this
}
func (this *AutoSql)Join(joinSql string) *AutoSql{
	this.joins = append(this.joins,joinSql)
	return this
}
func (this *AutoSql)Groupby(groupBySql string) *AutoSql{
	this.groupBy = groupBySql
	return this
}
func (this *AutoSql)Orderby(OrderbySql string) *AutoSql{
	this.orderby = OrderbySql
	return this
}
func (this *AutoSql)Having(having string) *AutoSql{
	this.having = having
	return this
}
func (this *AutoSql)Limit(limit int) *AutoSql{
	this.limit = limit
	return this
}
func (this *AutoSql) GenerateSql()string{
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

func (this *AutoSql)  QueryRows(args ...interface{})([]HelperRow, error) {
	sql:=this.GenerateSql()
	return this.sqlHelper.QueryRows(sql,args...)
}

func (this *AutoSql) 	QueryTable( args ...interface{})(*HelperTable, error) {
	sql:=this.GenerateSql()
	return this.sqlHelper.QueryTable(sql,args...)
}

func (this *AutoSql) Querying(args ...interface{})(*Querying,error){
	sql:=this.GenerateSql()
	return this.sqlHelper.Querying(sql,args...)
}

func (this *AutoSql) QueryRow( args ...interface{})(HelperRow, error) {
	sql:=this.GenerateSql()
	return this.sqlHelper.QueryRow(sql,args...)
}

func (this *AutoSql) QueryScalarInt(args ...interface{})(int, error) {
	sql:=this.GenerateSql()
	return this.sqlHelper.QueryScalarInt(sql,args...)
}

func (this *AutoSql) QueryOrm(orm IOrm, args ...interface{})(error) {
	sql:=this.GenerateSql()
	return this.sqlHelper.QueryOrm(orm,sql,args...)
}
