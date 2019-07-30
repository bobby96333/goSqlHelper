package goSqlHelper

type HelperTable struct{
	rows *[]HelperRow
	columns *[]string
}

func NewTable(rows []HelperRow,columns []string) *HelperTable {

	helper:=HelperTable{
		rows:&rows,
		columns:&columns,
	}
	return &helper

}


func(this HelperTable) Rows() *[]HelperRow {
	return this.rows
}
func(this HelperTable) Columns() *[]string {
	return this.columns
}