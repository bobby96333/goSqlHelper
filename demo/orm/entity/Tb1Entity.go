package entity
import(
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"strings"
)
type Tb1Entity struct {
	//remark name length:0
	Name        string
	Number      sql.NullInt64
	//good' 88
	Col1        sql.NullString
	XdFf        sql.NullInt64
	//insert time
	CreateTime  mysql.NullTime
	Id          uint32
	Val         sql.NullString
}

func(this *Tb1Entity) MapFields(columns []string) []interface{}{

	ret:=make([]interface{},len(columns))
	for i,col :=range columns{
		realCol:=col
		if index:=strings.Index(col,".");index != -1 {
			realCol=col[index+1:]
		}
		switch(realCol){
		case "xd ff":ret[i]       = &this.XdFf
		case "create_time":ret[i]  = &this.CreateTime
		case "id":ret[i]          = &this.Id
		case "val":ret[i]         = &this.Val
		case "col_name":ret[i]    = &this.Name
		case "number":ret[i]      = &this.Number
		case "col1":ret[i]        = &this.Col1

		}
	}
	return ret
}
func(this *Tb1Entity) MapColumn() map[string]interface{}{
	return map[string]interface{}{
		"xd ff"       : &this.XdFf,
		"create_time"  : &this.CreateTime,
		"id"          : &this.Id,
		"val"         : &this.Val,
		"col_name"    : &this.Name,
		"number"      : &this.Number,
		"col1"        : &this.Col1,

	}
}
func(this *Tb1Entity) PrimaryKeys() []string{
	return []string{"id",}
}
func(this *Tb1Entity) TableName()string{

	return "tb_tb1"
}

