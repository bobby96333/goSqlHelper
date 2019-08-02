package goSqlHelper

import (
	"database/sql"
	"github.com/bobby96333/goSqlHelper/HelperError"
)

type Querying struct {
	rows *sql.Rows
	cols []string
}


func (this *Querying) Close(){
	this.rows.Close()
	this.cols=nil
}

func (this *Querying) Columns() ([]string,HelperError.Error){
	if this.cols==nil {
		var err error
		this.cols,err = this.rows.Columns()
		if err!=nil {
			return nil,HelperError.NewParent(err)
		}
	}
	return this.cols,nil
}

func (this Querying) QueryObject(object interface{}) (HelperError.Error){
	err := this.rows.Scan(object)
	if err!=nil {
		return HelperError.NewParent(err)
	}
	return nil
}

func (this Querying) QueryRow() (HelperRow,HelperError.Error){

	scanArgs := make([]interface{}, len(this.cols))
	values := make([]interface{}, len(this.cols))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	if this.rows.Next() {

		err := this.rows.Scan(scanArgs...)
		if err!=nil {
			return nil,HelperError.NewParent(err)
		}
		record := make(HelperRow)
		for i, col := range values {
			if col == nil {
				continue
			}
			switch col.(type) {

			case []byte:
				record[this.cols[i]] = string(col.([]byte))
			default:
				record[this.cols[i]] = col

			}
		}

		return record,nil
	}

	return nil,HelperError.New(HelperError.ErrorEmpty)

}


func NewQuerying(rows *sql.Rows) (*Querying){
	querying:=new(Querying)
	querying.rows=rows
	return querying
}