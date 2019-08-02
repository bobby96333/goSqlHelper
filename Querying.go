package goSqlHelper

import (
	"database/sql"
	"github.com/bobby96333/goSqlHelper/HelperError"
)

type Querying struct {
	rows *sql.Rows
	_cols []string
}


func (this *Querying) Close(){
	this.rows.Close()
	this._cols=nil
}

func (this *Querying) Columns() ([]string,HelperError.Error){
	if this._cols==nil {
		var err error
		this._cols,err = this.rows.Columns()
		if err!=nil {
			return nil,HelperError.NewParent(err)
		}
	}
	return this._cols,nil
}

func (this Querying) QueryRow() (HelperRow,HelperError.Error){

	cols,err:=this.Columns()
	if err!=nil {
		return nil,err
	}
	scanArgs := make([]interface{}, len(cols))
	values := make([]interface{}, len(cols))
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
				record[cols[i]] = string(col.([]byte))
			default:
				record[cols[i]] = col

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