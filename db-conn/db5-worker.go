package db_conn

import (
	"context"
	"strconv"
	"telegram-aggregatio-data/helper"
)


type DB5 interface {
	GetStatusIter()int
	SetStatusIter(iter int)
	UpdateServerStatus(stat string)
	GetServerStatusTime() map[string]string
	Connect()
}

type Db5 struct {
	ConfigDb map[string]interface{}
	DbConn
}



func (db5 *Db5) testIterTable() string{
	var ctx = context.Background()
	res := "ok"
	conn := db5.ConnectOpen
	param,err := conn.HMGet(ctx,"Iter", "iterIp").Result()
	if err != nil{panic("Error RedisDb5 testIterTable")}
	if param[0] == nil{res = "non"}
	return res
}

func (db5 *Db5) SetStatusIter(iter int) {
	var ctx = context.Background()
	conn := db5.ConnectOpen
	conn.HMSet(ctx,"Iter", "iterIp", iter)
}

func (db5 *Db5)GetStatusIter()int{
	var ctx = context.Background()
	conn := db5.ConnectOpen

	param,err := conn.HMGet(ctx,"Iter", "iterIp").Result()
	if err != nil{panic("Error RedisDb5 GetStatusIter")}
	if param[0] != nil{
		st :=param[0].(string)
		iterInt,errInt := strconv.Atoi(st)
		if  errInt!= nil{panic("Error RedisDb5 GetStatusIter")}
		return iterInt
	}
	return 0
}

func (db5 *Db5) UpdateServerStatus(stat string)  {
	var ctx = context.Background()
	conn := db5.ConnectOpen
	conn.HMSet(ctx,"ServerStatus", "time", helper.GetTime(),"status",stat)

}

func (db5 *Db5) GetServerStatusTime() map[string]string {
	var ctx = context.Background()
	conn := db5.ConnectOpen

	data := conn.HMGet(ctx,"ServerStatus","time","status").Val()
	resMap := map[string]string{
		"time":"non",
		"status":"non",
	}

	if len(data) != 0 {
		resMap["time"]=data[0].(string)
		resMap["status"]=data[1].(string)
	}
	return resMap
}

func (db5 *Db5) GetBreakNum() int {
	var ctx = context.Background()
	conn := db5.ConnectOpen

	param,err := conn.HMGet(ctx,"BreakNum","Break").Result()
	if err != nil{panic("Error RedisDb5 GetStatusIter")}
	if param[0] != nil{
		st :=param[0].(string)
		iterInt,errInt := strconv.Atoi(st)
		if  errInt!= nil{panic("Error RedisDb5 GetStatusIter")}
		return iterInt
	}
	return 0
}

func (db5 *Db5)Connect() {
	db5.DbConn.DbHost = db5.ConfigDb["DbHost"].(string)
	db5.DbConn.DbPort = db5.ConfigDb["DbPort"].(string)
	db5.DbConn.DbNum  = 5
	db5.DbConn.DbPassword = db5.ConfigDb["DbPassword"].(string)
	db5.DbConn.PoolSiz = db5.ConfigDb["PoolSiz"].(int)
	db5.DbConn.Connect()
}