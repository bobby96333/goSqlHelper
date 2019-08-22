package goSqlHelper

import (
	"context"
	"database/sql"
	"github.com/bobby96333/commonLib/stackError"
)

type SqlHelperRunner struct{
	SqlHelper
}

func(this *SqlHelperRunner) SetContext(ctx context.Context){
	this.context=ctx
}

func(this *SqlHelperRunner) BeginTx(ctx context.Context, opts *sql.TxOptions) stackError.StackError {
	tx,err:=this.Connection.BeginTx(ctx, opts)
	if err!=nil {
		return stackError.NewFromError(err,this.stckErrorPowerId)
	}
	this.tx=tx
	return nil
}

func(this *SqlHelperRunner) Begin() stackError.StackError {
	tx,err:=this.Connection.Begin()
	if err!=nil {
		return stackError.NewFromError( err,this.stckErrorPowerId)
	}
	this.tx=tx
	return nil
}