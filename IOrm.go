package goSqlHelper

type IOrm interface{

	MapFields(columns []string) []interface{}
	PrimaryKeys()[]string
	TableName()string
	MapColumn() map[string]interface{}
}
