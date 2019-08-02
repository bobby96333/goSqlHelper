package goSqlHelper

import (
	"context"
	"database/sql"
)

type SqlHelperRunner struct{
	SqlHelper
}

func(this *SqlHelperRunner) SetContext(ctx context.Context){
	this.context=ctx
}

func(this *SqlHelperRunner) BeginTx(ctx context.Context, opts *sql.TxOptions) error {
	tx,err:=this.Connection.BeginTx(ctx, opts)
	if err!=nil {
		return err
	}
	this.tx=tx
	return nil
}

func(this *SqlHelperRunner) Begin() error {
	tx,err:=this.Connection.Begin()
	if err!=nil {
		return err
	}
	this.tx=tx
	return nil
}