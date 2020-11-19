package db_conn

import (
	"context"
	"fmt"
)

type DB0USReg interface {
	GetUser(name string) map[string]string
	SetData(id string,data interface{})
	DelUser(id string)
}

type Db0usReg struct {
	ConfigDb map[string]interface{}
	DbConn
}

func (db0usR *Db0usReg) DelUser(id string) {
	conn := db0usR.ConnectOpen
	var ctx = context.Background()
	err := conn.Del(ctx,id).Err()
	if err != nil{
		fmt.Println(err)
	}
}

func (db0usR *Db0usReg) SetData(id string,data interface{})  {
	conn := db0usR.ConnectOpen
	var ctx = context.Background()
	err := conn.HMSet(ctx,id,data).Err()
	if err != nil{
		fmt.Println(err)
	}
}

func (db0usR *Db0usReg) GetUser(name string) map[string]string {
	conn := db0usR.ConnectOpen
	var ctx = context.Background()
	v := conn.HGetAll(ctx,name).Val()
	return v

}

func (db0usR *Db0usReg)Connect() {
	db0usR.DbConn.DbHost = db0usR.ConfigDb["DbHost"].(string)
	db0usR.DbConn.DbPort = db0usR.ConfigDb["DbPort"].(string)
	db0usR.DbConn.DbNum  = 0
	db0usR.DbConn.DbPassword = db0usR.ConfigDb["DbPassword"].(string)
	db0usR.DbConn.PoolSiz = db0usR.ConfigDb["PoolSiz"].(int)
	db0usR.DbConn.Connect()
}