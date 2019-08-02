package goSqlHelper

type IOrm interface{
	MapFields(columns []string) []interface{}
}
