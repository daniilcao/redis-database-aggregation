package db_conn

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type DbConn struct {
	DbHost string
	DbPort string
	DbPassword string
	DbNum int
	PoolSiz int

	ConnectOpen *redis.Client `json:"-"`
}

func (Rd *DbConn) Connect() {

	context.Background()
	var client = redis.NewClient(&redis.Options{
		Addr:     Rd.DbHost + ":" + Rd.DbPort,
		Password: Rd.DbPassword, // no password set
		PoolSize: Rd.PoolSiz,
		DB:       Rd.DbNum, // use default DB
	})

	//ctx,_ := context.WithTimeout(context.TODO(),10 * time.Second)
	ctx := context.Background()
	pong,err := client.Ping(ctx).Result()

	if err != nil {
		fmt.Println(pong, err)
		panic("Error conn db")
		return
	}

	Rd.ConnectOpen = client
}